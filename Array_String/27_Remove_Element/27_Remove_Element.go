package main

import "fmt"

/*
Дан массив целых чисел nums и целое число val. Удалить все вхождения val в массиве nums без изменения. Порядок элементов может быть изменён. Затем вернуть количество элементов в массиве nums, которые не равны val.

Считая, что количество элементов в массиве nums, которые не равны val, равно k. Чтобы получить результат, необходимо выполнить следующие действия:

Изменить массив nums так, чтобы первые k элементов nums содержали элементы, не равные val. Остальные элементы массива nums, как и размер массива nums, не важны.
*/

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println("Example 1: \nInput: nums =", nums, ", val =", val)
	fmt.Println("Output: k =", removeElement(nums, val), ", val =", nums)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	fmt.Println("Example 2: \nInput: nums =", nums, ", val =", val)
	fmt.Println("Output: k =", removeElement(nums, val), ", val =", nums)
}
