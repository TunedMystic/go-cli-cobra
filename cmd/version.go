package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display version",
	Long:  "Display the version of the program",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.0.1")
	},
}
