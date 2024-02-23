package cluster

import (
	"fmt"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/clusterapi"
	"github.com/PatrickLaabs/frigg/cmd/frigg/bootstrap/capd/helmchartproxies"
	"github.com/PatrickLaabs/frigg/internal/runtime"
	"github.com/PatrickLaabs/frigg/pkg/common/kubeconfig"
	"github.com/PatrickLaabs/frigg/pkg/common/postbootstrap"
	"github.com/PatrickLaabs/frigg/pkg/common/wait"
	"github.com/PatrickLaabs/frigg/tmpl/mgmtmanifestgen"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/PatrickLaabs/frigg/cmd"
	"github.com/PatrickLaabs/frigg/pkg/cluster"
	"github.com/PatrickLaabs/frigg/pkg/common/workdir"
	"github.com/PatrickLaabs/frigg/pkg/errors"
	"github.com/PatrickLaabs/frigg/pkg/log"

	"github.com/PatrickLaabs/frigg/internal/cli"
)

type flagpole struct {
	Name       string
	Config     string
	ImageName  string
	Retain     bool
	Wait       time.Duration
	Kubeconfig string
}

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	homedir, _ := os.UserHomeDir()

	friggDirName := ".frigg"
	kubeconfigName := "bootstrapcluster.kubeconfig"

	kubeconfigFlagPath := homedir + "/" + friggDirName + "/" + kubeconfigName

	flags := &flagpole{}
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "cluster",
		Short: "Creates a local Kubernetes cluster",
		Long:  "Creates a local Kubernetes cluster using Docker container 'nodes'",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli.OverrideDefaultName(cmd.Flags())
			return runE(logger, streams, flags)
		},
	}
	c.Flags().StringVarP(
		&flags.Name,
		"name",
		"n",
		"bootstrapcluster",
		"cluster name, overrides KIND_CLUSTER_NAME, config (default kind)",
	)
	c.Flags().StringVar(
		&flags.Config,
		"config",
		"templates/kind-config.yaml",
		"path to a kind config file",
	)
	c.Flags().StringVar(
		&flags.ImageName,
		"image",
		"",
		"node docker image to use for booting the cluster",
	)
	c.Flags().BoolVar(
		&flags.Retain,
		"retain",
		false,
		"retain nodes for debugging when cluster creation fails",
	)
	c.Flags().DurationVar(
		&flags.Wait,
		"wait",
		time.Duration(0),
		"wait for control plane node to be ready (default 0s)",
	)
	c.Flags().StringVar(
		&flags.Kubeconfig,
		"kubeconfig",
		kubeconfigFlagPath,
		"sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config",
	)
	return c
}

func runE(logger log.Logger, streams cmd.IOStreams, flags *flagpole) error {

	// Get GITHUB_TOKEN environment var
	// Will exit, if the Token is not set, since we need this Token for further configurations
	// Like deploying the token as a kubernetes secret on your clusters, creating repositories.
	if os.Getenv("GITHUB_TOKEN") == "" {
		fmt.Println("Missing Github Token, please set it. Exiting now.")
		os.Exit(1)
	} else {
		os.Getenv("GITHUB_TOKEN")
		fmt.Println("Found Github Token Environment variable. Continuing..")
	}

	// Create working directory named .frigg inside the users homedirectory.
	workdir.CreateDir()
	
	provider := cluster.NewProvider(
		cluster.ProviderWithLogger(logger),
		runtime.GetDefault(logger),
	)

	// handle config flag, we might need to read from stdin
	withConfig, err := configOption(flags.Config, streams.In)
	if err != nil {
		return err
	}

	// create the cluster
	if err = provider.Create(
		flags.Name,
		withConfig,
		cluster.CreateWithNodeImage(flags.ImageName),
		cluster.CreateWithRetain(flags.Retain),
		cluster.CreateWithWaitForReady(flags.Wait),
		cluster.CreateWithKubeconfigPath(flags.Kubeconfig),
		cluster.CreateWithDisplayUsage(true),
		cluster.CreateWithDisplaySalutation(true),
	); err != nil {
		return errors.Wrap(err, "failed to create cluster")
	}

	// Installs capi components on the bootstrap cluster. 
	// clustername is bootstrapcluster
	wait.Wait(10 * time.Second)
	clusterapi.ClusterAPI()

	// Installs a CNI solution helm chart proxy to the bootstrapcluster
	// This is needed, to make the worker nodes ready and complete the bootstrap deployment
	wait.Wait(10 * time.Second)
	helmchartproxies.InstallBootstrapHelmCharts()

	// Generates a manifest for the management cluster, named frigg-mgmt-cluster
	wait.Wait(10 * time.Second)
	mgmtmanifestgen.Gen()

	// Modifies the manifest of mgmt, to add the helmchart labels to it
	// ToDo
	// This is still in development process, hence the fact that modification on YAMLs are not that easy..
	//manifestmodifier.ModifyMgmt()

	// Applies the frigg-mgmt-cluster manifest to the bootstrap cluster
	// to create the first 'real' management cluster
	wait.Wait(5 * time.Second)
	clusterapi.KubectlApplyMgmt()

	// Retrieves the kubeconfig for the frigg-mgmt-cluster from the bootstrap cluster
	// so that we can later on use the kubeconfig to target the correct cluster for deployments.
	wait.Wait(10 * time.Second)
	kubeconfig.RetrieveMgmtKubeconfig()

	// Modifes the kubeconfig, to let us interact with the newly created kubernetes cluster.
	// On MacOS, there is an issue, where you need to replace ip and the port address, in order to 
	// successfully connect to the cluster.
	// ToDo:
	// We shall check, if the user is running on macOS, Linux and/or Windows.
	// Depending on the OS, the modification, of the kubeconfig, may not be needed.
	wait.Wait(5 * time.Second)
	err = kubeconfig.ModifyMgmtKubeconfig()
	if err != nil {
		return err
	}

	// Applies the Github Token and the default ArgoCD Login Credentials as a 
	// kubernetes secret on the argo namespace.
	// This is needed to let us interact with github, to clone, refactor and push the needed
	// gitops repositories.
	//
	// On the deployment, of new workload kubernetes clusters - which will be attached to the management
	// cluster - we run ArgoCD Workflows, which will create a pod, which runs a script. 
	// This script logs in to the argocd instance, and adds the new kubernetes cluster to it, and 
	// also adds a label to the cluster, with which we can proceed the automation steps.
	wait.Wait(5 * time.Second)
	// Github Token Secret deployment
	clusterapi.ApplyGithubSecretMgmt()
	// ArgoCD Default Login Secret deployment
	clusterapi.ApplyArgoSecretMgmt()

	// Installs the capi components to the frigg-mgmt-cluster
	// This part might take a while.
	wait.Wait(5 * time.Second)
	clusterapi.ClusterAPIMgmt()

	// Moves the capi components from the bootstrap cluster to the frigg-mgmt-cluster
	wait.Wait(5 * time.Second)
	clusterapi.Pivot()

	// Deletes the bootstrap cluster, since we don't need it any longer
	// and to free up some hardware resources.
	postbootstrap.DeleteBootstrapcluster()

	// Installs the HelmChartProxies onto the mgmt-cluster
	wait.Wait(10 * time.Second)
	helmchartproxies.InstallMgmtHelmCharts()

	// Generates a workload-cluster manifest
	// Modifies the manifest of the workload cluster, to add the helmchart labels to it
	// ToDo
	// This is still in development process, hence the fact that modification on YAMLs are not that easy..
	
	// Modifies the generated manifest with the needed helmchartproxy labels
	// This step can we "on hold", since we directly write a yaml file from the templates directory, to the
	// .frigg working directory.
	
	// Applies the workload cluster manifest to the frigg-mgmt-cluster
	wait.Wait(5 * time.Second)
	clusterapi.KubectlApplyWorkload()

	// Retrieves the kubeconfig, like we did for the management cluster previously.
	wait.Wait(10 * time.Second)
	kubeconfig.RetrieveWorkloadKubeconfig()

	// Modifies the kubeconfig, same pattern applies like for the management cluster.
	wait.Wait(5 * time.Second)
	err = kubeconfig.ModifyWorkloadKubeconfig()
	if err != nil {
		return err
	}

	return nil
}

// configOption converts the raw --config flag value to a cluster creation
// option matching it. it will read from stdin if the flag value is `-`
func configOption(rawConfigFlag string, stdin io.Reader) (cluster.CreateOption, error) {
	// if not - then we are using a real file
	if rawConfigFlag != "-" {
		return cluster.CreateWithConfigFile(rawConfigFlag), nil
	}
	// otherwise read from stdin
	raw, err := io.ReadAll(stdin)
	if err != nil {
		return nil, errors.Wrap(err, "error reading config from stdin")
	}
	return cluster.CreateWithRawConfig(raw), nil
}
