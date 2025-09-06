package main

/* Задача Одно и то же дерево
Для корней двух двоичных деревьев p и q напишите функцию для проверки их идентичности.
Два двоичных дерева считаются одинаковыми, если их структура идентична, а узлы имеют одинаковые значения.
Ограничения:
Количество узлов в обоих деревьях находится в диапазоне [0, 100].
-10^4 <= Node.val <= 10^4   */

import (
	"fmt"
)

// Определение узла бинарного дерева.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ─────────────╮
/*  isSameTree — рекурсивная проверка "одинаковости" двух бинарных деревьев.
1) Если оба текущих узла равны nil — поддеревья тождественны (пустые).
2) Если один nil, а другой нет — структуры различаются.
3) Если оба не nil, но значения различаются — деревья различаются.
4) Рекурсивно сравниваем левые поддеревья и правые поддеревья.
Временная сложность: O(n), где n — минимальное количество узлов в p и q (в худшем — суммарно по всем узлам). */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Случай 1: оба пустые — одинаковые
	if p == nil && q == nil {
		return true
	}
	// Случай 2: один пустой, другой нет — точно разные
	if p == nil || q == nil {
		return false
	}
	// Случай 3: оба существуют, но значения различаются — разные
	if p.Val != q.Val {
		return false
	}
	// Случай 4: значения равны — сравниваем левые и правые поддеревья
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

/*
	Альтернативная итеративная версия на очереди (BFS). Полезно, если вы хотите избежать рекурсии.

- Держим очередь пар узлов (из p и q), начиная с корней.
- На каждом шаге достаём пару:
  - Если обе nil — норм, продолжаем.
  - Если одна nil — деревья различаются.
  - Если оба не nil, но значения не равны — различаются.
  - Иначе кладём в очередь их детей-пары: (Left, Left) и (Right, Right).

Сложность: такая же по порядку, что и у рекурсивного решения.
*/
func isSameTreeIter(p *TreeNode, q *TreeNode) bool {
	type pair struct {
		a, b *TreeNode
	}
	queue := []pair{{p, q}}
	for len(queue) > 0 {
		// Забираем первый элемент
		cur := queue[0]
		queue = queue[1:]
		a, b := cur.a, cur.b
		if a == nil && b == nil {
			// Оба пустые — эта "позиция" совпадает
			continue
		}
		if a == nil || b == nil {
			// Ровно один пустой — различаются
			return false
		}
		if a.Val != b.Val {
			// Значения различаются — деревья разные
			return false
		}
		// Добавляем соответствующие пары детей
		queue = append(queue, pair{a.Left, b.Left})
		queue = append(queue, pair{a.Right, b.Right})
	}
	return true
}

// Функция строит дерево из level-order (уровневого) представления, где nil означает отсутствие узла
// - Создаём корень из первого элемента.
// - Идём по слайсу, по очереди назначая для каждого узла левого и правого ребёнка.
func buildTreeLevelOrder(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	// Создаём корень
	root := &TreeNode{Val: *vals[0]}
	queue := []*TreeNode{root}
	i := 1 // индекс следующего значения в слайсе
	for i < len(vals) && len(queue) > 0 {
		// Берём текущий узел из очереди
		cur := queue[0]
		queue = queue[1:]
		// Левый ребёнок
		if i < len(vals) {
			if vals[i] != nil {
				cur.Left = &TreeNode{Val: *vals[i]}
				queue = append(queue, cur.Left)
			}
			i++
		}
		// Правый ребёнок
		if i < len(vals) {
			if vals[i] != nil {
				cur.Right = &TreeNode{Val: *vals[i]}
				queue = append(queue, cur.Right)
			}
			i++
		}
	}
	return root
}

// Удобный хелпер для получения *int из литерала (чтобы легко формировать []*int).
func ip(x int) *int { return &x }

// ─────────────╯

// ─────────────╮╰──>
func main() {
	fmt.Println()
	// Пример 1:
	// p = [1,2,3], q = [1,2,3] -> true
	p1 := buildTreeLevelOrder([]*int{ip(1), ip(2), ip(3)})
	q1 := buildTreeLevelOrder([]*int{ip(1), ip(2), ip(3)})
	fmt.Println("Пример 1 (рекурсия):", isSameTree(p1, q1))     // true
	fmt.Println("Пример 1 (итерация):", isSameTreeIter(p1, q1)) // true
	// Пример 2:
	// p = [1,2], q = [1,null,2] -> false
	p2 := buildTreeLevelOrder([]*int{ip(1), ip(2)})
	q2 := buildTreeLevelOrder([]*int{ip(1), nil, ip(2)})
	fmt.Println("Пример 2 (рекурсия):", isSameTree(p2, q2))     // false
	fmt.Println("Пример 2 (итерация):", isSameTreeIter(p2, q2)) // false
	// Пример 3:
	// p = [1,2,1], q = [1,1,2] -> false
	p3 := buildTreeLevelOrder([]*int{ip(1), ip(2), ip(1)})
	q3 := buildTreeLevelOrder([]*int{ip(1), ip(1), ip(2)})
	fmt.Println("Пример 3 (рекурсия):", isSameTree(p3, q3))     // false
	fmt.Println("Пример 3 (итерация):", isSameTreeIter(p3, q3)) // false
}
