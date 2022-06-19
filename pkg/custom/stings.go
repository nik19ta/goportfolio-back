package custom

import (
	"strings"
)

func GetArray(str string) []string {
	arr := strings.Split(str, ",")
	return arr
}

func GetString(str []string) string {
	return strings.Join(str[:], ",")
}

// Поменять местами
func ConfuseData(start, finish int, str string) string {
	temp := ""
	array := strings.Split(str, ",")

	temp = array[finish]

	array[finish] = array[start]
	array[start] = temp

	return strings.Join(array[:], ",")
}

// Вставить новый элемент
func InsertIntoString(str, new string) string {
	if len(str) > 0 {
		array := strings.Split(str, ",")

		array = append(array, new)
		return strings.Join(array[:], ",")
	} else {
		return new
	}
}

// Удалить элемент из строки
func DeleteFromString(str, id string) string {
	array := strings.Split(str, ",")

	var temp_array []string

	for _, item := range array {
		if strings.TrimSpace(item) != id {
			temp_array = append(temp_array, item)
		}
	}

	return strings.Join(temp_array[:], ",")
}
