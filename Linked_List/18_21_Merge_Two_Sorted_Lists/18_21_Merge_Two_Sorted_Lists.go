package main

/* Задача Объединение двух отсортированных списков
Вам даны заголовки двух отсортированных связных списков: list1 и list2.
Объедините два списка в один отсортированный список. Список должен быть получен путем объединения узлов первых двух списков.
Вернуть заголовок объединённого связного списка.
Ограничения:
Количество узлов в обоих списках находится в диапазоне [0, 50].
-100 <= Node.val <= 100
Оба списка: list1 и list2 отсортированы по неубывающему значению. */

import (
	"fmt"
)

// Определение структуры узла односвязного списка
type ListNode struct {
	Val  int       // значение узла
	Next *ListNode // ссылка на следующий узел
}

// ─────────────╮
// Функция mergeTwoLists объединяет два отсортированных списка в один отсортированный
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Создаём "фиктивный" (вспомогательный) узел, чтобы упростить работу
	// С его помощью можно избежать множества проверок на "nil"
	dummy := &ListNode{}
	// Указатель, который будет двигаться по результирующему списку
	current := dummy
	// Пока оба списка не пустые
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			// Если значение в list1 меньше или равно значению в list2,
			// то "подвязываем" этот узел к результату
			current.Next = list1
			// Сдвигаем указатель list1 вперёд
			list1 = list1.Next
		} else {
			// Иначе берём узел из list2
			current.Next = list2
			list2 = list2.Next
		}
		// Сдвигаем указатель результата вперёд
		current = current.Next
	}
	// Если один из списков закончился, подвязываем оставшуюся часть второго списка
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}
	// Возвращаем "реальный" результат (начиная с узла после фиктивного)
	return dummy.Next
}

// Вспомогательная функция для создания списка из среза
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

// Вспомогательная функция для вывода списка в консоль
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Пример 1
	list1 := createList([]int{1, 2, 4})
	list2 := createList([]int{1, 3, 4})
	result := mergeTwoLists(list1, list2)
	fmt.Print("Пример 1: ")
	printList(result) // Ожидаем: 1 -> 1 -> 2 -> 3 -> 4 -> 4

	// Пример 2
	list1 = createList([]int{})
	list2 = createList([]int{})
	result = mergeTwoLists(list1, list2)
	fmt.Print("Пример 2: ")
	printList(result) // Ожидаем: пустой вывод

	// Пример 3
	list1 = createList([]int{})
	list2 = createList([]int{0})
	result = mergeTwoLists(list1, list2)
	fmt.Print("Пример 3: ")
	printList(result) // Ожидаем: 0
	fmt.Println()

}
