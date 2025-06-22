package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"cli-time-tracker/internal/tracker"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add timer",
	Long:  `You can add timer using this command`,
	Run: func(cmd *cobra.Command, args []string) {

		tracker := tracker.NewTracker("config/timers.json")

		err := tracker.Load()
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Set time for the task im minutes: ")
		result, _ := reader.ReadString('\n')
		minutes := strings.TrimSpace(result)

		fmt.Println("Set name for the timer: ")
		result, _ = reader.ReadString('\n')
		name := strings.TrimSpace(result)

		tracker.AddTimer(name, minutes)

		err = tracker.Save()
		if err != nil {
			log.Fatal(err)
		}
	},
}
