package cmd

import (
	"cli-time-tracker/internal/tracker"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().String("name", "", "Gives status of timer by name")
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status",
	Long:  `Show status by timer id (use id as a flag)`,
	Run: func(cmd *cobra.Command, args []string) {

		tracker := tracker.NewTracker()

		err := tracker.Load()
		if err != nil {
			log.Fatal(err)
		}

		name := args[0]

		err = tracker.Status(name)
		if err != nil {
			log.Fatal(err)
		}
	},
}
