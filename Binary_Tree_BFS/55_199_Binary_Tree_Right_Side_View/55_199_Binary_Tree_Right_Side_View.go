package main

/* Задача Вид двоичного дерева справа
Дан корень двоичного дерева. Представьте, что вы стоите справа от него. Верните значения узлов, которые вы видите, упорядоченные сверху вниз.
Пример 2:
Входные данные: root = [1,2,3,4,null,null,null,5]
Выходные данные: [1,3,4,5]
Пояснение в картинке */

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
// Функция для получения "правого вида" бинарного дерева
func rightSideView(root *TreeNode) []int {
	// Если дерево пустое, возвращаем пустой срез
	if root == nil {
		return []int{}
	}
	// Результат, где будут храниться значения правых узлов
	result := []int{}
	// Очередь для обхода в ширину (BFS)
	queue := []*TreeNode{root}
	// Пока очередь не пуста, обрабатываем уровни дерева
	for len(queue) > 0 {
		levelSize := len(queue) // количество элементов на текущем уровне
		// Перебираем все узлы текущего уровня
		for i := 0; i < levelSize; i++ {
			node := queue[0]  // берем первый элемент из очереди
			queue = queue[1:] // удаляем его из очереди
			// Если это последний элемент уровня, добавляем его в результат
			if i == levelSize-1 {
				result = append(result, node.Val)
			}
			// Добавляем детей в очередь для следующего уровня
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}

// ─────────────╯

// Вспомогательная функция для создания дерева из примера
func exampleTree1() *TreeNode {
	// Построим дерево: [1,2,3,null,5,null,4]
	return &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3,
			Right: &TreeNode{Val: 4}}}
}

func exampleTree2() *TreeNode {
	// Построим дерево: [1,2,3,4,null,null,null,5]
	return &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 4,
				Left: &TreeNode{Val: 5}}},
		Right: &TreeNode{Val: 3}}
}

func exampleTree3() *TreeNode {
	// Построим дерево: [1,null,3]
	return &TreeNode{Val: 1,
		Right: &TreeNode{Val: 3}}
}

func exampleTree4() *TreeNode {
	// Пустое дерево []
	return nil
}

// ─────────────╮╰──>
func main() {
	fmt.Println()
	fmt.Println("Дерево: [1,2,3,null,5,null,4].  Результат:", rightSideView(exampleTree1()))      // Ожидаем: [1, 3, 4]
	fmt.Println("Дерево: [1,2,3,4,null,null,null,5].  Результат:", rightSideView(exampleTree2())) // Ожидаем: [1, 3, 4, 5]
	fmt.Println("Дерево: [1,null,3].  Результат:", rightSideView(exampleTree3()))                 // Ожидаем: [1, 3]
	fmt.Println("Дерево: [].  Результат:", rightSideView(exampleTree4()))                         // Ожидаем: []
	fmt.Println()
}
