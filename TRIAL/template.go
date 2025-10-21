package main

/* Задача
Сложение двоичных чисел
Даны две двоичные строки a и b. Выведите их сумму в виде двоичной строки.
*/

import (
	"fmt"
)

// ─────────────╮
// .
func sumBinary(a, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	maxLen := max(len(a), len(b))
	rezult := make([]byte, maxLen+1)
	k := maxLen

	for i >= 0 || j >= 0 || carry > 0 {
		digitA, digitB := 0, 0
		if i >= 0 {
			digitA = int(a[i] - '0')
		}
		if j >= 0 {
			digitB = int(b[j] - '0')
		}
		sum := digitA + digitB + carry
		rezult[k] = byte(sum%2 + '0')
		carry = sum / 2
		i--
		j--
		k--
	}

	if rezult[0] == '0' || rezult[0] == 0 {
		rezult = rezult[1:]
	}

	return string(rezult)
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	fmt.Println(sumBinary("1010", "1011")) // "10101"
	fmt.Println()
}
