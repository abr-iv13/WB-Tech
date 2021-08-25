package anagram

import (
	"sort"
	"strings"
)

//Проверка на присутствие строки в слайсе
func isWordInSlice(s string, sl []string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false
}
//Преобразование строки в слайс , сортировка по алфавиту, объединение элементов в единую строку.
func convertStringForCheckUniqueKey (w string) string {
	//Разбитие строки на слайс элементов
	sl := strings.Split(w, "")
	//Сортируем слайс по алфавиту
	sort.Strings(sl)
	//Объединение отсортированных элементов для создания единой строки
	return strings.Join(sl, "")
}

func sortAndClearFirstKeyInMap(tempMap map[string][]string) map[string][]string {
	result := make(map[string][]string)
	for _, w := range tempMap {
		sort.Strings(w)
		//Задать ключом первый элемент салйса
		result[w[0]] = w
	}
	return result
}
//FindSetsAnagram func search sets anagram in slice sting
func FindSetsAnagram (sl []string) map[string][]string {
	tempMap := make(map[string][]string)

	for _, word := range sl {
		//Привести к нижнему регистру
		word = strings.ToLower(word)
		key := convertStringForCheckUniqueKey(word)

		//Проверка на присутсвие ключа в мапе
		if _, ok := tempMap[key]; !ok {
			tempMap[key] = []string{word}
		} else {
			if !isWordInSlice(word, tempMap[key]) {
				tempMap[key] = append(tempMap[key], word)
			}
		}
	}
	//Вернуть отсортированную мапу
	return sortAndClearFirstKeyInMap(tempMap)
}