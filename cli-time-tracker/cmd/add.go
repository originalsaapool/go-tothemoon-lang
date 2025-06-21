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

	"cli-time-tracker/internal/domain"

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

		// fmt.Println(os.Getwd())

		parentFolderPath := "config"
		dataFile := "timers.json"
		err := os.MkdirAll(parentFolderPath, 0750)
		if err != nil {
			log.Fatal(err)
		}

		_, err = os.Stat(parentFolderPath + "/" + dataFile)
		if os.IsNotExist(err) {
			file, err := os.Create(parentFolderPath + "/" + dataFile)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
		} else if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Set time for the task im minutes: ")
		result, _ := reader.ReadString('\n')
		result = strings.TrimSpace(result)

		duration := domain.TimerDuration{String: result}

		duration.Number, err = strconv.Atoi(duration.String)
		if err != nil {
			panic(err)
		}

		duration.Number = duration.Number * 60
		duration.String = strconv.Itoa(duration.Number)
		duration.Duration = time.Duration(duration.Number)

		var timers []domain.Timer

		id := fmt.Sprintf("%d", time.Now().UnixNano())
		startTime := time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST")
		// endTime := startTime.Add(time.Duration(duration.Duration) * time.Second)

		file, err := os.ReadFile(parentFolderPath + "/" + dataFile)
		if err != nil {
			log.Fatal(err)
		}

		//здесь нужно обработать по другому
		err = json.Unmarshal(file, &timers)
		if err != nil {
			fmt.Println("Ошибка десериализации")
			log.Fatal(err)
		}

		newTimer := domain.Timer{
			ID:       id,
			Start:    startTime,
			Duration: duration.String,
			Status:   "running",
		}

		timers = append(timers, newTimer)

		data, err := json.MarshalIndent(timers, "", "  ")
		if err != nil {
			fmt.Println("Ошибка сериализации")
			log.Fatal(err)
		}

		err = os.WriteFile(parentFolderPath+"/"+dataFile, data, 0644)
		if err != nil {
			fmt.Println("Ошибка записи файла")
			log.Fatal(err)
		}
		fmt.Println("The timer has started!")
		//закрываем файл по окончании функции

		// file.WriteString()

	},
}
