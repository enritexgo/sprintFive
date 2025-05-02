package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return errors.New("mistake: lenght of data !=2")
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil || steps <= 0 {
		return errors.New("mistake steps count")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(data[1])
	if err != nil || duration <= 0 {
		return errors.New("mistake duration count")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	kkal, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, kkal), nil
}
