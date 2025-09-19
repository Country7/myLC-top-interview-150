package main

/* Задача Буквосочетания в телефонном номере
Дана строка, содержащая цифры от 2 до 9 включительно. Верните все возможные буквосочетанья, которые может представлять это число. Верните ответ в любом порядке.
Ниже приведено соответствие цифр буквам (как на кнопках телефона). Обратите внимание, что 1 не соответствует ни одной букве. */

import (
	"fmt"
)

// ─────────────╮
// Функция принимает строку digits (цифры от '2' до '9')
// и возвращает все возможные комбинации букв, соответствующие этим цифрам
func letterCombinations(digits string) []string {
	// Если входная строка пустая — сразу возвращаем пустой список
	if len(digits) == 0 {
		return []string{}
	}
	// Создаём словарь (map), который отображает цифру в соответствующий набор букв - аналог кнопок на телефоне
	digitToLetters := map[byte]string{
		'0': "",
		'1': "",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	// Результат — список строк
	var result []string
	// Рекурсивная функция backtracking для построения всех комбинаций
	var backtrack func(index int, current string)
	backtrack = func(index int, current string) {
		// Если текущая длина комбинации равна длине digits, значит, мы построили готовую комбинацию
		if index == len(digits) {
			result = append(result, current)
			return
		}
		// Берём цифру по индексу
		digit := digits[index]
		// Получаем все возможные буквы для этой цифры
		letters := digitToLetters[digit]
		// Перебираем все буквы и строим комбинации рекурсивно
		for i := 0; i < len(letters); i++ {
			// Добавляем букву к текущей комбинации и двигаемся дальше
			backtrack(index+1, current+string(letters[i]))
		}
		fmt.Println("  ╰──> ", result)
	}
	// Запускаем рекурсию с первой цифры и пустой строкой
	backtrack(0, "")
	// Возвращаем результат
	return result
}

// Итеративное решение задачи Letter Combinations of a Phone Number
func letterCombinationsIter(digits string) []string {
	// Если входная строка пустая — сразу возвращаем пустой список
	if len(digits) == 0 {
		return []string{}
	}
	// Отображение цифр в буквы (аналог кнопок телефона)
	digitToLetters := map[byte]string{
		'0': "",
		'1': "",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	// Начинаем с одного "пустого" варианта
	result := []string{""}
	// Проходим по каждой цифре из входной строки
	for i := 0; i < len(digits); i++ {
		// Берём текущую цифру
		digit := digits[i]
		// Получаем все буквы для этой цифры
		letters := digitToLetters[digit]
		// Временный список для хранения новых комбинаций
		var newResult []string
		// Для каждой комбинации, которую уже построили ранее,
		// добавляем каждую букву, связанную с текущей цифрой
		for _, combination := range result {
			for j := 0; j < len(letters); j++ {
				newResult = append(newResult, combination+string(letters[j]))
			}
		}
		// Обновляем result, чтобы на следующем шаге использовать новые комбинации
		result = newResult
		fmt.Println("  ╰──> ", result)
	}
	return result
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	examples := []string{"23", "", "2", "987"}
	fmt.Println("\nРекурсивно \n ")
	for _, str := range examples {
		fmt.Println("Номер телефона:", str)
		fmt.Println("  => Результат:", letterCombinations(str))
	}
	fmt.Println("\nИтеративно \n ")
	for _, str := range examples {
		fmt.Println("Номер телефона:", str)
		fmt.Println("  => Результат:", letterCombinationsIter(str))
	}
	fmt.Println()
}
