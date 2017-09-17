package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var destination string
var spaceCrew int

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch spaceship",
	Long:  "Launch the S.S. Galactic and fly off into space!",
	Run: func(cmd *cobra.Command, args []string) {
		// Create message.
		message := launch(destination, spaceCrew)
		// Print message.
		fmt.Println(message)
		// Print args.
		printArgs(args)
	},
}

func init() {
	// Create flags for command.
	launchCmd.Flags().StringVarP(&destination, "destination", "d", "The Moon", "Your destination (can be anywhere in the galaxy!)")
	launchCmd.Flags().IntVarP(&spaceCrew, "crew", "c", 7, "The amount of people in your trusty space crew.")
}

func launch(destination string, spaceCrew int) string {
	message := `As you set course for %v, you and ` +
		`your trusty space crew of %v brave ` +
		`souls launch the S.S. Galactic ` +
		"to the stars!\n" +
		`Zoom!`
	return fmt.Sprintf(message, destination, spaceCrew)
}
