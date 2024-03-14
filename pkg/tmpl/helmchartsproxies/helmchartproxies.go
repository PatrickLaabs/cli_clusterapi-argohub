package helmchartsproxies

// The package helmchartproxies generates various helmchartproxy yaml files to the .frigg directory
// which will be installed onto the management cluster.
import (
	"fmt"
	"github.com/PatrickLaabs/frigg/pkg/consts"
	"github.com/PatrickLaabs/frigg/pkg/vars"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

type CniAutoGenerated struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		ClusterSelector struct {
			MatchLabels struct {
			} `yaml:"matchLabels"`
		} `yaml:"clusterSelector"`
		ReleaseName    string `yaml:"releaseName"`
		RepoURL        string `yaml:"repoURL"`
		ChartName      string `yaml:"chartName"`
		Namespace      string `yaml:"namespace"`
		Version        string `yaml:"version"`
		ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
	} `yaml:"spec"`
}

type VaultAutoGenerated struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		ClusterSelector struct {
			MatchLabels struct {
				VaultmgmtEnabled string `yaml:"vaultmgmtEnabled"`
			} `yaml:"matchLabels"`
		} `yaml:"clusterSelector"`
		RepoURL     string `yaml:"repoURL"`
		ChartName   string `yaml:"chartName"`
		Namespace   string `yaml:"namespace"`
		ReleaseName string `yaml:"releaseName"`
		Version     string `yaml:"version"`
		Options     struct {
			WaitForJobs bool   `yaml:"waitForJobs"`
			Wait        bool   `yaml:"wait"`
			Timeout     string `yaml:"timeout"`
			Install     struct {
				CreateNamespace bool `yaml:"createNamespace"`
			} `yaml:"install"`
		} `yaml:"options"`
		ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
	} `yaml:"spec"`
}

type ArgoCDWorkloadClustersAutoGenerated struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		ClusterSelector struct {
			MatchLabels struct {
				ArgocdWorkloadclustersClustersEnabled string `yaml:"argocd-workloadclusters-clusters-Enabled"`
			} `yaml:"matchLabels"`
		} `yaml:"clusterSelector"`
		RepoURL     string `yaml:"repoURL"`
		ChartName   string `yaml:"chartName"`
		Namespace   string `yaml:"namespace"`
		ReleaseName string `yaml:"releaseName"`
		Version     string `yaml:"version"`
		Options     struct {
			WaitForJobs bool   `yaml:"waitForJobs"`
			Wait        bool   `yaml:"wait"`
			Timeout     string `yaml:"timeout"`
			Install     struct {
				CreateNamespace bool `yaml:"createNamespace"`
			} `yaml:"install"`
		} `yaml:"options"`
		ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
	} `yaml:"spec"`
}

type ArgoEventsAutoGenerated struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		ClusterSelector struct {
			MatchLabels struct {
				ArgoEventsEnabled string `yaml:"argo-events-Enabled"`
			} `yaml:"matchLabels"`
		} `yaml:"clusterSelector"`
		RepoURL     string `yaml:"repoURL"`
		ChartName   string `yaml:"chartName"`
		Namespace   string `yaml:"namespace"`
		ReleaseName string `yaml:"releaseName"`
		Version     string `yaml:"version"`
		Options     struct {
			WaitForJobs bool   `yaml:"waitForJobs"`
			Wait        bool   `yaml:"wait"`
			Timeout     string `yaml:"timeout"`
			Install     struct {
				CreateNamespace bool `yaml:"createNamespace"`
			} `yaml:"install"`
		} `yaml:"options"`
	} `yaml:"spec"`
}

type ArgoRolloutsAutoGenerated struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		ClusterSelector struct {
			MatchLabels struct {
				ArgoRolloutsEnabled string `yaml:"argo-rollouts-Enabled"`
			} `yaml:"matchLabels"`
		} `yaml:"clusterSelector"`
		RepoURL     string `yaml:"repoURL"`
		ChartName   string `yaml:"chartName"`
		Namespace   string `yaml:"namespace"`
		ReleaseName string `yaml:"releaseName"`
		Version     string `yaml:"version"`
		Options     struct {
			WaitForJobs bool   `yaml:"waitForJobs"`
			Wait        bool   `yaml:"wait"`
			Timeout     string `yaml:"timeout"`
			Install     struct {
				CreateNamespace bool `yaml:"createNamespace"`
			} `yaml:"install"`
		} `yaml:"options"`
		ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
	} `yaml:"spec"`
}

type ArgoWorkflowsAutoGenerated struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		ClusterSelector struct {
			MatchLabels struct {
				ArgoWorkflowsEnabled string `yaml:"argo-workflows-Enabled"`
			} `yaml:"matchLabels"`
		} `yaml:"clusterSelector"`
		RepoURL     string `yaml:"repoURL"`
		ChartName   string `yaml:"chartName"`
		Namespace   string `yaml:"namespace"`
		ReleaseName string `yaml:"releaseName"`
		Version     string `yaml:"version"`
		Options     struct {
			WaitForJobs bool   `yaml:"waitForJobs"`
			Wait        bool   `yaml:"wait"`
			Timeout     string `yaml:"timeout"`
			Install     struct {
				CreateNamespace bool `yaml:"createNamespace"`
			} `yaml:"install"`
		} `yaml:"options"`
		ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
	} `yaml:"spec"`
}

type ArgoMgmtAppsAutoGenerated struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		ClusterSelector struct {
			MatchLabels struct {
				ArgocdHubAppsEnalbed string `yaml:"argocd-hub-appsEnabled"`
			} `yaml:"matchLabels"`
		} `yaml:"clusterSelector"`
		RepoURL     string `yaml:"repoURL"`
		ChartName   string `yaml:"chartName"`
		Namespace   string `yaml:"namespace"`
		ReleaseName string `yaml:"releaseName"`
		Version     string `yaml:"version"`
		Options     struct {
			WaitForJobs bool   `yaml:"waitForJobs"`
			Wait        bool   `yaml:"wait"`
			Timeout     string `yaml:"timeout"`
			Install     struct {
				CreateNamespace bool `yaml:"createNamespace"`
			} `yaml:"install"`
		} `yaml:"options"`
		ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
	} `yaml:"spec"`
}

type ArgoMgmtAutoGenerated struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		ClusterSelector struct {
			MatchLabels struct {
				ArgocdHubEnalbed string `yaml:"argocd-hub-Enabled"`
			} `yaml:"matchLabels"`
		} `yaml:"clusterSelector"`
		RepoURL     string `yaml:"repoURL"`
		ChartName   string `yaml:"chartName"`
		Namespace   string `yaml:"namespace"`
		ReleaseName string `yaml:"releaseName"`
		Version     string `yaml:"version"`
		Options     struct {
			WaitForJobs bool   `yaml:"waitForJobs"`
			Wait        bool   `yaml:"wait"`
			Timeout     string `yaml:"timeout"`
			Install     struct {
				CreateNamespace bool `yaml:"createNamespace"`
			} `yaml:"install"`
		} `yaml:"options"`
		ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
	} `yaml:"spec"`
}

type ClusterApiOperatorAutoGenerated struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
	} `yaml:"metadata"`
	Spec struct {
		ClusterSelector struct {
			MatchLabels struct {
				CapiOperatorEnabled string `yaml:"capi-operator-Enabled"`
			} `yaml:"matchLabels"`
		} `yaml:"clusterSelector"`
		RepoURL     string `yaml:"repoURL"`
		ChartName   string `yaml:"chartName"`
		Namespace   string `yaml:"namespace"`
		ReleaseName string `yaml:"releaseName"`
		Version     string `yaml:"version"`
		Options     struct {
			WaitForJobs bool   `yaml:"waitForJobs"`
			Wait        bool   `yaml:"wait"`
			Timeout     string `yaml:"timeout"`
			Install     struct {
				CreateNamespace bool `yaml:"createNamespace"`
			} `yaml:"install"`
		} `yaml:"options"`
		ValuesTemplate string `yaml:"valuesTemplate"`
	} `yaml:"spec"`
}

// Cni generates the CNI Helmchartproxy YAML file to the .frigg directory
func Cni() {
	data := &CniAutoGenerated{
		APIVersion: "addons.cluster.x-k8s.io/v1alpha1",
		Kind:       "HelmChartProxy",
		Metadata: struct {
			Name string `yaml:"name"`
		}{Name: "calico-cni"},
		Spec: struct {
			ClusterSelector struct {
				MatchLabels struct{} `yaml:"matchLabels"`
			} `yaml:"clusterSelector"`
			ReleaseName    string `yaml:"releaseName"`
			RepoURL        string `yaml:"repoURL"`
			ChartName      string `yaml:"chartName"`
			Namespace      string `yaml:"namespace"`
			Version        string `yaml:"version"`
			ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
		}{
			ReleaseName: "calico",
			RepoURL:     "https://docs.tigera.io/calico/charts",
			ChartName:   "tigera-operator",
			Namespace:   "kube-system",
			Version:     "3.27.0",
			ValuesTemplate: `installation:
  registry: quay.io/`,
		},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marsheling yaml data %v\n", err))
	}

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	newfilePath := filepath.Join(friggDir, vars.CniName)

	// Write to file
	err = os.WriteFile(newfilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing cni yaml: %v\n", err))
	}

}

// Vault generates the Vault Helmchartproxy YAML file to the .frigg directory
func Vault() {
	data := VaultAutoGenerated{
		APIVersion: "addons.cluster.x-k8s.io/v1alpha1",
		Kind:       "HelmChartProxy",
		Metadata: struct {
			Name string `yaml:"name"`
		}{
			Name: "vault",
		},
		Spec: struct {
			ClusterSelector struct {
				MatchLabels struct {
					VaultmgmtEnabled string `yaml:"vaultmgmtEnabled"`
				} `yaml:"matchLabels"`
			} `yaml:"clusterSelector"`
			RepoURL     string `yaml:"repoURL"`
			ChartName   string `yaml:"chartName"`
			Namespace   string `yaml:"namespace"`
			ReleaseName string `yaml:"releaseName"`
			Version     string `yaml:"version"`
			Options     struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			} `yaml:"options"`
			ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
		}{
			ClusterSelector: struct {
				MatchLabels struct {
					VaultmgmtEnabled string `yaml:"vaultmgmtEnabled"`
				} `yaml:"matchLabels"`
			}{
				MatchLabels: struct {
					VaultmgmtEnabled string `yaml:"vaultmgmtEnabled"`
				}{
					VaultmgmtEnabled: "default",
				},
			},
			RepoURL:     "https://helm.releases.hashicorp.com",
			ChartName:   "vault",
			Namespace:   "vault",
			ReleaseName: "vault",
			Version:     "0.27.0",
			Options: struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			}{
				WaitForJobs: true,
				Wait:        true,
				Timeout:     "5m",
				Install: struct {
					CreateNamespace bool `yaml:"createNamespace"`
				}{
					CreateNamespace: true,
				},
			},
			ValuesTemplate: `server:
  ha:
    config: |
      ui = true
      listener "tcp" {
        tls_disable     = 1
        address         = "[::]:8200"
        cluster_address = "[::]:8201"
      }
      service_registration "kubernetes" {}
ui:
  enabled: true
  serviceType: "ClusterIP"
  serviceNodePort: null
  externalPort: 8200
  targetPort: 8200`,
		},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marsheling yaml data %v\n", err))
	}

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	newfilePath := filepath.Join(friggDir, vars.VaultName)

	// Write to file
	err = os.WriteFile(newfilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing vault yaml: %v\n", err))
	}
}

// ArgoCDWorkloadClusters generates the ArgoCD Helmchartproxy YAML file to the .frigg directory
func ArgoCDWorkloadClusters() {
	data := ArgoCDWorkloadClustersAutoGenerated{
		APIVersion: "addons.cluster.x-k8s.io/v1alpha1",
		Kind:       "HelmChartProxy",
		Metadata: struct {
			Name string `yaml:"name"`
		}{
			Name: "argocd-workloadclusters-clusters",
		},
		Spec: struct {
			ClusterSelector struct {
				MatchLabels struct {
					ArgocdWorkloadclustersClustersEnabled string `yaml:"argocd-workloadclusters-clusters-Enabled"`
				} `yaml:"matchLabels"`
			} `yaml:"clusterSelector"`
			RepoURL     string `yaml:"repoURL"`
			ChartName   string `yaml:"chartName"`
			Namespace   string `yaml:"namespace"`
			ReleaseName string `yaml:"releaseName"`
			Version     string `yaml:"version"`
			Options     struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			} `yaml:"options"`
			ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
		}{
			ClusterSelector: struct {
				MatchLabels struct {
					ArgocdWorkloadclustersClustersEnabled string `yaml:"argocd-workloadclusters-clusters-Enabled"`
				} `yaml:"matchLabels"`
			}{
				MatchLabels: struct {
					ArgocdWorkloadclustersClustersEnabled string `yaml:"argocd-workloadclusters-clusters-Enabled"`
				}{
					ArgocdWorkloadclustersClustersEnabled: "default",
				},
			},
			RepoURL:     "https://argoproj.github.io/argo-helm",
			ChartName:   "argo-cd",
			Namespace:   "argocd",
			ReleaseName: "argocd",
			Version:     "5.52.1",
			Options: struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			}{
				WaitForJobs: true,
				Wait:        true,
				Timeout:     "5m",
				Install: struct {
					CreateNamespace bool `yaml:"createNamespace"`
				}{
					CreateNamespace: true,
				},
			},
			ValuesTemplate: `server:
  name: server
  crds:
    keep: false
  configs:
    params:
      server.insecure: true
    ssh:
      knownHosts: |
        github.com ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOMqqnkVzrm0SdG6UOoqKLsabgH5C9okWi0dh2l9GKJl
        github.com ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEmKSENjQEezOmxkZMy7opKgwFB9nkt5YRrYMjNuG5N87uRgg6CLrbo5wAdT/y6v0mKV0U2w0WZ2YB/++Tpockg=
        github.com ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCj7ndNxQowgcQnjshcLrqPEiiphnt+VTTvDP6mHBL9j1aNUkY4Ue1gvwnGLVlOhGeYrnZaMgRK6+PKCUXaDbC7qtbW8gIkhL7aGCsOr/C56SJMy/BCZfxd1nWzAOxSDPgVsmerOBYfNqltV9/hWCqBywINIR+5dIg6JTJ72pcEpEjcYgXkE2YEFXV1JHnsKgbLWNlhScqb2UmyRkQyytRLtL+38TGxkxCflmO+5Z8CSSNY7GidjMIZ7Q4zMjA2n1nGrlTDkzwDCsw+wqFPGQA179cnfGWOWRVruj16z6XyvxvjJwbz0wQZ75XK5tKSb7FNyeIEs4TT4jk+S4dhPeAUC5y+bDYirYgM4GC7uEnztnZyaVWQ7B381AK4Qdrwt51ZqExKbQpTUNn+EjqoTwvqNj4kqx5QUCI0ThS/YkOxJCXmPUWZbhjpCg56i+2aB6CmK2JGhn57K5mj0MNdBXA4/WnwH6XoPWJzK5Nyu2zB3nAZp+S5hpQs+p1vN1/wsjk=
    secret:
      argocdServerAdminPassword: "$2a$10$UfHxzEstRBKFAiTH0ZlI8u95SOaRBcXDCxBTBxfmOz14FHC6Vv3de"`,
		},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marsheling yaml data %v\n", err))
	}

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	newfilePath := filepath.Join(friggDir, vars.ArgoCDWorkload)

	// Write to file
	err = os.WriteFile(newfilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing argocd workloadclusters yaml: %v\n", err))
	}

}

// ArgoWorkflows generates the ArgoCD Workflows Helmchartproxy YAML file to the .frigg directory
func ArgoWorkflows() {
	data := ArgoWorkflowsAutoGenerated{
		APIVersion: "addons.cluster.x-k8s.io/v1alpha1",
		Kind:       "HelmChartProxy",
		Metadata: struct {
			Name string `yaml:"name"`
		}{
			Name: "argo-workflows",
		},
		Spec: struct {
			ClusterSelector struct {
				MatchLabels struct {
					ArgoWorkflowsEnabled string `yaml:"argo-workflows-Enabled"`
				} `yaml:"matchLabels"`
			} `yaml:"clusterSelector"`
			RepoURL     string `yaml:"repoURL"`
			ChartName   string `yaml:"chartName"`
			Namespace   string `yaml:"namespace"`
			ReleaseName string `yaml:"releaseName"`
			Version     string `yaml:"version"`
			Options     struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			} `yaml:"options"`
			ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
		}{
			ClusterSelector: struct {
				MatchLabels struct {
					ArgoWorkflowsEnabled string `yaml:"argo-workflows-Enabled"`
				} `yaml:"matchLabels"`
			}{
				MatchLabels: struct {
					ArgoWorkflowsEnabled string `yaml:"argo-workflows-Enabled"`
				}{
					ArgoWorkflowsEnabled: "default",
				},
			},
			RepoURL:     "https://argoproj.github.io/argo-helm",
			ChartName:   "argo-workflows",
			Namespace:   "argocd",
			ReleaseName: "argo-workflows",
			Version:     "0.40.5",
			Options: struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			}{
				WaitForJobs: true,
				Wait:        true,
				Timeout:     "5m",
				Install: struct {
					CreateNamespace bool `yaml:"createNamespace"`
				}{
					CreateNamespace: false,
				},
			},
			ValuesTemplate: `workflow:
  serviceAccount:
    create: true
    name: "argocd-workflow"
  rbac:
    create: true
controller:
  rbac:
    create: true
    accessAllSecrets: true
  workflowNamespaces:
    - argo
extraObjects:
  - apiVersion: rbac.authorization.k8s.io/v1
    kind: ClusterRole
    metadata:
      name: argocd-workflow-cluster
    rules:
    - apiGroups: ["cluster.x-k8s.io"]
      resources: ["clusters"]
      verbs: ["get", "list", "watch"]
    - apiGroups: ["infrastructure.cluster.x-k8s.io"]
      resources: ["vsphereclusters", "azureclusters", "dockerclusters"]
      verbs: ["get", "list", "watch"]
    - apiGroups: [""]
      resources: ["secrets", "configmaps"]
      verbs: ["get", "list", "watch"]
  - apiVersion: rbac.authorization.k8s.io/v1
    kind: ClusterRoleBinding
    metadata:
      name: argocd-workflow-cluster
    subjects:
    - kind: ServiceAccount
      name: argocd-workflow
      namespace: argo
    roleRef:
      kind: ClusterRole
      name: argocd-workflow-cluster
      apiGroup: rbac.authorization.k8s.io
  - apiVersion: rbac.authorization.k8s.io/v1
    kind: Role
    metadata:
      name: argo-events-workflows
      namespace: argo
    rules:
    - apiGroups: ["argoproj.io"]
      verbs: ["*"]
      resources: ["workflows", "workflowtemplates"]
  - apiVersion: rbac.authorization.k8s.io/v1
    kind: RoleBinding
    metadata:
      name: argocd-events-workflows
      namespace: argo
    subjects:
    - kind: ServiceAccount
      name: argocd-workflow
      namespace: argo
    roleRef:
      kind: Role
      name: argo-events-workflows
      apiGroup: rbac.authorization.k8s.io
server:
  extraArgs:
  - --auth-mode=server
  - --auth-mode=client
  sso:
    enabled: false
  secure: false
  ingress:
    enabled: false`,
		},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marsheling yaml data %v\n", err))
	}

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	newfilePath := filepath.Join(friggDir, vars.ArgoWorkflowsMgmt)

	// Write to file
	err = os.WriteFile(newfilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing argocd workflows yaml: %v\n", err))
	}

}

// ArgoRollouts generates the ArgoCD Rollouts Helmchartproxy YAML file to the .frigg directory
func ArgoRollouts() {
	data := ArgoRolloutsAutoGenerated{
		APIVersion: "addons.cluster.x-k8s.io/v1alpha1",
		Kind:       "HelmChartProxy",
		Metadata: struct {
			Name string `yaml:"name"`
		}{
			Name: "argo-rollouts",
		},
		Spec: struct {
			ClusterSelector struct {
				MatchLabels struct {
					ArgoRolloutsEnabled string `yaml:"argo-rollouts-Enabled"`
				} `yaml:"matchLabels"`
			} `yaml:"clusterSelector"`
			RepoURL     string `yaml:"repoURL"`
			ChartName   string `yaml:"chartName"`
			Namespace   string `yaml:"namespace"`
			ReleaseName string `yaml:"releaseName"`
			Version     string `yaml:"version"`
			Options     struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			} `yaml:"options"`
			ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
		}{
			ClusterSelector: struct {
				MatchLabels struct {
					ArgoRolloutsEnabled string `yaml:"argo-rollouts-Enabled"`
				} `yaml:"matchLabels"`
			}{
				MatchLabels: struct {
					ArgoRolloutsEnabled string `yaml:"argo-rollouts-Enabled"`
				}{
					ArgoRolloutsEnabled: "default",
				},
			},
			RepoURL:     "https://argoproj.github.io/argo-helm",
			ChartName:   "argo-rollouts",
			Namespace:   "argo-rollouts",
			ReleaseName: "argo-rollouts",
			Version:     "2.34.1",
			Options: struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			}{
				WaitForJobs: true,
				Wait:        true,
				Timeout:     "5m",
				Install: struct {
					CreateNamespace bool `yaml:"createNamespace"`
				}{
					CreateNamespace: true,
				},
			},
			ValuesTemplate: `dashboard:
  enabled: false
  service:
    type: ClusterIP
    portName: dashboard
    port: 3100
    portTarget: 3100
  serviceAccount:
    create: true
  ingress:
    enabled: false`,
		},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marsheling yaml data %v\n", err))
	}

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	newfilePath := filepath.Join(friggDir, vars.ArgoRolloutsMgmt)

	// Write to file
	err = os.WriteFile(newfilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing argocd rollouts yaml: %v\n", err))
	}

}

// ArgoEvents generates the ArgoCD Events Helmchartproxy YAML file to the .frigg directory
func ArgoEvents() {
	data := ArgoEventsAutoGenerated{
		APIVersion: "addons.cluster.x-k8s.io/v1alpha1",
		Kind:       "HelmChartProxy",
		Metadata: struct {
			Name string `yaml:"name"`
		}{
			Name: "argo-events",
		},
		Spec: struct {
			ClusterSelector struct {
				MatchLabels struct {
					ArgoEventsEnabled string `yaml:"argo-events-Enabled"`
				} `yaml:"matchLabels"`
			} `yaml:"clusterSelector"`
			RepoURL     string `yaml:"repoURL"`
			ChartName   string `yaml:"chartName"`
			Namespace   string `yaml:"namespace"`
			ReleaseName string `yaml:"releaseName"`
			Version     string `yaml:"version"`
			Options     struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			} `yaml:"options"`
		}{
			ClusterSelector: struct {
				MatchLabels struct {
					ArgoEventsEnabled string `yaml:"argo-events-Enabled"`
				} `yaml:"matchLabels"`
			}{
				MatchLabels: struct {
					ArgoEventsEnabled string `yaml:"argo-events-Enabled"`
				}{
					ArgoEventsEnabled: "default",
				},
			},
			RepoURL:     "https://argoproj.github.io/argo-helm",
			ChartName:   "argo-events",
			Namespace:   "argocd",
			ReleaseName: "argo-events",
			Version:     "2.4.2",
			Options: struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			}{
				WaitForJobs: true,
				Wait:        true,
				Timeout:     "5m",
				Install: struct {
					CreateNamespace bool `yaml:"createNamespace"`
				}{
					CreateNamespace: false,
				},
			},
		},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marsheling yaml data %v\n", err))
	}

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	newfilePath := filepath.Join(friggDir, vars.ArgoEventsMgmt)

	// Write to file
	err = os.WriteFile(newfilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing argocd events yaml: %v\n", err))
	}

}

// MgmtArgoApps generates the ArgoCD Apps Helmchartproxy YAML file to the .frigg directory
func MgmtArgoApps() {
	username, err := retrieveUsername()
	if err != nil {
		println(color.RedString("Error retrieving token: %v\n", err))
		os.Exit(1)
	}

	url := "git@github.com:" + username + "/" + vars.FriggMgmtGitOpsName

	data := ArgoMgmtAppsAutoGenerated{
		APIVersion: "addons.cluster.x-k8s.io/v1alpha1",
		Kind:       "HelmChartProxy",
		Metadata: struct {
			Name string `yaml:"name"`
		}{
			Name: "argocd-hub-apps",
		},
		Spec: struct {
			ClusterSelector struct {
				MatchLabels struct {
					ArgocdHubAppsEnalbed string `yaml:"argocd-hub-appsEnabled"`
				} `yaml:"matchLabels"`
			} `yaml:"clusterSelector"`
			RepoURL     string `yaml:"repoURL"`
			ChartName   string `yaml:"chartName"`
			Namespace   string `yaml:"namespace"`
			ReleaseName string `yaml:"releaseName"`
			Version     string `yaml:"version"`
			Options     struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			} `yaml:"options"`
			ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
		}{
			ClusterSelector: struct {
				MatchLabels struct {
					ArgocdHubAppsEnalbed string `yaml:"argocd-hub-appsEnabled"`
				} `yaml:"matchLabels"`
			}{
				MatchLabels: struct {
					ArgocdHubAppsEnalbed string `yaml:"argocd-hub-appsEnabled"`
				}{
					ArgocdHubAppsEnalbed: "default",
				},
			},
			RepoURL:     "https://argoproj.github.io/argo-helm",
			ChartName:   "argocd-apps",
			Namespace:   "argocd-apps",
			ReleaseName: "argocd-apps",
			Version:     "1.4.1",
			Options: struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			}{
				WaitForJobs: true,
				Wait:        true,
				Timeout:     "5m",
				Install: struct {
					CreateNamespace bool `yaml:"createNamespace"`
				}{
					CreateNamespace: true,
				},
			},
			ValuesTemplate: `applications:
  - name: argocd-hub
    namespace: argocd
    annotations:
      argocd.argoproj.io/sync-wave: '0'
    finalizers:
    - resources-finalizer.argocd.argoproj.io
    project: default
    source:
      repoURL: ` + url + `
      path: frigg-mgmt-cluster
      targetRevision: main
    destination:
      server: https://kubernetes.default.svc
      namespace: argocd
    syncPolicy:
      automated:
        prune: true
        selfHeal: true`,
		},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marsheling yaml data %v\n", err))
	}

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	newfilePath := filepath.Join(friggDir, vars.ArgoCDAppsMgmt)

	// Write to file
	err = os.WriteFile(newfilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing argocd apps yaml: %v\n", err))
	}

}

// MgmtArgoCD generates the ArgoCD Apps Helmchartproxy YAML file to the .frigg directory
func MgmtArgoCD() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	newfilePath := filepath.Join(friggDir, vars.ArgoCDMgmt)
	sshprivatekeyPath := filepath.Join(friggDir, vars.PrivatekeyName)

	username, err := retrieveUsername()
	if err != nil {
		println(color.RedString("Error retrieving token: %v\n", err))
		os.Exit(1)
	}

	url := "git@github.com:" + username + "/" + vars.FriggMgmtGitOpsName

	sshprivatekey, err := os.ReadFile(sshprivatekeyPath)
	// Remove trailing newline character (if present)
	trimmedKey := strings.TrimSuffix(string(sshprivatekey), "\n")
	// Replace newlines with 8 spaces of indentation
	formattedKey := strings.ReplaceAll(trimmedKey, "\n", "\n        ")

	data := ArgoMgmtAutoGenerated{
		APIVersion: "addons.cluster.x-k8s.io/v1alpha1",
		Kind:       "HelmChartProxy",
		Metadata: struct {
			Name string `yaml:"name"`
		}{
			Name: "argocd-hub",
		},
		Spec: struct {
			ClusterSelector struct {
				MatchLabels struct {
					ArgocdHubEnalbed string `yaml:"argocd-hub-Enabled"`
				} `yaml:"matchLabels"`
			} `yaml:"clusterSelector"`
			RepoURL     string `yaml:"repoURL"`
			ChartName   string `yaml:"chartName"`
			Namespace   string `yaml:"namespace"`
			ReleaseName string `yaml:"releaseName"`
			Version     string `yaml:"version"`
			Options     struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			} `yaml:"options"`
			ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
		}{
			ClusterSelector: struct {
				MatchLabels struct {
					ArgocdHubEnalbed string `yaml:"argocd-hub-Enabled"`
				} `yaml:"matchLabels"`
			}{
				MatchLabels: struct {
					ArgocdHubEnalbed string `yaml:"argocd-hub-Enabled"`
				}{
					ArgocdHubEnalbed: "default",
				},
			},
			RepoURL:     "https://argoproj.github.io/argo-helm",
			ChartName:   "argo-cd",
			Namespace:   "argocd",
			ReleaseName: "argocd",
			Version:     "5.52.1",
			Options: struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			}{
				WaitForJobs: true,
				Wait:        true,
				Timeout:     "5m",
				Install: struct {
					CreateNamespace bool `yaml:"createNamespace"`
				}{
					CreateNamespace: true,
				},
			},
			ValuesTemplate: `server:
  name: server
crds:
  keep: false
configs:
  params:
    server.insecure: true
  ssh:
    knownHosts: |
      github.com ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIOMqqnkVzrm0SdG6UOoqKLsabgH5C9okWi0dh2l9GKJl
      github.com ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEmKSENjQEezOmxkZMy7opKgwFB9nkt5YRrYMjNuG5N87uRgg6CLrbo5wAdT/y6v0mKV0U2w0WZ2YB/++Tpockg=
      github.com ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCj7ndNxQowgcQnjshcLrqPEiiphnt+VTTvDP6mHBL9j1aNUkY4Ue1gvwnGLVlOhGeYrnZaMgRK6+PKCUXaDbC7qtbW8gIkhL7aGCsOr/C56SJMy/BCZfxd1nWzAOxSDPgVsmerOBYfNqltV9/hWCqBywINIR+5dIg6JTJ72pcEpEjcYgXkE2YEFXV1JHnsKgbLWNlhScqb2UmyRkQyytRLtL+38TGxkxCflmO+5Z8CSSNY7GidjMIZ7Q4zMjA2n1nGrlTDkzwDCsw+wqFPGQA179cnfGWOWRVruj16z6XyvxvjJwbz0wQZ75XK5tKSb7FNyeIEs4TT4jk+S4dhPeAUC5y+bDYirYgM4GC7uEnztnZyaVWQ7B381AK4Qdrwt51ZqExKbQpTUNn+EjqoTwvqNj4kqx5QUCI0ThS/YkOxJCXmPUWZbhjpCg56i+2aB6CmK2JGhn57K5mj0MNdBXA4/WnwH6XoPWJzK5Nyu2zB3nAZp+S5hpQs+p1vN1/wsjk=
  credentialTemplates:
    ssh-crds-mgmt:
      url: ` + url + `
      sshPrivateKey: |
        ` + formattedKey + `
  repositories:
    private-repo-mgmt:
      url: ` + url + `
  secret:
    argocdServerAdminPassword: "$2a$10$UfHxzEstRBKFAiTH0ZlI8u95SOaRBcXDCxBTBxfmOz14FHC6Vv3de"`,
		},
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marsheling yaml data %v\n", err))
	}

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	// Write to file
	err = os.WriteFile(newfilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing argocd yaml: %v\n", err))
	}
}

// MgmtClusterApiOperator generates the ClusterAPI Operator Helmchartproxy YAML file to the .frigg directory
func MgmtClusterApiOperator() {
	data := ClusterApiOperatorAutoGenerated{
		APIVersion: "addons.cluster.x-k8s.io/v1alpha1",
		Kind:       "HelmChartProxy",
		Metadata: struct {
			Name string `yaml:"name"`
		}{
			Name: "capi-operator",
		},
		Spec: struct {
			ClusterSelector struct {
				MatchLabels struct {
					CapiOperatorEnabled string `yaml:"capi-operator-Enabled"`
				} `yaml:"matchLabels"`
			} `yaml:"clusterSelector"`
			RepoURL     string `yaml:"repoURL"`
			ChartName   string `yaml:"chartName"`
			Namespace   string `yaml:"namespace"`
			ReleaseName string `yaml:"releaseName"`
			Version     string `yaml:"version"`
			Options     struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			} `yaml:"options"`
			ValuesTemplate string `yaml:"valuesTemplate"`
		}(struct {
			ClusterSelector struct {
				MatchLabels struct {
					CapiOperatorEnabled string `yaml:"capi-operator-Enabled"`
				} `yaml:"matchLabels"`
			} `yaml:"clusterSelector"`
			RepoURL     string `yaml:"repoURL"`
			ChartName   string `yaml:"chartName"`
			Namespace   string `yaml:"namespace"`
			ReleaseName string `yaml:"releaseName"`
			Version     string `yaml:"version"`
			Options     struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			} `yaml:"options"`
			ValuesTemplate string `yaml:"valuesTemplate,omitempty"`
		}{
			ClusterSelector: struct {
				MatchLabels struct {
					CapiOperatorEnabled string `yaml:"capi-operator-Enabled"`
				} `yaml:"matchLabels"`
			}{
				MatchLabels: struct {
					CapiOperatorEnabled string `yaml:"capi-operator-Enabled"`
				}{
					CapiOperatorEnabled: "default",
				},
			},
			RepoURL:     "https://kubernetes-sigs.github.io/cluster-api-operator",
			ChartName:   "cluster-api-operator",
			Namespace:   "capi-operator-system",
			ReleaseName: "cluster-api-operator",
			Version:     "0.9.0",
			Options: struct {
				WaitForJobs bool   `yaml:"waitForJobs"`
				Wait        bool   `yaml:"wait"`
				Timeout     string `yaml:"timeout"`
				Install     struct {
					CreateNamespace bool `yaml:"createNamespace"`
				} `yaml:"install"`
			}{
				WaitForJobs: true,
				Wait:        true,
				Timeout:     "5m",
				Install: struct {
					CreateNamespace bool `yaml:"createNamespace"`
				}{
					CreateNamespace: true,
				},
			},
			ValuesTemplate: `cert-manager:
  enabled: false
core: "capi-system:` + consts.ClusterApiVersion + `"
bootstrap: "capi-system:` + consts.KubeadmVersion + `"
controlPlane: "capi-system:` + consts.KubeadmVersion + `"
infrastructure: "capd-system:` + consts.DockerInfraVersion + `"
addon: "capi-system:` + consts.CaaphVersion + `"
manager.featureGates:
  MachinePool: true
  ClusterTopology: true`,
		}),
	}

	// Marshal to YAML
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		println(color.RedString("error on marsheling yaml data %v\n", err))
	}

	// Prepend "---" to the YAML data
	yamlData = append([]byte("---\n"), yamlData...)

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("error on accessing home directory: %v\n", err))
	}

	friggDir := filepath.Join(homedir, vars.FriggDirName)
	newfilePath := filepath.Join(friggDir, vars.ClusterApiHelmChartProxyName)

	// Write to file
	err = os.WriteFile(newfilePath, yamlData, 0644)
	if err != nil {
		println(color.RedString("error on writing argocd apps yaml: %v\n", err))
	}

}

func retrieveUsername() (string, error) {
	// Get GITHUB_USERNAME environment var
	var username string

	if os.Getenv("GITHUB_USERNAME") == "" {
		fmt.Println("Missing Github Username, please set it. Exiting now.")
		os.Exit(1)
	} else {
		username = os.Getenv("GITHUB_USERNAME")
	}

	return username, nil
}
