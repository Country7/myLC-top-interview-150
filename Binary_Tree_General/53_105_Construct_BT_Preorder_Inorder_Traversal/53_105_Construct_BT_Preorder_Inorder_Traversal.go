package main

/* Задача Построение двоичного дерева с помощью прямого и симметричного обхода
Даны два целочисленных массива preorder и inorder, где preorder — прямой обход двоичного дерева, а inorder — симметричный обход того же дерева. Постройте и верните двоичное дерево.
Пример 1:
Входные данные: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Выходные данные: [3,9,20,null,null,15,7] */

import (
	"fmt"
)

// Определение структуры узла бинарного дерева
type TreeNode struct {
	Val   int       // Значение узла
	Left  *TreeNode // Указатель на левое поддерево
	Right *TreeNode // Указатель на правое поддерево
}

// ─────────────╮
/* Алгоритм:
1. Preorder (корень → левое → правое)
Первый элемент массива preorder всегда является корнем текущего поддерева.
2. Inorder (левое → корень → правое)
По положению корня в inorder можно понять:
Какие элементы слева от него принадлежат левому поддереву.
Какие справа — правому поддереву.
3. Рекурсия:
Берем текущий корень из preorder.
Находим его индекс в inorder.
Определяем размер левого поддерева.
Рекурсивно строим левое и правое поддеревья.
4. Оптимизация через map:
Чтобы каждый раз не искать индекс корня в inorder (что дало бы O(n²)), используем map для хранения всех индексов — тогда поиск занимает O(1). Общая сложность — O(n).  */
// Функция для построения дерева по массивам preorder и inorder
func buildTree(preorder []int, inorder []int) *TreeNode {
	// Создадим хэш-таблицу (map), которая будет хранить индексы значений из inorder.
	// Это позволит быстро находить позицию корня в inorder-массиве.
	inorderIndexMap := make(map[int]int)
	for i, val := range inorder {
		inorderIndexMap[val] = i
	}
	// Вспомогательная рекурсивная функция
	// preLeft, preRight -> диапазон в preorder
	// inLeft, inRight   -> диапазон в inorder
	var helper func(preLeft, preRight, inLeft, inRight int) *TreeNode
	helper = func(preLeft, preRight, inLeft, inRight int) *TreeNode {
		// Базовый случай: если диапазоны пересеклись — возвращаем nil
		if preLeft > preRight || inLeft > inRight {
			return nil
		}
		// Первый элемент в preorder — это всегда корень поддерева
		rootVal := preorder[preLeft]
		root := &TreeNode{Val: rootVal}
		// Найдем индекс корня в inorder с помощью заранее построенной map
		inRootIndex := inorderIndexMap[rootVal]
		// Количество узлов в левом поддереве
		leftSubtreeSize := inRootIndex - inLeft
		// Рекурсивно строим левое поддерево
		// В preorder элементы левого поддерева идут сразу после корня
		root.Left = helper(
			preLeft+1, preLeft+leftSubtreeSize,
			inLeft, inRootIndex-1,
		)
		// Рекурсивно строим правое поддерево
		root.Right = helper(
			preLeft+leftSubtreeSize+1, preRight,
			inRootIndex+1, inRight,
		)
		return root
	}
	// Запускаем рекурсию на всем диапазоне массивов
	return helper(0, len(preorder)-1, 0, len(inorder)-1)
}

// Вспомогательная функция для вывода дерева (обход в ширину)
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	queue := []*TreeNode{root}
	res := []interface{}{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			res = append(res, node.Val)
			queue = append(queue, node.Left, node.Right)
		} else {
			res = append(res, nil)
		}
	}
	fmt.Println(res)
}

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Пример 1
	preorder1 := []int{3, 9, 20, 15, 7}
	inorder1 := []int{9, 3, 15, 20, 7}
	tree1 := buildTree(preorder1, inorder1)
	printTree(tree1) // Ожидаемый вывод: [3 9 20 <nil> <nil> 15 7]
	// Пример 2
	preorder2 := []int{-1}
	inorder2 := []int{-1}
	tree2 := buildTree(preorder2, inorder2)
	printTree(tree2) // Ожидаемый вывод: [-1]
	fmt.Println()
}
