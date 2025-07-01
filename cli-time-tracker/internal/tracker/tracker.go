package tracker

import (
	"cli-time-tracker/internal/domain"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Tracker struct {
	Timers      map[string]*domain.Timer
	StoragePath string
}

func NewTracker() *Tracker {
	return &Tracker{
		Timers:      make(map[string]*domain.Timer),
		StoragePath: "config/timers.json",
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

		data, err := json.Marshal("[]")
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

func (t *Tracker) AddTimer(name string, minutes string) error {

	duration := domain.TimerDuration{}
	dur, err := time.ParseDuration(minutes + "m")
	if err != nil {
		return fmt.Errorf("duration parsing error")
	}

	duration.Duration = dur

	// duration.Number, err = strconv.Atoi(duration.String)
	// if err != nil {
	// 	panic(err)
	// }
	// duration.Number = duration.Number * 60
	// duration.String = strconv.Itoa(duration.Number)
	// duration.Duration = time.Duration(duration.Number)

	id := fmt.Sprintf("%d", time.Now().UnixNano())
	startTime := time.Now()
	endTime := startTime.Add(duration.Duration)

	newTimer := domain.Timer{
		ID:        id,
		Name:      name,
		Duration:  duration.Duration,
		StartedAt: startTime,
		StoppedAt: endTime,
	}

	// t.Timers = append(t.Timers, newTimer)
	t.Timers[newTimer.Name] = &newTimer

	fmt.Println("The timer has started!")

	return nil
}

func (t *Tracker) Stop(name string) error {

	_, ok := t.Timers[name]
	if ok {
		nowTime := time.Now()
		if nowTime.Before(t.Timers[name].StoppedAt) {
			t.Timers[name].StoppedAt = nowTime
			fmt.Printf("The timer %s has been stopped!\n", name)
			return nil
		} else {
			return fmt.Errorf("the timer is already stopped")
		}
	} else {
		return fmt.Errorf("the timer is not found")
	}
}

func (t *Tracker) Status(name string) error {

	_, ok := t.Timers[name]
	if ok {
		nowTime := time.Now()
		if nowTime.Before(t.Timers[name].StoppedAt) {
			fmt.Printf("The timer %s is now running\n", t.Timers[name].Name)
			timeLeft := t.Timers[name].StoppedAt.Sub(nowTime)
			fmt.Printf("Осталось до конца таймера: %s\n", timeLeft)
			elapsed := nowTime.Sub(t.Timers[name].StartedAt)
			fmt.Printf("Прошло с запуска таймера: %s\n", elapsed)
			return nil
		} else {
			return fmt.Errorf("the timer is already stopped")
		}

	} else {
		return fmt.Errorf("the timer is not found")
	}
}
