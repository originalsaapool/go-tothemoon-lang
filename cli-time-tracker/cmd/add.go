package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

		fmt.Println(os.Getwd())

		parentFolderPath := "config"
		err := os.MkdirAll(parentFolderPath, 0750)
		if err != nil {
			log.Fatal(err)
		}

		type Timer struct {
			ID       string        `json:"id"`
			Start    time.Time     `json:"start"`
			Duration time.Duration `json:"duration"`
			End      time.Time     `json:"end"`
			// ID       string `json: "id"`
		}

		type TimerDuration struct {
			Number   int
			String   string
			Duration time.Duration
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Set time for the task im minutes: ")
		result, _ := reader.ReadString('\n')
		result = strings.TrimSpace(result)

		duration := TimerDuration{String: result}

		duration.Number, err = strconv.Atoi(duration.String)
		if err != nil {
			panic(err)
		}

		duration.Number = duration.Number * 60
		duration.Duration = time.Duration(duration.Number)

		var timers []Timer

		id := fmt.Sprintf("%d", time.Now().UnixNano())
		startTime := time.Now()
		endTime := startTime.Add(time.Duration(duration.Duration) * time.Second)

		fmt.Printf("%T", startTime)

		// duration := strconv.Itoa(1500)

		file, err := os.ReadFile(parentFolderPath + "/timers.json")
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(file, &timers)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Содержимое файла:\n", string(file))
		fmt.Println(timers)

		newTimer := Timer{
			ID:       id,
			Start:    startTime,
			Duration: duration.Duration,
			End:      endTime,
		}

		timers = append(timers, newTimer)

		data, err := json.MarshalIndent(timers, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(parentFolderPath+"/timers.json", data, 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("The timer has started!")
		//закрываем файл по окончании функции

		// file.WriteString()

	},
}
