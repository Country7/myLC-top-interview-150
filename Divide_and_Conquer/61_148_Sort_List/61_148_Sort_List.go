package main

/* Задача Сортировка списка
Дан заголовок связного списка, вернуть список после его сортировки по возрастанию. */

import (
	"fmt"
)

// Определение структуры узла односвязного списка
type ListNode struct {
	Val  int
	Next *ListNode
}

// ─────────────╮
// оптимально использовать сортировку слиянием (Merge Sort), так как она работает за O(n log n) и не требует дополнительной памяти (кроме рекурсии)
// Основная функция сортировки списка
func sortList(head *ListNode) *ListNode {
	fmt.Println("   ╰──> head", listToArray(head))
	// Базовый случай: если список пустой или в нём только один элемент - возвращаем как есть
	if head == nil || head.Next == nil {
		return head
	}
	// Находим середину списка с помощью двух указателей (быстрый и медленный)
	mid := getMiddle(head)
	// Разделяем список на две части
	right := mid.Next
	mid.Next = nil
	fmt.Println("   ╰──> left", listToArray(head), "right", listToArray(right))
	// Рекурсивно сортируем левую и правую половины
	leftSorted := sortList(head)
	rightSorted := sortList(right)
	fmt.Println("   ╰──> leftSorted", listToArray(leftSorted), ", rightSorted", listToArray(rightSorted))
	// Сливаем отсортированные половины в один отсортированный список
	return merge(leftSorted, rightSorted)
}

// Функция для поиска середины списка (используется алгоритм "быстрый и медленный указатель")
func getMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	// Двигаем fast на 2 узла, slow на 1 узел, пока fast и fast.Next не достигнут конца
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// Функция слияния двух отсортированных списков
func merge(l1 *ListNode, l2 *ListNode) *ListNode {
	// Создаём "фейковый" узел для удобства
	dummy := &ListNode{}
	current := dummy
	// Пока оба списка не закончились, выбираем минимальный элемент и добавляем в результат
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	// Если один список закончился, добавляем оставшуюся часть другого списка
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}
	return dummy.Next
}

// Вспомогательная функция для создания списка из массива
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, v := range nums[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

// Вспомогательная функция для преобразования списка в массив (для вывода)
func listToArray(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Пример 1
	list1 := createList([]int{4, 2, 1, 3})
	fmt.Println("Исходный список: {4, 2, 1, 3}")
	sorted1 := sortList(list1)
	fmt.Println("Отсортированный:", listToArray(sorted1), "\n ") // [1 2 3 4]
	// Пример 2
	list2 := createList([]int{-1, 5, 3, 4, 0})
	fmt.Println("Исходный список: {-1, 5, 3, 4, 0}")
	sorted2 := sortList(list2)
	fmt.Println("Отсортированный:", listToArray(sorted2), "\n ") // [-1 0 3 4 5]
	// Пример 3 (пустой список)
	list3 := createList([]int{})
	fmt.Println("Исходный список: {}")
	sorted3 := sortList(list3)
	fmt.Println("Отсортированный:", listToArray(sorted3), "\n ") // []
	fmt.Println()
}
