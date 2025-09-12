package main

/* Задача Сумма пути
Дан корень двоичного дерева и целочисленная сумма targetSum. Верните значение true, если дерево имеет путь от корня к листьям, такой, что сумма всех значений на пути равна targetSum.
Лист — это узел без дочерних узлов. */

import (
	"fmt"
)

// Определение структуры для узла бинарного дерева
type TreeNode struct {
	Val   int       // значение узла
	Left  *TreeNode // левый потомок
	Right *TreeNode // правый потомок
}

// ─────────────╮
// Функция проверяет, существует ли путь от корня до листа, сумма значений на котором равна targetSum.
func hasPathSum(root *TreeNode, targetSum int) bool {
	// Базовый случай: если дерево пустое, то пути не существует
	if root == nil {
		return false
	}
	// Если мы находимся в листе (оба потомка равны nil),
	// то проверяем, равна ли сумма на этом пути targetSum
	if root.Left == nil && root.Right == nil {
		// Если значение узла совпадает с targetSum, значит путь найден
		return root.Val == targetSum
	}
	// Рекурсивно проверяем левое и правое поддерево. При этом уменьшаем targetSum на значение текущего узла.
	// Если хотя бы в одном из поддеревьев найден путь, возвращаем true.
	return hasPathSum(root.Left, targetSum-root.Val) ||
		hasPathSum(root.Right, targetSum-root.Val)
}

// Вспомогательная функция для наглядного примера: строим дерево вручную
func buildExampleTree() *TreeNode {
	/*
	   Пример дерева:
	           5
	          / \
	         4   8
	        /   / \
	       11  13  4
	      /  \       \
	     7    2       1
	*/
	return &TreeNode{Val: 5,
		Left: &TreeNode{Val: 4,
			Left: &TreeNode{Val: 11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2}}},
		Right: &TreeNode{Val: 8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{Val: 4,
				Right: &TreeNode{Val: 1}}}}
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Строим дерево пример
	root := buildExampleTree()
	// Проверим пример: targetSum = 22
	targetSum := 22
	fmt.Println("Есть ли путь с суммой", targetSum, "?", hasPathSum(root, targetSum)) // Ожидаем: true
	// Проверим другой пример: root = [1,2,3], targetSum = 5
	root2 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3}}
	fmt.Println("Есть ли путь с суммой 5?", hasPathSum(root2, 5)) // Ожидаем: false
	// Проверим случай пустого дерева
	fmt.Println("Есть ли путь в пустом дереве с суммой 0?", hasPathSum(nil, 0)) // Ожидаем: false
	fmt.Println()
}
