package main

/* Задача Цикл в связанном списке
Для заданного head, главы связанного списка, определить, есть ли в связанном списке цикл.
В связанном списке есть цикл, если в списке есть узел, к которому можно вернуться, непрерывно следуя по указателю next. Внутри списка pos используется для обозначения индекса узла, к которому привязан указатель next в tail. Обратите внимание, что pos не передаётся как параметр.
Возвращает true, если в связанном списке есть цикл. В противном случае возвращает false.
Ограничения:
Количество узлов в списке находится в диапазоне [0, 10^4].
-10^5 <= Node.val <= 10^5
pos равно -1 или допустимому индексу в связном списке.
Дополнение:
Можете ли вы решить эту задачу, используя O(1) (т.е. константную) память?
это мы сделали с помощью алгоритма Флойда (Floyd’s Cycle Detection Algorithm).
Мы не используем хэш-таблицы или дополнительные массивы (которые дали бы решение за O(n) памяти).
Мы используем только два указателя, значит дополнительная память = O(1).  */

import (
	"fmt"
)

// Определение структуры узла односвязного списка
type ListNode struct {
	Val  int
	Next *ListNode
}

// ─────────────╮
// Функция hasCycle проверяет наличие цикла в связном списке.
// Возвращает true, если цикл есть, иначе false.
func hasCycle(head *ListNode) bool {
	// Если список пустой или состоит только из одного элемента, то цикла быть не может.
	if head == nil || head.Next == nil {
		return false
	}
	// Используем два указателя: "медленный" и "быстрый".
	// Медленный двигается на 1 шаг, быстрый на 2 шага.
	slow := head
	fast := head
	// Пока быстрый и его следующий не nil
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // шаг на 1
		fast = fast.Next.Next // шаг на 2
		// Если указатели встретились, значит есть цикл
		if slow == fast {
			return true
		}
	}
	// Если вышли из цикла — значит fast дошёл до конца списка, цикла нет
	return false
}

// Вспомогательные функции для тестирования
// Функция createLinkedList создает список из массива значений
// и формирует цикл, если pos != -1 (индекс узла, куда указывает хвост).
func createLinkedList(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	// Создаём первый элемент (голову списка)
	head := &ListNode{Val: values[0]}
	current := head
	var cycleNode *ListNode // указатель на узел, куда нужно замкнуть цикл
	// Проходим по массиву и создаём узлы
	for i := 1; i < len(values); i++ {
		newNode := &ListNode{Val: values[i]}
		current.Next = newNode
		current = newNode
		// Если текущий индекс равен pos, запоминаем этот узел
		if i == pos {
			cycleNode = newNode
		}
	}
	// Если pos == 0, цикл указывает на голову
	if pos == 0 {
		cycleNode = head
	}
	// Если pos != -1, создаём цикл
	if pos != -1 {
		current.Next = cycleNode
	}
	return head
}

// ─────────────╯

// ─────────────╮
func main() {

	// Примеры
	tests := []struct {
		values []int
		pos    int
	}{
		{values: []int{3, 2, 0, -4}, pos: 1}, // true
		{values: []int{1, 2}, pos: 0},        // true
		{values: []int{1}, pos: -1},          // false
	}

	for i, test := range tests {
		fmt.Printf("\nExample %d:\n", i+1)
		fmt.Println("Input: head = ", test.values, ", pos = ", test.pos)
		head := createLinkedList(test.values, test.pos)
		fmt.Println(hasCycle(head))
	}
	fmt.Println()
}
