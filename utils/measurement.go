package utils

import "github.com/inventory-service/model"

func StandarizeMeasurement(value float64, currentUnit string, expectedUnit string) float64 {
	switch expectedUnit {
	case model.Kilogram:
		return toKilogram(value, currentUnit)
	case model.Liter:
		return toLiter(value, currentUnit)
	case model.Gram:
		return toGram(value, currentUnit)
	case model.Mililiter:
		return toMililiter(value, currentUnit)
	}
	return 0.0
}

func toMililiter(value float64, currentUnit string) float64 {
	switch currentUnit {
	case model.Liter:
		return value * 1000
	case model.Mililiter:
		return value

	}
	return 0.0
}

func toLiter(value float64, currentUnit string) float64 {
	switch currentUnit {
	case model.Liter:
		return value
	case model.Mililiter:
		return value * 0.001
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
