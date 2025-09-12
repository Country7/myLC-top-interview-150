package main

/* Задача Корректная анаграмма
Даны две строки s и t. Верните true, если t является анаграммой s, и false в противном случае.  */

import (
	"fmt"
)

// ─────────────╮
// Функция проверяет, является ли строка t анаграммой строки s
// Реализация через map, что более универсально, но немного медленнее,
// чем через фиксированный массив на 26 букв
func isAnagramMap(s string, t string) bool {
	// Если длины разные — строки точно не анаграммы
	if len(s) != len(t) {
		return false
	}
	// Создаём словарь (map), где ключ = символ, значение = количество вхождений
	count := make(map[rune]int)
	// Считаем частоту каждой буквы в строке s
	for _, ch := range s {
		count[ch]++
	}
	// Уменьшаем частоту букв, встречающихся в строке t
	for _, ch := range t {
		count[ch]--
		// Если какая-то буква встречается чаще в t, чем в s счётчик станет отрицательным → не анаграмма
		if count[ch] < 0 {
			return false
		}
	}
	// На этом этапе все значения в count должны быть равны 0
	// Но так как мы проверяем отрицательные значения сразу, можно не проходить ещё раз по map
	return true
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Примеры
	fmt.Println("Анаграмма или нет: anagram, nagaram = ", isAnagramMap("anagram", "nagaram")) // true
	fmt.Println("Анаграмма или нет: rat, car = ", isAnagramMap("rat", "car"))                 // false
	// Дополнительные тесты
	fmt.Println("Анаграмма или нет: a, a = ", isAnagramMap("a", "a"))     // true
	fmt.Println("Анаграмма или нет: ab, ba = ", isAnagramMap("ab", "ba")) // true
	fmt.Println("Анаграмма или нет: ab, aa = ", isAnagramMap("ab", "aa")) // false
	fmt.Println()
}
