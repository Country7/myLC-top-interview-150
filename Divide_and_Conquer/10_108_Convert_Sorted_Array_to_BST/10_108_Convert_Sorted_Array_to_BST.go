package main

/* Задача Преобразование отсортированного массива в двоичное дерево поиска
Дан целочисленный массив nums, элементы которого отсортированы по возрастанию, и преобразование его в сбалансированное по высоте двоичное дерево поиска. */

/* Для этого эффективным решением будет использование подхода, который рекурсивно выбирает средний элемент массива в качестве корня дерева, а затем рекурсивно строит поддеревья из левой и правой части массива. */

import (
	"fmt"
)

// ─────────────╮
// Определяем структуру для узла бинарного дерева поиска.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Функция для преобразования отсортированного массива в сбалансированное бинарное дерево поиска.
func sortedArrayToBST(nums []int) *TreeNode {
	// Вызов рекурсивной функции для создания дерева.
	return convertToBST(nums, 0, len(nums)-1)
}

// Рекурсивная функция для создания дерева.
func convertToBST(nums []int, left, right int) *TreeNode {
	// Базовый случай: если левая граница больше правой, возвращаем nil (пустое поддерево).
	if left > right {
		return nil
	}
	// Находим индекс среднего элемента для сбалансированного разделения массива.
	mid := left + (right-left)/2
	// Создаем новый узел дерева с данным значением.
	root := &TreeNode{
		Val: nums[mid],
	}
	// Рекурсивно строим левое и правое поддерево.
	root.Left = convertToBST(nums, left, mid-1)
	root.Right = convertToBST(nums, mid+1, right)
	// Возвращаем корень текущего поддерева.
	return root
}

// Функция для печати бинарного дерева в виде массива (для удобства проверки результата).
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("nil ")
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

// ─────────────╯

// ─────────────╮
func main() {
	fmt.Println()
	// Пример 1
	nums1 := []int{-10, -3, 0, 5, 9}
	tree1 := sortedArrayToBST(nums1)
	fmt.Print("Tree 1: ")
	printTree(tree1) // Ожидаемый вывод: [0, -3, 9, -10, nil, 5]
	fmt.Println()
	// Пример 2
	nums2 := []int{1, 3}
	tree2 := sortedArrayToBST(nums2)
	fmt.Print("Tree 2: ")
	printTree(tree2) // Ожидаемый вывод: [3, 1, nil]
	fmt.Println()
}
