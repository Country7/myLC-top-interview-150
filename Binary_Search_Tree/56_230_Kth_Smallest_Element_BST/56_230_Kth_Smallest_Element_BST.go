package main

/* Задача K-й наименьший элемент в BST-дереве
Дан корень двоичного дерева поиска и целое число k. Вернуть k-е наименьшее значение (индексированное с 1) среди всех значений узлов дерева.  */

import (
	"fmt"
)

// Определение структуры узла бинарного дерева поиска (BST)
type TreeNode struct {
	Val   int       // значение в узле
	Left  *TreeNode // ссылка на левое поддерево
	Right *TreeNode // ссылка на правое поддерево
}

// ─────────────╮
// Функция для нахождения k-го наименьшего элемента в BST
func kthSmallest(root *TreeNode, k int) int {
	// Переменные для хранения результата и счётчика посещённых узлов
	count := 0
	result := -1
	// Вспомогательная рекурсивная функция обхода "in-order" (лево → корень → право)
	// В BST такой порядок обхода даёт отсортированный по возрастанию список значений
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		// Сначала обходим левое поддерево
		inorder(node.Left)
		// Увеличиваем счётчик при посещении узла
		count++
		// Если текущий узел оказался k-м по счёту — сохраняем результат
		if count == k {
			result = node.Val
			return
		}
		// Обходим правое поддерево
		inorder(node.Right)
	}
	// Запускаем рекурсивный обход с корня дерева
	inorder(root)
	return result
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 4}
	k := 1
	fmt.Printf("k-й наименьший элемент (k=%d): %d\n", k, kthSmallest(root, k))
	fmt.Println()
}
