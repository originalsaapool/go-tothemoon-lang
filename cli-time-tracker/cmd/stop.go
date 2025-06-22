package cmd

import (
	"log"

	"cli-time-tracker/internal/tracker"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stopCmd)
	stopCmd.Flags().String("name", "", "Stops timer using exact name")

}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop",
	Long:  `You can stop timer using this command`,
	Run: func(cmd *cobra.Command, args []string) {

		tracker := tracker.NewTracker("config/timers.json")

		err := tracker.Load()
		if err != nil {
			log.Fatal(err)
		}

		name := args[0]

		tracker.Stop(name)

		err = tracker.Save()
		if err != nil {
			log.Fatal(err)
		}
	},
}
