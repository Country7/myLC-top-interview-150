package main

/* Задача Найдите индекс первого вхождения слова «needle» в строке
Данные две строки: «needle» и «haystack». Верните индекс первого вхождения слова «needle» в строке «haystack» или -1, если «needle» не входит в строку «haystack».  */

import (
	"fmt"
)

// ─────────────╮
// Функция ищет первое вхождение подстроки needle в строке haystack.
// Если needle не содержится в haystack, возвращаем -1.
func strStr(haystack string, needle string) int {
	// Если needle пустая строка, по условию задачи это не должно происходить
	// (ограничение: длина >= 1), но на всякий случай можно вернуть 0.
	if len(needle) == 0 {
		return 0
	}
	// Перебираем возможные позиции начала вхождения needle в haystack.
	// Ограничение: мы не должны выходить за границы строки.
	// Последняя возможная позиция — это len(haystack) - len(needle).
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// Проверяем, совпадает ли подстрока haystack[i : i+len(needle)] с needle.
		if haystack[i:i+len(needle)] == needle {
			return i // возвращаем индекс первого совпадения
		}
	}
	// Если совпадений не найдено, возвращаем -1.
	return -1
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Пример 1
	haystack1 := "sadbutsad"
	needle1 := "sad"
	result1 := strStr(haystack1, needle1)
	fmt.Printf("Input: haystack = %q, needle = %q\n", haystack1, needle1)
	fmt.Printf("Output: %d\n", result1) // Ожидается: 0
	// Пример 2
	haystack2 := "leetcode"
	needle2 := "leeto"
	result2 := strStr(haystack2, needle2)
	fmt.Printf("Input: haystack = %q, needle = %q\n", haystack2, needle2)
	fmt.Printf("Output: %d\n", result2) // Ожидается: -1
	fmt.Println()
}
