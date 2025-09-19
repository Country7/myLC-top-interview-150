package main

/* Задача Разбиение на слова
Для строки s и словаря строк wordDict верните значение true, если s можно сегментировать в последовательность из одного или нескольких слов словаря, разделённых пробелами.
Обратите внимание, что одно и то же слово из словаря может использоваться в сегментации несколько раз. */

import (
	"fmt"
)

// ─────────────╮
// Функция проверяет, можно ли строку s разбить на слова из wordDict
func wordBreak(s string, wordDict []string) bool {
	// Создаем множество для словаря для быстрой проверки наличия слова
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	// dp[i] означает, что s[0:i] можно разбить на слова из словаря
	dp := make([]bool, len(s)+1)
	dp[0] = true // пустая строка всегда может быть разбита
	// Проходим по всей строке
	for i := 1; i <= len(s); i++ {
		// Проверяем все возможные разбиения s[0:i]
		for j := 0; j < i; j++ {
			// Если s[0:j] можно разбить (dp[j] == true)
			// и s[j:i] находится в словаре
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break // нашли подходящее разбиение, дальше проверять не нужно
			}
		}
	}
	// Возвращаем результат для всей строки s
	return dp[len(s)]
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Примеры использования
	s1 := "leetcode"
	wordDict1 := []string{"leet", "code"}
	fmt.Println(wordBreak(s1, wordDict1)) // true
	s2 := "applepenapple"
	wordDict2 := []string{"apple", "pen"}
	fmt.Println(wordBreak(s2, wordDict2)) // true
	s3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s3, wordDict3)) // false
	fmt.Println()
}
