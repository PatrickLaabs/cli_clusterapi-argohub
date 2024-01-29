/*
Copyright © 2024 Patrick Laabs <patrick.laabs@me.com>
*/
package k3d

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Test represents the k3d/test command
var Test = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("k3d/test called")
	},
}
