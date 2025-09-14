package main

/* Задача

 */

import (
	"fmt"
)

// ─────────────╮
// Функция для поворота массива nums вправо на k шагов
func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 || k%n == 0 {
		// Если массив пустой или k кратно длине массива, ничего делать не нужно
		return
	}
	k = k % n // если k больше длины массива, берем остаток от деления
	// Вспомогательная функция для реверса подмассива с индекса start до end
	reverse := func(nums []int, start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	// Алгоритм "три реверса":
	// 1. Реверсируем весь массив
	reverse(nums, 0, n-1)
	// 2. Реверсируем первые k элементов
	reverse(nums, 0, k-1)
	// 3. Реверсируем оставшиеся n-k элементов
	reverse(nums, k, n-1)
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Пример 1
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	k1 := 3
	rotate(nums1, k1)
	fmt.Println(nums1) // [5 6 7 1 2 3 4]
	// Пример 2
	nums2 := []int{-1, -100, 3, 99}
	k2 := 2
	rotate(nums2, k2)
	fmt.Println(nums2) // [3 99 -1 -100]
	fmt.Println()
}
