package main

/* Задача 383. Записка с требованием выкупа
Даны две строки: ransomNote и magazine. Верните true, если ransomNote можно составить из букв слова magazine, и false в противном случае.
Каждая буква из слова magazine может быть использована в слове ransomNote только один раз.
Ограничения:
1 <= ransomNote.length, magazine.length <= 10^5
Слова ransomNote и magazine состоят из строчных английских букв. */

import (
	"fmt"
)

// ─────────────╮
// canConstruct проверяет, можно ли составить ransomNote из букв magazine
func canConstruct(ransomNote string, magazine string) bool {
	// Проверяем ограничения Constraints
	if len(ransomNote) < 1 || len(ransomNote) > 100000 {
		fmt.Println("Error: ransomNote length must be between 1 and 105")
		return false
	}
	if len(magazine) < 1 || len(magazine) > 100000 {
		fmt.Println("Error: magazine length must be between 1 and 105")
		return false
	}
	// Создаем карту для подсчета количества букв в magazine
	magCount := make(map[rune]int)
	for _, ch := range magazine {
		magCount[ch]++
	}
	// Проверяем, достаточно ли каждой буквы из ransomNote
	for _, ch := range ransomNote {
		if magCount[ch] > 0 {
			// Если буква есть в magazine, уменьшаем счетчик
			magCount[ch]--
		} else {
			// Если буквы нет или она закончилась, возвращаем false
			return false
		}
	}
	// Если все буквы ransomNote найдены в magazine
	return true
}

// ─────────────╯

// ─────────────╮
func main() {
	fmt.Println(canConstruct("a", "b"))    // false
	fmt.Println(canConstruct("aa", "ab"))  // false
	fmt.Println(canConstruct("aa", "aab")) // true
}
