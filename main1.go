package main

import (
	"fmt"
	"sort"
)

func main() {

	nums1 := []int{1, 2, 3, 4, 5}
	k := 3
	fmt.Println(nums1)
	fmt.Println(Solution1(nums1, k))

	nums2 := []int{4, 15, 18, 3, 15, 4, 18, 19, 19}
	fmt.Println(Solution2(nums2))

	nums3 := []int{4, 7, 6, 5, 8, 10}
	fmt.Println(Solution3(nums3))

	nums4 := []int{2, 3, 1, 5}
	fmt.Println(Solution4(nums4))

}

func Solution1(A []int, K int) []int { // первое задание: СДВИГ СЛАЙСА ВПРАВО
	l := len(A)
	K = K % l
	A = append(A, A[:]...)
	A = A[l-K : 2*l-K]

	return A
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
