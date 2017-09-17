package cmd

import "github.com/spf13/cobra"

// RootCmd command
var RootCmd = &cobra.Command{
	Use:   "spaceship",
	Short: "The short spaceship command",
	Long:  "The long spaceship command!",
	Run: func(cmd *cobra.Command, args []string) {
		// Print help.
		cmd.Help()
	},
}

func addCommands() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(launchCmd)
}

// Execute function runs the root command.
func Execute() error {
	addCommands()
	return RootCmd.Execute()
}
