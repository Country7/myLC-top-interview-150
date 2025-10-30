package main

/* Задача Содержит дубликаты II
Дан массив целых чисел nums и целое число k. Вернуть true, если в массиве есть два различных индекса i и j, таких, что nums[i] == nums[j] и abs(i - j) <= k. (Абсолютное значение числа (модуль числа) — это его величина без учёта знака.) */

import (
	"fmt"
)

// ─────────────╮
// Функция проверяет, есть ли в массиве два одинаковых элемента, расстояние между которыми не больше k.
func containsNearbyDuplicate(nums []int, k int) bool {
	// Используем map для хранения последнего индекса каждого числа
	lastSeen := make(map[int]int)
	// Проходим по всем элементам массива
	for i, num := range nums {
		// Проверяем, встречалось ли число ранее
		if prevIndex, exists := lastSeen[num]; exists {
			// Если да, проверяем условие на расстояние
			if i-prevIndex <= k {
				return true // нашли два одинаковых числа на допустимом расстоянии
			}
		}
		// Обновляем индекс последнего вхождения числа
		lastSeen[num] = i
	}
	// Если не нашли подходящих пар, возвращаем false
	return false
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Примеры
	nums1 := []int{1, 2, 3, 1}
	k1 := 3
	fmt.Println(containsNearbyDuplicate(nums1, k1)) // true
	nums2 := []int{1, 0, 1, 1}
	k2 := 1
	fmt.Println(containsNearbyDuplicate(nums2, k2)) // true
	nums3 := []int{1, 2, 3, 1, 2, 3}
	k3 := 2
	fmt.Println(containsNearbyDuplicate(nums3, k3)) // false
	fmt.Println()
}
