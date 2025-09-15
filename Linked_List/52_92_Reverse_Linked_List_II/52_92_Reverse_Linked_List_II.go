package main

/* Задача Обратносвязанный список II
Дан заголовок односвязного списка и два целых числа left и right, где left <= right, перевернуть узлы списка с позиции left на позицию right и вернуть обратный список.
Пример 1:
Входные данные: head = [1,2,3,4,5], left = 2, right = 4
Выходные данные: [1,4,3,2,5] */

import (
	"fmt"
)

// Определение структуры узла односвязного списка
type ListNode struct {
	Val  int
	Next *ListNode
}

// ─────────────╮
// Функция разворота подсписка с позиции left до right
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// Создаём фиктивный узел перед головой списка
	// Это удобно для обработки случаев, когда разворот начинается с первого элемента
	dummy := &ListNode{Next: head}
	// prev будет указывать на узел перед позицией "left"
	prev := dummy
	for i := 1; i < left; i++ {
		prev = prev.Next
	}
	// start указывает на первый узел подсписка (позиция left)
	start := prev.Next
	// then будет указывать на следующий узел после start
	then := start.Next
	// Разворот выполняем по принципу: "вытаскиваем then и вставляем его после prev"
	for i := 0; i < right-left; i++ {
		// Переставляем ссылки, чтобы "then" переместился в начало подсписка
		start.Next = then.Next
		then.Next = prev.Next
		prev.Next = then
		// Сдвигаем then на следующий узел, который нужно обработать
		then = start.Next
	}
	// Возвращаем новую голову списка (dummy.Next)
	return dummy.Next
}

// Вспомогательная функция для создания списка из массива
func createList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, v := range nums {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return dummy.Next
}

// Вспомогательная функция для печати списка
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
	// Пример 1: head = [1,2,3,4,5], left = 2, right = 4
	head := createList([]int{1, 2, 3, 4, 5})
	fmt.Print("Исходный список: ")
	printList(head)
	newHead := reverseBetween(head, 2, 4)
	fmt.Print("После разворота (2-4): ")
	printList(newHead)
	// Пример 2: head = [5], left = 1, right = 1
	head2 := createList([]int{5})
	fmt.Print("Исходный список: ")
	printList(head2)
	newHead2 := reverseBetween(head2, 1, 1)
	fmt.Print("После разворота (1-1): ")
	printList(newHead2)
	fmt.Println()
}
