package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"cli-time-tracker/internal/domain"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().String("id", "", "Gives status of timer by id")
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show status",
	Long:  `Show status by timer id (use id as a flag)`,
	Run: func(cmd *cobra.Command, args []string) {

		parentFolderPath := "config"
		err := os.MkdirAll(parentFolderPath, 0750)
		if err != nil {
			log.Fatal(err)
		}

		var timers []domain.Timer

		file, err := os.ReadFile(parentFolderPath + "/timers.json")
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(file, &timers)
		if err != nil {
			log.Fatal(err)
		}

		id := args[0]

		// fmt.Printf("%+v\n", timers[0])
		// fmt.Println(id)

		for _, v := range timers {
			if v.ID == id {
				fmt.Printf("The timer %s is now %s\n", v.ID, v.Status)
				startTime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", v.Start)
				if err != nil {
					log.Fatal(err)
				}
				now := time.Now()
				diff := now.Sub(startTime).Round(time.Second).String()
				fmt.Printf("%v from the beggining\n", diff)
				// endTime := startTime.Add(time.Duration(duration.Duration) * time.Second)
			}
		}

		data, err := json.MarshalIndent(timers, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(parentFolderPath+"/timers.json", data, 0644)
		if err != nil {
			log.Fatal(err)
		}
		//закрываем файл по окончании функции

		// file.WriteString()

	},
}
