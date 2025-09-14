package main

/* Задача Самый длинный общий префикс
Напишите функцию для поиска самого длинного общего префикса среди строк в массиве.
Если общего префикса нет, верните пустую строку "".  */

import (
	"fmt"
)

// ─────────────╮
// Функция для нахождения наибольшего общего префикса среди массива строк
func longestCommonPrefix(strs []string) string {
	// Если массив пустой, то общего префикса нет
	if len(strs) == 0 {
		return ""
	}
	// Берем первую строку за основу (потенциальный общий префикс)
	prefix := strs[0]
	// Проходим по всем строкам в массиве начиная со второй
	for i := 1; i < len(strs); i++ {
		// Если текущая строка короче prefix, нужно также подрезать префикс
		for len(prefix) > len(strs[i]) {
			prefix = prefix[:len(prefix)-1]
		}
		// Пока текущая строка не начинается с prefix,
		// мы постепенно укорачиваем prefix на один символ с конца
		for len(prefix) > 0 && len(strs[i]) >= len(prefix) && strs[i][:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}
		// Если префикс стал пустым, значит общего префикса нет
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Пример 1
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Println("Input:", strs1)
	fmt.Println("Output:", longestCommonPrefix(strs1)) // Ожидается "fl"
	// Пример 2
	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println("Input:", strs2)
	fmt.Println("Output:", longestCommonPrefix(strs2)) // Ожидается ""
	strs3 := []string{"flower", "fkow"}
	fmt.Println("Input:", strs3)
	fmt.Println("Output:", longestCommonPrefix(strs3)) // Ожидается ""
	fmt.Println()
}
