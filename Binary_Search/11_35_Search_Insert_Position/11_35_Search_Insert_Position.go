package main

/* Задача Поиск позиции вставки
Дан отсортированный массив различных целых чисел и искомое значение. Верните индекс, если искомое значение найдено. Если нет, верните индекс, где оно было бы, если бы вставка производилась по порядку.
Необходимо написать алгоритм со сложностью выполнения O(log n). */

import (
	"fmt"
)

// ─────────────╮
// searchInsert — функция для поиска позиции вставки элемента.
// Принимает на вход отсортированный массив nums и целевое значение target.
// Возвращает индекс, по которому либо найден target, либо место, куда его нужно вставить.
func searchInsert(nums []int, target int) int {
	// Задаём левую и правую границы поиска
	left, right := 0, len(nums)-1
	// Используем бинарный поиск, так как требуется O(log n) сложность
	for left <= right {
		// Находим середину текущего диапазона
		mid := left + (right-left)/2
		// Если значение в середине равно target — возвращаем индекс
		if nums[mid] == target {
			return mid
		}
		// Если target больше, чем значение в середине —
		// значит, он может находиться только справа
		if nums[mid] < target {
			left = mid + 1
		} else {
			// Иначе target меньше — он должен быть слева
			right = mid - 1
		}
	}
	// Если target не найден, то левая граница (left)
	// укажет на позицию, куда нужно вставить target
	return left
}

// ─────────────╯

// ─────────────╮
func main() {
	// Примеры использования
	nums1 := []int{1, 3, 5, 6}
	target1 := 5
	fmt.Printf("Input: nums = %v, target = %d\nOutput: %d\n\n", nums1, target1, searchInsert(nums1, target1))
	// Ожидаемый вывод: 2
	nums2 := []int{1, 3, 5, 6}
	target2 := 2
	fmt.Printf("Input: nums = %v, target = %d\nOutput: %d\n\n", nums2, target2, searchInsert(nums2, target2))
	// Ожидаемый вывод: 1
	nums3 := []int{1, 3, 5, 6}
	target3 := 7
	fmt.Printf("Input: nums = %v, target = %d\nOutput: %d\n\n", nums3, target3, searchInsert(nums3, target3))
	// Ожидаемый вывод: 4
	nums4 := []int{1, 3, 5, 6}
	target4 := 0
	fmt.Printf("Input: nums = %v, target = %d\nOutput: %d\n\n", nums4, target4, searchInsert(nums4, target4))
	// Ожидаемый вывод: 0
}
