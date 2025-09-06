package main

/* Задача Инвертировать двоичное дерево
Дан корень двоичного дерева, инвертировать дерево и вернуть его корень.
Количество узлов в дереве находится в диапазоне [0, 100].
-100 <= Node.val <= 100
*/

import (
	"fmt"
)

// Определение структуры узла бинарного дерева
type TreeNode struct {
	Val   int       // значение узла
	Left  *TreeNode // указатель на левое поддерево
	Right *TreeNode // указатель на правое поддерево
}

// ─────────────╮
// Функция для инверсии бинарного дерева
func invertTree(root *TreeNode) *TreeNode {
	// Базовый случай рекурсии: если узел пустой (nil), просто возвращаем nil
	if root == nil {
		return nil
	}
	// Меняем местами левое и правое поддерево
	root.Left, root.Right = root.Right, root.Left
	// Рекурсивно инвертируем левое поддерево
	invertTree(root.Left)
	// Рекурсивно инвертируем правое поддерево
	invertTree(root.Right)
	// Возвращаем текущий корень (уже инвертированный)
	return root
}

// Вспомогательная функция для печати дерева (обход в ширину)
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	queue := []*TreeNode{root}
	result := []interface{}{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		} else {
			// Если узел пустой, добавляем nil для наглядности
			result = append(result, nil)
		}
	}
	fmt.Println(result)
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Пример 1: root = [4,2,7,1,3,6,9]
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}
	fmt.Print("Оригинальное дерево: ")
	printTree(root)
	inverted := invertTree(root)
	fmt.Print("Инвертированное дерево: ")
	printTree(inverted)
	// Пример 2: root = [2,1,3]
	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 1}
	root2.Right = &TreeNode{Val: 3}
	fmt.Print("Оригинальное дерево: ")
	printTree(root2)
	inverted2 := invertTree(root2)
	fmt.Print("Инвертированное дерево: ")
	printTree(inverted2)
	// Пример 3: root = []
	var root3 *TreeNode = nil
	fmt.Print("Оригинальное дерево: ")
	printTree(root3)
	inverted3 := invertTree(root3)
	fmt.Print("Инвертированное дерево: ")
	printTree(inverted3)
	fmt.Println()
}
