package main

import (
	"strconv"
	"strings"
)

//input := "[[[1,6,3,4,5,2,8],3],[[3,4,6,7,5,8,10],4],[[1,3,4,5,2,7],5]]"

func Ans_for_solution1(body []byte) string {
	a := string(body[2 : len(body)-2]) //обрезаем внешние скобки
	b := strings.Split(a, "],[")       //делим на пакеты данных для одного решения

	datastructs := make([]Sol1, len(b))

	for i := range b {

		c := []byte(b[i])[1:] //обрезаем две первые скобки

		d := strings.Split(string(c), "],") // d[1] = это переменная для сдвига

		e := strings.Split(d[0], ",")
		f := make([]int, len(e))
		for i := range e { // получаем слайс интов, который уже можно передать в Solution1
			f[i], _ = strconv.Atoi(e[i])
		}

		g, _ := strconv.Atoi(d[1])
		datastructs[i] = *NewSol1(f, g)

	}
	solvedDatastucts := make([][]int, len(datastructs))
	for i := range solvedDatastucts {
		solvedDatastucts[i] = Solution1(datastructs[i].slice, datastructs[i].k)
	}

	solvedStrings := []string{}
	for i := range solvedDatastucts {
		solvedStrings = append(solvedStrings, fromSliceIntsToStirng(solvedDatastucts[i]))
	}

	return strings.Join(solvedStrings, ",")
}

type Sol1 struct {
	slice []int
	k     int
}

func NewSol1(nums []int, k int) *Sol1 {
	struct1 := Sol1{}
	struct1.slice = nums
	struct1.k = k
	return &struct1
}

func Solution1(A []int, K int) []int { // первое задание: СДВИГ СЛАЙСА ВПРАВО
	l := len(A)
	K = K % l
	A = append(A, A[:]...)
	A = A[l-K : 2*l-K]

	return A
}

func fromSliceIntsToStirng(nums []int) string {
	solvedStrings := []string{}

	for i := range nums {
		solvedStrings = append(solvedStrings, strconv.Itoa(nums[i]))
	}
	str := "[" + strings.Join(solvedStrings, ",") + "]"
	return str
}
