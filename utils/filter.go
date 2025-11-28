package utils

import "github.com/inventory-service/dto"

func CheckKeyExist(key string, filters []dto.Filter) (bool, []string) {
	for _, filter := range filters {
		if filter.Key == key {
			return true, filter.Values
		}
	}
	return false, []string{}
}

func CheckKeyExistWithDefaultValue(key string, filters []dto.Filter, defaultValue any) (bool, any) {
	if exist, values := CheckKeyExist(key, filters); exist {
		return true, values
	}
	return false, defaultValue
}
