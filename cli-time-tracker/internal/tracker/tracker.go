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
	Timers      map[string]*domain.Timer
	StoragePath string
}

func NewTracker(filepath string) *Tracker {
	return &Tracker{
		StoragePath: filepath,
	}
}

func (t *Tracker) Load() error {

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

		err = t.UnmarshalJSON(file)
		if err != nil {
			return err
		}
	}

	return nil

}

func (t *Tracker) UnmarshalJSON(data []byte) error {
	// создаем переменную - массив из структуры типа Timer (из domain)
	var timers []domain.Timer
	// получаем из файла и декодируем таймеры
	if err := json.Unmarshal(data, &timers); err != nil {
		return err
	}
	// делаем мапу где ключ имя а значение структура типа Timer
	t.Timers = make(map[string]*domain.Timer)
	// проходим по переменной которая содержим струкруы и переписываем данные в мапу чтобы дальше с ней работать
	for i := range timers {
		t.Timers[timers[i].Name] = &timers[i]
	}
	// теперь таймеры хранятся в мапе
	return nil
}

func (t *Tracker) MarshalJSON() ([]byte, error) {
	// после того как набор таймеров изменен - вытаскиваем данные в срез из мапы
	timers := make([]domain.Timer, 0, len(t.Timers))
	for _, tmr := range t.Timers {
		timers = append(timers, *tmr)
	}
	//кодируем полученный срез в json
	return json.Marshal(timers)
}

func (t *Tracker) Save() error {
	// data, err := json.MarshalIndent(t.Timers, "", "  ")
	data, err := t.MarshalJSON()
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

	// t.Timers = append(t.Timers, newTimer)
	t.Timers[newTimer.Name] = &newTimer

	fmt.Println("The timer has started!")

}

func (t *Tracker) Stop(name string) {

	t.Timers[name].Status = "stopped"
	fmt.Printf("The timer %s has been stopped!\n", name)
}

func (t *Tracker) Status(name string) {

	_, ok := t.Timers[name]
	if ok {
		fmt.Printf("The timer %s is now %s\n", t.Timers[name].Name, t.Timers[name].Status)
		if t.Timers[name].Status == "running" {
			startTime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", t.Timers[name].Start)
			if err != nil {
				log.Fatal(err)
			}
			now := time.Now()
			diff := now.Sub(startTime).Round(time.Second).String()
			fmt.Printf("%v from the beggining\n", diff)
		}
	} else {
		fmt.Println("The timer is not found")
	}
}
