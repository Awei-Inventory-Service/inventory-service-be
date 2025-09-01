package utils

import "github.com/inventory-service/model"

func StandarizeMeasurement(value float64, currentUnit string, expectedUnit string) float64 {
	switch expectedUnit {
	case model.Kilogram:
		return toKilogram(value, currentUnit)
	case model.Gram:
		return toGram(value, currentUnit)
	}
	return 0.0
}

func toKilogram(value float64, currentUnit string) float64 {
	switch currentUnit {
	case model.Kilogram:
		return value
	case model.Gram:
		return value * 0.001
	}
	return 0.0
}

func toGram(value float64, currentUnit string) float64 {
	switch currentUnit {
	case model.Kilogram:
		return value * 1000
	case model.Gram:
		return value
	}
	return 0.0
}
