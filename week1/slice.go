package main

import "fmt"

/* 实现删除切片特定下标元素的方法。

要求一：能够实现删除操作就可以。
要求二：考虑使用比较高性能的实现。 (e.g.: 减少内存分配, 资料复制)
要求三：改造为泛型方法
要求四：支持缩容，并旦设计缩容机制。 (e.g.: 删除一定数量再缩容)
**/

const (
	ShrinkThreshold = 0.75
)

func main() {
	deleteFn := DeleteV3[int]
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("original slice: %v, cap %d, len %d, position: %p, position of first element: %p\n", slice, cap(slice), len(slice), slice, &(slice[0]))

	i := 0
	slice = testDelete[int](deleteFn, slice, i)

	i = 1
	slice = testDelete[int](deleteFn, slice, i)

	i = len(slice) - 2
	slice = testDelete[int](deleteFn, slice, i)

	i = len(slice) - 1
	slice = testDelete[int](deleteFn, slice, i)

	// 测试缩容
	for i := len(slice) - 1; i >= 0; i-- {
		slice = testDelete(deleteFn, slice, i)
	}

	fmt.Printf("append after being assigned as nil: %v", append(slice, 1))
}

func testDelete[T any](deleteFn func([]T, int) []T, slice []T, i int) []T {
	slice = deleteFn(slice, i)
	if len(slice) > 0 {
		fmt.Printf("delete %d: result slice: %v, cap %d, len %d, position: %p, position of first element: %p\n", i, slice, cap(slice), len(slice), slice, &(slice[0]))
	} else {
		fmt.Printf("delete %d: result slice: %v, cap %d, len %d, position: %p\n", i, slice, cap(slice), len(slice), slice)
	}
	return slice
}

func DeleteV1[T any](slice []T, i int) []T {
	if i == len(slice)-1 {
		return slice
	}
	return append((slice)[:i], (slice)[i+1:]...)
}

// 依照元素位置处理, 减少资料复制
func DeleteV2[T any](slice []T, i int) []T {
	if i > len(slice) {
		panic("index out of range")
	}

	mid := len(slice) / 2
	if i > mid {
		// 第一个元素位置不变时, capacity不会减少
		for j := i; j < len(slice)-1; j++ {
			slice[j] = slice[j+1]
		}
		slice = slice[:len(slice)-1]
	} else {
		// 第一个元素位置改变时, capacity会减少, 应该就算缩容？？ 没有被任何slice reference到就会被GC回收？？
		for j := i; j > 0; j-- {
			slice[j] = slice[j-1]
		}
		slice = slice[1:]
	}
	return slice
}

// capacity 与 length 差距达threshold时, 主动缩容
func DeleteV3[T any](slice []T, i int) []T {

	if i > len(slice) {
		panic("index out of range")
	}

	// 主动缩容
	if len(slice) == 1 {
		return nil
	}
	if float64(len(slice))/float64(cap(slice)) <= ShrinkThreshold {
		newSlice := make([]T, len(slice)-1, len(slice)-1)
		copy(newSlice, slice[:i])
		if i+1 < len(slice) {
			copy(newSlice[i:], slice[i+1:])
		}
		return newSlice
	}

	mid := len(slice) / 2
	if i > mid {
		// 第一个元素位置不变时, capacity不会减少
		for j := i; j < len(slice)-1; j++ {
			slice[j] = slice[j+1]
		}
		slice = slice[:len(slice)-1]
	} else {
		// 第一个元素位置改变时, capacity会减少, 应该就算缩容？？ 没有被任何slice reference到就会被GC回收？？
		for j := i; j > 0; j-- {
			slice[j] = slice[j-1]
		}
		slice = slice[1:]
	}
	return slice
}

/* TODO:
1. 缩容时可保留一定的capacity供未来append使用
2. capacity 少于一定数量e.g. 64时可不用缩容
*/
