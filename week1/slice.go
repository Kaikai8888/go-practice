package main

import "fmt"

/* 实现删除切片特定下标元素的方法。

要求一：能够实现删除操作就可以。
要求二：考虑使用比较高性能的实现。
要求三：改造为泛型方法
要求四：支持缩容，并旦设计缩容机制。
**/

func main() {
	deleteFn := SimpleDelete[int]
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("original slice: %v, position: %p\n", slice, slice)

	i := 1
	slice = deleteFn(slice, i)
	fmt.Printf("delete %d: result slice: %v, position: %p\n", i, slice, slice)

	i = 0
	slice = deleteFn(slice, i)
	fmt.Printf("delete %d: result slice: %v, position: %p\n", i, slice, slice)

	i = len(slice) - 2
	slice = deleteFn(slice, i)
	fmt.Printf("delete %d: result slice: %v, position: %p\n", i, slice, slice)

	i = len(slice) - 1
	slice = deleteFn(slice, i)
	fmt.Printf("delete %d: result slice: %v, position: %p\n", i, slice, slice)
}

func SimpleDelete[T any](slice []T, i int) []T {
	if i == len(slice)-1 {
		return slice
	}
	slice = append((slice)[:i], (slice)[i+1:]...)
	return slice
}

func AdvancedDelete[T any](slice []T, i int) []T {
	if i > len(slice) {
		panic("index out of range")
	}

	mid := len(slice) / 2
	if i > mid {
		for j := i; j < len(slice)-1; j++ {
			slice[j], slice[j+1] = slice[j+1], slice[j]
		}
		slice = slice[:len(slice)-1]
	} else {
		for j := i; j > 0; j-- {
			slice[j], slice[j-1] = slice[j-1], slice[j]
		}
		slice = slice[1:]
	}

	return slice
}
