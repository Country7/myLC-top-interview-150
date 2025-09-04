package main

/* Задача Сложение двоичных чисел
Даны две двоичные строки a и b. Выведите их сумму в виде двоичной строки.  */

import (
	"fmt"
)

// ─────────────╮
// Функция принимает две двоичные строки (a и b) и возвращает их сумму также в виде двоичной строки
// Улучшенная версия без реверса строки
func addBinary(a string, b string) string {
	// Индексы для обхода строк справа налево
	i := len(a) - 1
	j := len(b) - 1
	// перенос, который возникает при сложении (0 или 1)
	carry := 0
	// Максимальная длина результата — на 1 больше, чем max(len(a), len(b))
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	// Байтовый срез для формирования результата
	result := make([]byte, maxLen+1)
	// индекс, с которого будем заполнять результат справа налево
	k := maxLen
	fmt.Println("\na=", a, ",  b=", b, ",  maxLen=", maxLen, ",  k= maxLen =", k)

	// Цикл выполняется пока есть символы в любой из строк или остался перенос
	for i >= 0 || j >= 0 || carry > 0 {
		digitA := 0
		if i >= 0 {
			digitA = int(a[i] - '0') // Преобразуем символ в число
		}
		digitB := 0
		if j >= 0 {
			digitB = int(b[j] - '0') // Преобразуем символ в число
		}
		sum := digitA + digitB + carry
		fmt.Println("    ╰──> sum = digitA", digitA, "+ digitB", digitB, "+ carry", carry, "=", sum)
		// Пишем цифру сразу на нужное место
		result[k] = byte(sum%2 + '0') // Оператор % в Go возвращает остаток от деления целых чисел
		fmt.Println("    ╰──> result[k", k, "] = byte(sum", sum, "% 2 + '0') =", string(result[k]))

		// перенос, который возникает при сложении (0 или 1)
		carry = sum / 2
		fmt.Println("      ╰──> carry = sum", sum, " / 2  =", carry, "\n ")
		i--
		j--
		k--
	}

	// Если результат начинается с '0', значит перенос не понадобился — убираем лишний ноль
	if result[0] == '0' {
		return string(result[1:])
	}
	return string(result)
}

// ─────────────╯

// ─────────────╮╰──>

func main() {
	fmt.Println()
	fmt.Println(addBinary("11", "1"))       // "100"
	fmt.Println(addBinary("1010", "1011"))  // "10101"
	fmt.Println(addBinary("1111", "11111")) // "101110"
	fmt.Println()
}
