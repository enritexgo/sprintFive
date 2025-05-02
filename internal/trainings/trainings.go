package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		err = errors.New("mistake: lenght of data != 3")
		return
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return
	}
	if steps <= 0 {
		return errors.New("mistake steps count")
	}
	duration, err := time.ParseDuration(data[2])
	if err != nil {
		return
	}
	if duration <= 0 {
		return errors.New("mistake duration count")
	}

	t.Steps = steps
	t.TrainingType = data[1]
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var (
		kkal   float64
		result string
		err    error
	)

	typeRun, typeWalk := "Бег", "Ходьба"
	switch t.TrainingType {
	case typeRun:
		kkal, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case typeWalk:
		kkal, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("mistake training type")
	}
	if err != nil {
		return "", err
	}

	result = fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, kkal)

	return result, nil
}
