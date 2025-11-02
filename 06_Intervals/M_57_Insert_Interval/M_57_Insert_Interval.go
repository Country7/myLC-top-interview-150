package main

import "fmt"

/* Задача 57. Вставка интервала  ✅
Вам дан массив неперекрывающихся интервалов intervals, где intervals[i] = [starti, endi] представляют начало и конец i-го интервала, а intervals отсортирован по возрастанию starti. Вам также дан интервал newInterval = [start, end], представляющий начало и конец другого интервала.
Вставьте newInterval в intervals так, чтобы intervals по-прежнему был отсортирован по возрастанию starti, а в intervals по-прежнему не было перекрывающихся интервалов (при необходимости объедините перекрывающиеся интервалы).
Верните intervals после вставки.
Обратите внимание, что вам не нужно изменять intervals на месте. Вы можете создать новый массив и вернуть его.

Пример 1:
Входные данные: intervals = [[1,3],[6,9]], newInterval = [2,5]
Выходные данные: [[1,5],[6,9]]

Пример 2:
Входные данные: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Выходные данные: [[1,2],[3,10],[12,16]]
Пояснение: Поскольку новый интервал [4,8] пересекается с [3,5],[6,7],[8,10].

Ограничения:
intervals[i].length == 2
intervals отсортирован по starti в порядке возрастания.
newInterval.length == 2  */

func insert(intervals [][]int, newInterval []int) [][]int {
	i := 0
	result := [][]int{}
	lenInt := len(intervals)
	newStart := newInterval[0]
	newEnd := newInterval[1]

	// intervals = [[1,2] | ,[3,5],[6,7],[8,10], | [12,16]],
	// newInterval = [4,8]

	// все интервалы до newInterval
	for i < lenInt && intervals[i][1] < newStart {
		result = append(result, intervals[i])
		i++
	}

	// перекрывающиеся интервалы
	for i < lenInt && intervals[i][0] <= newEnd {
		if intervals[i][0] < newStart {
			newStart = intervals[i][0] // 3
		}
		if intervals[i][1] > newEnd {
			newEnd = intervals[i][1] // 10
		}
		i++
	}
	result = append(result, []int{newStart, newEnd})

	// оставшиеся интервалы
	for i < lenInt {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	// [[1 5] [6 9]]

	fmt.Println(insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	// [[1 2] [3 10] [12 16]]
}
