/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package clusters implements the `delete` command for multiple clusters
package clusters

import (
	"github.com/spf13/cobra"

	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/cluster"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/errors"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/pkg/log"

	"github.com/PatrickLaabs/cli_clusterapi-argohub/internal/runtime"
)

type flagpole struct {
	Kubeconfig string
	All        bool
}

// NewCommand returns a new cobra.Command for cluster deletion
func NewCommand(logger log.Logger) *cobra.Command {
	flags := &flagpole{}
	c := &cobra.Command{
		Args:  cobra.MinimumNArgs(0),
		Use:   "clusters",
		Short: "Deletes one or more clusters",
		Long:  "Deletes a resource",
		RunE: func(cmd *cobra.Command, args []string) error {
			if !flags.All && len(args) == 0 {
				return errors.New("no cluster names provided")
			}

			return deleteClusters(logger, flags, args)
		},
	}
	c.Flags().StringVar(
		&flags.Kubeconfig,
		"kubeconfig",
		"",
		"sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config",
	)
	c.Flags().BoolVarP(
		&flags.All,
		"all",
		"A",
		false,
		"delete all clusters",
	)
	return c
}

func deleteClusters(logger log.Logger, flags *flagpole, clusters []string) error {
	provider := cluster.NewProvider(
		cluster.ProviderWithLogger(logger),
		runtime.GetDefault(logger),
	)
	var err error
	if flags.All {
		//Delete all clusters
		if clusters, err = provider.List(); err != nil {
			return errors.Wrap(err, "failed listing clusters for delete")
		}
	}
	var success []string
	for _, c := range clusters {
		if err = provider.Delete(c, flags.Kubeconfig); err != nil {
			logger.V(0).Infof("%s\n", errors.Wrapf(err, "failed to delete c %q", c))
			continue
		}
		success = append(success, c)
	}
	logger.V(0).Infof("Deleted clusters: %q", success)
	return nil
}
