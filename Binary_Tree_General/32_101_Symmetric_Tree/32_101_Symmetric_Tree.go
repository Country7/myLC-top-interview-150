package main

/* Задача Симметричное дерево
Дан корень двоичного дерева. Проверьте, является ли он зеркальным отражением самого себя (т.е. симметричен ли он относительно своего центра). */

import (
	"fmt"
)

// Определение структуры узла бинарного дерева
type TreeNode struct {
	Val   int       // значение узла
	Left  *TreeNode // ссылка на левое поддерево
	Right *TreeNode // ссылка на правое поддерево
}

// ─────────────╮
// Функция isSymmetric проверяет, является ли дерево симметричным
func isSymmetric(root *TreeNode) bool {
	// Пустое дерево считается симметричным
	if root == nil {
		return true
	}
	// Проверяем симметричность левого и правого поддерева
	return isMirror(root.Left, root.Right)
}

// Вспомогательная функция isMirror рекурсивно проверяет два поддерева на зеркальность
func isMirror(t1, t2 *TreeNode) bool {
	// Если оба поддерева пустые — они зеркальны
	if t1 == nil && t2 == nil {
		return true
	}
	// Если только одно из поддеревьев пустое — они не зеркальны
	if t1 == nil || t2 == nil {
		return false
	}
	// Условие зеркальности:
	// 1) Значения текущих узлов должны быть равны
	// 2) Левое поддерево первого узла зеркально правому поддереву второго
	// 3) Правое поддерево первого узла зеркально левому поддереву второго
	return (t1.Val == t2.Val) &&
		isMirror(t1.Left, t2.Right) &&
		isMirror(t1.Right, t2.Left)
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Пример 1: [1,2,2,3,4,4,3]
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 2}
	root1.Left.Left = &TreeNode{Val: 3}
	root1.Left.Right = &TreeNode{Val: 4}
	root1.Right.Left = &TreeNode{Val: 4}
	root1.Right.Right = &TreeNode{Val: 3}
	fmt.Println("Симметричность дерева: [1,2,2,3,4,4,3] = ", isSymmetric(root1)) // true
	// Пример 2: [1,2,2,null,3,null,3]
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 2}
	root2.Left.Right = &TreeNode{Val: 3}
	root2.Right.Right = &TreeNode{Val: 3}
	fmt.Println("Симметричность дерева: [1,2,2,null,3,null,3] = ", isSymmetric(root2)) // false
	fmt.Println()
}
