package tracker

import (
	"cli-time-tracker/internal/domain"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Tracker struct {
	Timers      []domain.Timer
	StoragePath string
}

func NewTracker(filepath string) *Tracker {
	return &Tracker{
		StoragePath: filepath,
	}
}

func (t *Tracker) Load() error {

	// parentFolderPath := "config"
	// dataFile := "timers.json"
	err := os.MkdirAll("config", 0750)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(t.StoragePath)
	if os.IsNotExist(err) {

		file, err := os.Create(t.StoragePath)
		if err != nil {
			return err
		}

		data, err := json.MarshalIndent("[]", "", "  ")
		if err != nil {
			fmt.Println("Ошибка сериализации")
			return err
		}

		err = os.WriteFile(t.StoragePath, data, 0644)
		if err != nil {
			fmt.Println("Ошибка записи файла")
			return err
		}

		defer file.Close()

	} else if err != nil {
		return err
	} else {

		file, err := os.ReadFile(t.StoragePath)
		if err != nil {
			return err
		}

		err = json.Unmarshal(file, &t.Timers)
		if err != nil {
			return err
		}
	}

	return nil

}

func (t *Tracker) Save() error {
	data, err := json.MarshalIndent(t.Timers, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(t.StoragePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tracker) AddTimer(name string, minutes string) {

	duration := domain.TimerDuration{String: minutes}

	var err error

	duration.Number, err = strconv.Atoi(duration.String)
	if err != nil {
		panic(err)
	}
	duration.Number = duration.Number * 60
	duration.String = strconv.Itoa(duration.Number)
	duration.Duration = time.Duration(duration.Number)

	id := fmt.Sprintf("%d", time.Now().UnixNano())
	startTime := time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST")

	newTimer := domain.Timer{
		ID:       id,
		Start:    startTime,
		Duration: duration.String,
		Status:   "running",
		Name:     name,
	}

	t.Timers = append(t.Timers, newTimer)

	fmt.Println("The timer has started!")

}

func (t *Tracker) Stop(name string) {
	for i, v := range t.Timers {
		if v.Name == name {
			t.Timers[i].Status = "stopped"
			fmt.Printf("The timer %s has been stopped!\n", name)
		}
	}
}

func (t *Tracker) Status(name string) {
	for _, v := range t.Timers {
		if v.Name == name {
			fmt.Printf("The timer %s is now %s\n", v.Name, v.Status)
			if v.Status == "running" {

				startTime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", v.Start)
				if err != nil {
					log.Fatal(err)
				}
				now := time.Now()
				diff := now.Sub(startTime).Round(time.Second).String()
				fmt.Printf("%v from the beggining\n", diff)
			}
		}
	}
}
