package main

/* Задача - Корректный палиндром
Фраза является палиндромом, если после преобразования всех заглавных букв в строчные и удаления всех символов, кроме букв и цифр, она читается одинаково слева направо и слева направо. К буквенно-цифровым символам относятся буквы и цифры.

В задаче есть ограничения:
1 <= s.length <= 2 * 105
s состоит только из печатных символов ASCII

В учебных целях добавлена функция UNICODE (на будущее) */

import (
	"fmt"
	"strings"
	"unicode"
)

// Ограничение по длине s
const MaxLength = 2 * 100000

// Основная функция проверки палидрома в ASCII
func isPalindromeASCII(s string) bool {
	// Проверка длины строки
	if len(s) < 1 || len(s) > MaxLength {
		panic("s.length out of allowed range")
	}

	left, right := 0, len(s)-1

	for left < right {
		// Пропускаем неалфанумерные символы слева
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		// Пропускаем неалфанумерные символы справа
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		// Сравниваем символы в нижнем регистре
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

// Проверка, является ли ASCII-символ буквой или цифрой
func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

// Перевод ASCII-буквы в нижний регистр
func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

// Основная функция проверки палидрома в UNICODE
func isPalindromeUNICODE(s string) bool {
	var builder strings.Builder
	/* Почему нужен strings.Builder?
	В Go строки иммутабельны (неизменяемы). Это значит, что при каждом конкатенировании (str = str + "x") создаётся новая строка, и все символы копируются. Если так делать в цикле — это будет очень неэффективно (O(n²) по времени).
	strings.Builder решает эту проблему: внутри себя он хранит буфер (срез байт), при добавлении новых символов он просто дописывает их в буфер, и только в конце мы получаем итоговую строку методом .String().
	Основные методы:
	var b strings.Builder — создание билдера.
	b.WriteRune(r) — записывает руну (символ).
	b.WriteString(str) — записывает строку.
	b.WriteByte(c) — записывает один байт.
	b.String() — возвращает собранную строку. */

	/* Очищаем строку: оставляем только буквы и цифры, приводим к нижнему регистру
	unicode.IsLetter(r rune) bool в Go проверяет, является ли руна (символ) буквой, IsDigit цифрой */
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			builder.WriteRune(unicode.ToLower(ch))
		}
	}
	cleaned := builder.String()
	n := len(cleaned)
	// Проверяем палиндром двумя указателями
	for i := 0; i < n/2; i++ {
		if cleaned[i] != cleaned[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	tests := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		" ",
	}

	fmt.Println("\nПроверяем палидром в ASCII:")
	for _, t := range tests {
		fmt.Printf("Input: %q -> %v\n", t, isPalindromeASCII(t))
	}

	fmt.Println("\nПроверяем палидром в UNICODE:")
	for _, t := range tests {
		fmt.Printf("Input: %q -> %v\n", t, isPalindromeUNICODE(t))
	}
	fmt.Println()
}
