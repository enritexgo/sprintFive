package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	kkal, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return kkal * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("mistake in the initial parameters")
	}

	speed := MeanSpeed(steps, height, duration)
	durInMin := duration.Minutes()
	kkal := (weight * speed * durInMin) / minInH
	return kkal, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps < 0 || duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	speed := distance / duration.Hours()
	return speed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	step := height * stepLengthCoefficient
	distance := step * float64(steps)
	distanceInKm := distance / mInKm
	return distanceInKm
}
