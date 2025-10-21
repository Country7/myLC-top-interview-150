package main

/* Задача Длина последнего слова
Дана строка s, состоящая из слов и пробелов. Верните длину последнего слова в строке.
Слово — это максимальная подстрока, состоящая только из символов, не являющихся пробелами. */

import (
	"fmt"
	"strings"
)

// ─────────────╮
// Функция принимает строку s и возвращает длину последнего слова
// В этой задаче нужно запомнить s = strings.TrimSpace(s) и strings.LastIndex(s, " ")
func lengthOfLastWord(s string) int {
	// Убираем пробелы в начале и в конце строки с помощью функции strings.TrimSpace
	// Это важно, так как в условии могут быть лишние пробелы
	s = strings.TrimSpace(s)
	// Ищем индекс последнего пробела в строке
	lastSpaceIndex := strings.LastIndex(s, " ")
	// Если пробела нет, значит вся строка — это одно слово
	if lastSpaceIndex == -1 {
		return len(s)
	}
	// Последнее слово — это подстрока после последнего пробела
	lastWord := s[lastSpaceIndex+1:]
	// Возвращаем длину последнего слова
	return len(lastWord)
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Пример 1
	s1 := "Hello World"
	fmt.Printf("Input: %q → Output: %d\n", s1, lengthOfLastWord(s1))
	// Пример 2
	s2 := "   fly me   to   the moon  "
	fmt.Printf("Input: %q → Output: %d\n", s2, lengthOfLastWord(s2))
	// Пример 3
	s3 := "luffy is still joyboy"
	fmt.Printf("Input: %q → Output: %d\n", s3, lengthOfLastWord(s3))
	fmt.Println()
}
