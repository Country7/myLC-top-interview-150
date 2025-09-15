package main

/* Задача Построение двоичного дерева с помощью инпорядкового и постпорядкового обходов
Даны два целочисленных массива inorder и postorder, где inorder — инпорядковое обход двоичного дерева, а postorder — инпорядковое обход того же дерева. Постройте и верните двоичное дерево.
Пример 1:
Входные данные: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
Выходные данные: [3,9,20,null,null,15,7] */

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
/* 1. Основная идея:
Последний элемент в postorder — это корень текущего поддерева.
Находим этот элемент в inorder, разделяя массив на левое и правое поддерево.
Рекурсивно строим сначала правое поддерево, потом левое (так как мы двигаемся по postorder с конца).
2. Оптимизация:
Чтобы быстро находить индекс корня в inorder, используем map[value]index.
Это снижает сложность до O(n) вместо O(n²).
3. Вывод дерева:
Используется обход в ширину (BFS), чтобы показать результат в формате задачи (как массив уровнями). */
// Функция для построения дерева по inorder и postorder
func buildTree(inorder []int, postorder []int) *TreeNode {
	// Создаём хэш-таблицу для быстрого поиска индекса элемента в inorder
	// Ключ: значение узла, Значение: индекс в inorder
	inorderIndexMap := make(map[int]int)
	for i, val := range inorder {
		inorderIndexMap[val] = i
	}
	// Определяем индекс последнего элемента в postorder (корень дерева)
	postIndex := len(postorder) - 1
	// Вспомогательная рекурсивная функция
	var helper func(left, right int) *TreeNode
	helper = func(left, right int) *TreeNode {
		// Если нет элементов для построения поддерева
		if left > right {
			return nil
		}
		// Берём текущий корень из postorder
		rootVal := postorder[postIndex]
		postIndex--
		// Создаём узел
		root := &TreeNode{Val: rootVal}
		// Получаем индекс корня в inorder
		index := inorderIndexMap[rootVal]
		// ВАЖНО: сначала строим правое поддерево, потом левое
		// так как мы идём с конца postorder (сначала правый ребёнок)
		root.Right = helper(index+1, right)
		root.Left = helper(left, index-1)
		return root
	}
	// Запускаем рекурсию для всего диапазона
	return helper(0, len(inorder)-1)
}

// Функция для печати дерева в формате "уровневого обхода" (BFS)
// Чтобы проверить результат
func printLevelOrder(root *TreeNode) {
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
			result = append(result, nil)
		}
	}
	// Убираем лишние nil в конце
	i := len(result) - 1
	for i >= 0 && result[i] == nil {
		i--
	}
	result = result[:i+1]
	fmt.Println(result)
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Пример 1
	inorder1 := []int{9, 3, 15, 20, 7}
	postorder1 := []int{9, 15, 7, 20, 3}
	tree1 := buildTree(inorder1, postorder1)
	fmt.Print("Пример 1: ")
	printLevelOrder(tree1)
	// Пример 2
	inorder2 := []int{-1}
	postorder2 := []int{-1}
	tree2 := buildTree(inorder2, postorder2)
	fmt.Print("Пример 2: ")
	printLevelOrder(tree2)
	fmt.Println()
}
