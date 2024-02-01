/*
Copyright 2019 The Kubernetes Authors.

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

// Package completion implements the `completion` command
package completion

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/completion/bash"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/completion/fish"
	"github.com/PatrickLaabs/cli_clusterapi-argohub/cmd/argohub/completion/zsh"
)

// NewCommand returns a new cobra.Command for cluster creation
func NewCommand(streams cmd.IOStreams) *cobra.Command {
	c := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "completion",
		Short: "Output shell completion code for the specified shell (bash, zsh or fish)",
		Long:  longDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Help()
			if err != nil {
				return err
			}
			return errors.New("subcommand is required")
		},
	}
	c.AddCommand(zsh.NewCommand(streams))
	c.AddCommand(bash.NewCommand(streams))
	c.AddCommand(fish.NewCommand(streams))
	return c
}

const longDescription = `
Outputs kind shell completion for the given shell (bash or zsh)
This depends on the bash-completion binary.  Example installation instructions:
# for bash users
	$ kind completion bash > ~/.kind-completion
	$ source ~/.kind-completion

# for zsh users
	% kind completion zsh > /usr/local/share/zsh/site-functions/_kind
	% autoload -U compinit && compinit
# or if zsh-completion is installed via homebrew
    % kind completion zsh > "${fpath[1]}/_kind"
# or if you use oh-my-zsh (needs zsh-completions plugin)
	% mkdir $ZSH/completions/
	% kind completion zsh > $ZSH/completions/_kind

# for fish users
	% kind completion fish > ~/.config/fish/completions/kind.fish

Additionally, you may want to output the completion to a file and source in your .bashrc
Note for zsh users: [1] zsh completions are only supported in versions of zsh >= 5.2
`
