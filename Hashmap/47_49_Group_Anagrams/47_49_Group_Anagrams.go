package main

/* Задача Группировка анаграмм
Дан массив строк strs. Сгруппируйте анаграммы. Вы можете вернуть ответ в любом порядке. */

import (
	"fmt"
	"sort"
)

// ─────────────╮
// Основная функция, которая группирует анаграммы
func groupAnagrams(strs []string) [][]string {
	// Используем map, где:
	// ключ -> отсортированная строка (например "aet")
	// значение -> список всех слов, которые являются анаграммами
	anagramMap := make(map[string][]string)
	// Проходим по всем строкам
	for _, str := range strs {
		// Сортируем строку доп. функцией, чтобы получить "ключ"
		sorted := sortString(str)
		// Добавляем исходную строку в список по этому ключу
		anagramMap[sorted] = append(anagramMap[sorted], str)
	}
	// Теперь формируем итоговый результат
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

// Функция для сортировки букв в строке Например, "eat" -> "aet"
func sortString(s string) string {
	// Преобразуем строку в срез рун (или байт, т.к. только латинские буквы)
	runes := []rune(s)
	// Сортируем срез рун в алфавитном порядке
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	// Собираем обратно в строку
	return string(runes)
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Пример 1
	input1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output1 := groupAnagrams(input1)
	fmt.Println("Input:", input1)
	fmt.Println("Output:", output1)
	// Пример 2
	input2 := []string{""}
	output2 := groupAnagrams(input2)
	fmt.Println("\nInput:", input2)
	fmt.Println("Output:", output2)
	// Пример 3
	input3 := []string{"a"}
	output3 := groupAnagrams(input3)
	fmt.Println("\nInput:", input3)
	fmt.Println("Output:", output3)
	fmt.Println()
}
