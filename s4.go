package main

import (
	"sort"
	"strconv"
	"strings"
)

// input := "[[[1,6,3,4,5,2,8]],[[3,4,6,7,5,8,10]],[[1,3,4,5,2,7]]]"

func Ans_for_solution4(body []byte) string {
	input := string(body)
	slices := strings.Split(input[3:len(input)-3], "]],[[") //обрезаем внешние скобки
	slices_allstr := make([]string, len(slices))
	for i := range slices { // формируем слайс строк из лишних эл-тов
		slice_str := strings.Split(slices[i], ",")
		slices_allstr[i] = Sequence_find(slice_str)

	}
	return strings.Join(slices_allstr, ",")
}

func Sequence_find(slice_str []string) string {
	slice_int := make([]int, len(slice_str))
	for i := range slice_str {
		slice_int[i], _ = strconv.Atoi(slice_str[i])
	}
	return strconv.Itoa(Solution4(slice_int))
}

func Solution4(A []int) (x int) { // четвертое задание: ПОИСК НЕДОСТАЮЩЕГО ЭЛЕМЕНТА В ПОСЛЕДОВАТЕЛЬНОСТИ
	sort.Ints(A)

	for i := range A {
		if i == len(A)-1 {
			break
		}
		if A[i+1] != A[i]+1 {
			x = A[i] + 1
		}
	}

	return
}
