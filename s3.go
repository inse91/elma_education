package main

import (
	"sort"
	"strconv"
	"strings"
)

// input := "[[[1,6,3,4,5,2,7]],[[3,4,3,5,5,7,8,7,8]],[[1,3,4,5,2]]]"

func Ans_for_solution3(body []byte) string {
	input := string(body)
	slices := strings.Split(input[3:len(input)-3], "]],[[") //обрезаем внешние скобки
	slices_allstr := make([]string, len(slices))
	for i := range slices { // формируем слайс строк из лишних эл-тов
		slice_str := strings.Split(slices[i], ",")
		slices_allstr[i] = Sequence_auth(slice_str)
	}
	return strings.Join(slices_allstr, ",")
}

func Sequence_auth(slice_str []string) string {
	slice_int := make([]int, len(slice_str))
	for i := range slice_str {
		slice_int[i], _ = strconv.Atoi(slice_str[i])
	}
	return strconv.Itoa(Solution3(slice_int))
}

func Solution3(A []int) (x int) { // третье задание: ПРОВЕРКА НА ПОСЛЕДОВАТЕЛЬНОСТЬ
	sort.Ints(A)
	x = 1
	for i := range A {
		if i == len(A)-1 {
			break
		}
		if A[i+1] != A[i]+1 {
			x = 0
		}
	}

	return
}
