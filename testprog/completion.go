/*
Copyright The Helm Authors.
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

package main

import (
	"io"

	"github.com/spf13/cobra"
)

var disableCompDescriptions bool

func newCompletionCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{Use: "completion"}
	cmd.PersistentFlags().BoolVar(&disableCompDescriptions, "no-descriptions", false, "disable completion descriptions")

	bash := &cobra.Command{
		Use: "bash",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Root().GenBashCompletion(out)
		},
	}

	zsh := &cobra.Command{
		Use: "zsh",
		RunE: func(cmd *cobra.Command, args []string) error {
			if disableCompDescriptions {
				return cmd.Root().GenZshCompletionNoDesc(out)
			} else {
				return cmd.Root().GenZshCompletion(out)
			}
		},
	}

	fish := &cobra.Command{
		Use: "fish",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Root().GenFishCompletion(out, !disableCompDescriptions)
		},
	}

	cmd.AddCommand(bash, zsh, fish)

	return cmd
}
