package main

/* Задача Максимальная глубина двоичного дерева
Для корня двоичного дерева вернуть его максимальную глубину.
Максимальная глубина двоичного дерева — это количество узлов на самом длинном пути от корня до самого дальнего листового узла.
Временная сложность алгоритма: O(n)
Ограничения:
Количество узлов в дереве находится в диапазоне [0, 10^4].
-100 <= Node.val <= 100.  */

import (
	"fmt"
)

// Определение структуры узла бинарного дерева
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ─────────────╮
// Функция для вычисления максимальной глубины бинарного дерева
func maxDepth(root *TreeNode) int {
	// Базовый случай: если узел nil (пустое дерево или достигнут конец ветки)
	if root == nil {
		return 0
	}
	// Рекурсивно вычисляем глубину левого поддерева
	leftDepth := maxDepth(root.Left)
	// Рекурсивно вычисляем глубину правого поддерева
	rightDepth := maxDepth(root.Right)
	// Возвращаем максимальную глубину из левого и правого поддеревьев + 1 (текущий узел)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// Вспомогательная функция для создания нового узла дерева
func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// ─────────────╯

// ─────────────╮
func main() {
	fmt.Println("\nЗадача Максимальная глубина двоичного дерева")

	// Тест 1: Пример из условия [3,9,20,null,null,15,7]
	fmt.Println("\nТест 1:")
	root1 := newNode(3)
	root1.Left = newNode(9)
	root1.Right = newNode(20)
	root1.Right.Left = newNode(15)
	root1.Right.Right = newNode(7)

	result1 := maxDepth(root1)
	fmt.Printf("Ожидаемый результат: 3, Полученный результат: %d\n", result1)

	// Тест 2: Пример из условия [1,null,2]
	fmt.Println("\nТест 2:")
	root2 := newNode(1)
	root2.Right = newNode(2)

	result2 := maxDepth(root2)
	fmt.Printf("Ожидаемый результат: 2, Полученный результат: %d\n", result2)

	// Тест 3: Пустое дерево
	fmt.Println("\nТест 3:")
	var root3 *TreeNode = nil

	result3 := maxDepth(root3)
	fmt.Printf("Ожидаемый результат: 0, Полученный результат: %d\n", result3)

	// Тест 4: Дерево с одним узлом
	fmt.Println("\nТест 4:")
	root4 := newNode(5)

	result4 := maxDepth(root4)
	fmt.Printf("Ожидаемый результат: 1, Полученный результат: %d\n", result4)
}
