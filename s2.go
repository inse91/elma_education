package main

import (
	"strconv"
	"strings"
)

// input := "[[[1,6,1,3,5,6,5]],[[3,4,3,5,5,7,8,7,8]],[[1,3,1,7,3]]]"

func Ans_for_solution2(body []byte) string {
	input := string(body)
	slices := strings.Split(input[3:len(input)-3], "]],[[") //обрезаем внешние скобки
	slices_allstr := make([]string, len(slices))
	for i := range slices { // формируем слайс строк из лишних эл-тов
		slice_str := strings.Split(slices[i], ",")
		slices_allstr[i] = ElemWOApair(slice_str)
	}
	return strings.Join(slices_allstr, ",")
}

func ElemWOApair(slice_str []string) string {
	slice_int := make([]int, len(slice_str))
	for i := range slice_str {
		slice_int[i], _ = strconv.Atoi(slice_str[i])
	}
	return strconv.Itoa(Solution2(slice_int))
}

func Solution2(A []int) (x int) { // второе задание: ПОИСК ЭЛЕМЕНТА БЕЗ ПАРЫ

	numsbool := map[int]bool{}

	for _, val := range A {
		if _, ok := numsbool[val]; ok {
			numsbool[val] = !numsbool[val]
		} else {
			numsbool[val] = false
		}
	}

	for idx, val := range numsbool {
		if val == false {
			x = idx
		}
	}
	return
}
