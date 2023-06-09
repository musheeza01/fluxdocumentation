/*
Copyright 2020 The Flux authors

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
	"os"

	"github.com/spf13/cobra"
)

var completionFishCmd = &cobra.Command{
	Use:   "fish",
	Short: "Generates fish completion scripts",
	Long:  `The completion sub-command generates completion scripts for fish.`,
	Example: `To configure your fish shell to load completions for each session write this script to your completions dir:

flux completion fish > ~/.config/fish/completions/flux.fish

See http://fishshell.com/docs/current/index.html#completion-own for more details`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenFishCompletion(os.Stdout, true)
	},
}

func init() {
	completionCmd.AddCommand(completionFishCmd)
}
