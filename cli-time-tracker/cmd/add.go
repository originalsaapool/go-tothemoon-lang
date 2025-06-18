package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
		file, err := os.Create("../config/timers.json")
		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		defer file.Close()

		parentFolderPath := "../config"
		err = os.MkdirAll(parentFolderPath, 0750)
		if err != nil {
			log.Fatal(err)
		}
		// err = os.WriteFile("test/subdir/testfile.txt", []byte("Hello, Gophers!"), 0660)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		type Timer struct {
			ID       string `json:"id"`
			Start    string `json:"start"`
			Duration string `json:"duration"`
			// End      string `json:"end"`
			// ID       string `json: "id"`
		}

		var timers []Timer

		id := fmt.Sprintf("%d", time.Now().UnixNano())
		startTime := time.Now().String()
		duration := strconv.Itoa(1500)

		newTimer := Timer{
			ID:       id,
			Start:    startTime,
			Duration: duration,
			// End: startTime - duration,
		}

		timers = append(timers, newTimer)
		data, err := json.MarshalIndent(timers, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		// обработай err!
		err = os.WriteFile("../config/timers.json", data, 0644)
		if err != nil {
			log.Fatal(err)
		}

		//закрываем файл по окончании функции

		// file.WriteString()

	},
}
