package main

/* Задача

 */

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
// Функция использует алгоритм обхода в ширину (BFS) для обработки дерева по уровням
// Временная сложность алгоритма: O(n), где n - количество узлов в дереве
func averageOfLevels(root *TreeNode) []float64 {
	// Если дерево пустое, возвращаем пустой слайс
	if root == nil {
		return []float64{}
	}

	// Используем обход в ширину (BFS) с помощью очереди
	var result []float64
	queue := []*TreeNode{root} // Начинаем с корневого узла

	for len(queue) > 0 {
		levelSize := len(queue) // Количество узлов на текущем уровне
		levelSum := 0.0         // Сумма значений узлов текущего уровня

		// Обрабатываем все узлы текущего уровня
		for i := 0; i < levelSize; i++ {
			// Извлекаем узел из начала очереди
			currentNode := queue[0]
			queue = queue[1:]

			// Добавляем значение текущего узла к сумме уровня
			levelSum += float64(currentNode.Val)

			// Добавляем дочерние узлы в очередь для обработки на следующих уровнях
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}

		// Вычисляем среднее значение для текущего уровня и добавляем в результат
		average := levelSum / float64(levelSize)
		result = append(result, average)
	}

	return result
}

// Вспомогательная функция для создания нового узла дерева
func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// ─────────────╯

// ─────────────╮
func main() {
	fmt.Println()
	// Создаем дерево из примера 1: [3,9,20,null,null,15,7]
	//       3
	//      / \
	//     9  20
	//        / \
	//       15  7
	root1 := newNode(3)
	root1.Left = newNode(9)
	root1.Right = newNode(20)
	root1.Right.Left = newNode(15)
	root1.Right.Right = newNode(7)

	result1 := averageOfLevels(root1)
	fmt.Printf("Пример 1: %v\n", result1) // Ожидаемый результат: [3, 14.5, 11]

	// Создаем дерево из примера 2: [3,9,20,15,7]
	//       3
	//      / \
	//     9  20
	//    / \
	//   15  7
	root2 := newNode(3)
	root2.Left = newNode(9)
	root2.Right = newNode(20)
	root2.Left.Left = newNode(15)
	root2.Left.Right = newNode(7)

	result2 := averageOfLevels(root2)
	fmt.Printf("Пример 2: %v\n", result2) // Ожидаемый результат: [3, 14.5, 11]

	// Тест с одним узлом
	root3 := newNode(5)
	result3 := averageOfLevels(root3)
	fmt.Printf("Один узел: %v\n", result3) // Ожидаемый результат: [5]

}
