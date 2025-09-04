package main

/* Задача Число-палиндром
Для заданного целого числа x верните true, если x является палиндромом, и false в противном случае.
Ограничения: -2^31 <= x <= 2^31 - 1
Дополнительно: Можете ли вы решить эту задачу, не преобразуя целое число в строку? */

import (
	"fmt"
)

// ─────────────╮
// Функция проверяет, является ли число палиндромом (самый быстрый вариант преобразуя в строку)
func isPalindromeStr(x int) bool {
	fmt.Printf("Число x= %d  ", x)
	s := fmt.Sprintf("%d", x)
	length := len(s)
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}

// Функция проверяет, является ли число палиндромом (не преобразуя в строку, медленнее)
func isPalindrome(x int) bool {
	fmt.Println("Число x=", x)
	// Отрицательные числа не могут быть палиндромами
	// (например, -121 ≠ 121-)
	if x < 0 {
		return false
	}
	// Если число оканчивается на 0 (кроме самого числа 0),
	// оно не может быть палиндромом
	// (например, 10 → справа налево будет 01)
	if x%10 == 0 && x != 0 {
		return false
	}
	// reversed хранит перевёрнутую "вторую половину" числа
	reversed := 0
	// Пока оригинальное число больше перевёрнутой половины —
	// продолжаем переносить цифры
	for x > reversed {
		// Берём последнюю цифру числа
		lastDigit := x % 10
		// Добавляем её к перевёрнутой половине
		reversed = reversed*10 + lastDigit
		// Убираем последнюю цифру из числа
		x /= 10
		fmt.Println("   ╰──> reversed=", reversed, ",  x=", x)
	}
	// В этот момент у нас два варианта:
	// 1) x == reversed → чётное количество цифр (например, 1221 → 12 и 12)
	// 2) x == reversed/10 → нечётное количество цифр (например, 12321 → 12 и 123/10=12)
	return x == reversed || x == reversed/10
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Примеры проверки
	fmt.Println("Результат:", isPalindrome(121), "\n ")   // true
	fmt.Println("Результат:", isPalindrome(-121), "\n ")  // false
	fmt.Println("Результат:", isPalindrome(10), "\n ")    // false
	fmt.Println("Результат:", isPalindrome(12321), "\n ") // true
	fmt.Println("Результат:", isPalindrome(0), "\n ")     // true
	fmt.Println("Теперь тоже самое, но преобразуя число в строку: \n ")
	fmt.Println("Результат:", isPalindromeStr(121), "\n ")   // true
	fmt.Println("Результат:", isPalindromeStr(-121), "\n ")  // false
	fmt.Println("Результат:", isPalindromeStr(10), "\n ")    // false
	fmt.Println("Результат:", isPalindromeStr(12321), "\n ") // true
	fmt.Println("Результат:", isPalindromeStr(0), "\n ")     // true
}
