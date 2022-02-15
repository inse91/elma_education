package main

import (
	"fmt"
	"sort"
	//"strings"
)

type firstNameKey struct{}

type lastNameKey struct{}

// ctx := context.Background()

// 	firstNameCtx := context.WithValue(ctx, firstNameKey{}, "firstname")
// 	lastNameCtx := context.WithValue(ctx, lastNameKey{}, "lastname")

// 	print(firstNameCtx)
// 	print(lastNameCtx)

// func print(ctx context.Context) {
// 	val := ctx.Value(firstNameKey{})
// 	fmt.Println(val.(string))

// }

func main() {

	nums1 := []int{3, 8, 9, 7, 6}
	k1 := 3
	fmt.Println(nums1, Solution1(nums1, k1))

	nums2 := []int{4, 15, 18, 3, 15, 4, 18}
	fmt.Println(Solution2(nums2))

	nums3 := []int{3, 4, 2, 5}
	fmt.Println(Solution3(nums3))

	nums4 := []int{7, 10, 6, 9}
	fmt.Println(Solution4(nums4))

}

func Solution1(A []int, K int) []int { // первое задание: СДВИГ СЛАЙСА ВПРАВО

	for i := 0; i < K; i++ {
		y := len(A) - 1
		A = append(A, A[:y]...)
		A = A[y:]

	}
	return A
}

func Solution2(A []int) (x int) { // второе задание: ПОИСК ЭЛЕМЕНТА БЕЗ ПАРЫ
	for i := range A {
		count := 0
		for j := range A {
			if A[i] == A[j] {
				count++
			}
		}
		if count%2 != 0 {
			x = A[i]
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

func Solution4(A []int) (x int) { // четвертое задание: ПОИСК НЕДОСТАЮЩЕГО ЭЛЕМЕНТА
	// В ПОСЛЕДОВАТЕЛЬНОСТИ
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
