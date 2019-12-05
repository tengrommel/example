package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

var (
	a []int
	b = []int{}
	c = []int{1, 2, 3}
	d = c[:2]
	e = c[0:2:cap(c)]
	f = c[:0]
	g = make([]int, 3)
	h = make([]int, 2, 3)
	i = make([]int, 0, 3)
)

func TrimSpace(s []byte) []byte {
	b := s[:0]
	for _, x := range s {
		if x != ' ' {
			b = append(b, x)
		}
	}
	return b
}

func Filter(s []byte, fn func(x byte) bool) []byte {
	b := s[:0]
	for _, x := range s {
		if !fn(x) {
			b = append(b, x)
		}
	}
	return b
}

var aSlice = []float64{4, 2, 5, 7, 2, 1, 88, 1}

func SortFloat64FastV1(a []float64) {
	// 强制类型转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
	// 以int方式给float64排序
	sort.Ints(b)
}

func SortFloat64FastV2(a []float64) {
	// 通过reflect.SliceHeader更新切片头部信息实现转换
	var c []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	*cHdr = *aHdr
	// 以int方式给float64排序
	sort.Ints(c)
}

func main() {
	for i := range a {
		fmt.Printf("a[%d] : %d\n", i, a[i])
	}
	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]: %d\n", i, c[i])
	}
	var a []int
	a = append(a, 1)                 // 追加一个元素
	a = append(a, 1, 2, 3)           // 追加多个元素，手写解包方式
	a = append(a, []int{1, 2, 3}...) // 追加一个切片，切片需要解包

	var ab = []int{1, 2, 3}
	ab = append([]int{0}, ab...)          // 从开头添加一个元素
	ab = append([]int{-3, -2, -1}, ab...) // 在开头添加一个切片

	i := 1
	ab = append(a[:i], append([]int{1}, a[i:]...)...)    // 在第i个位置插入1
	ab = append(a[:i], append([]int{1, 2}, a[i:]...)...) // 在第i个位置插入切片

	// 在开头一般都会导致内存的重新分配，而且会导致已有的元素全部复制一次。
	N := 2
	a = []int{1, 2, 3}
	a = a[:len(a)-1]
	a = a[:len(a)-N]
	// 删除开头的元素可以直接移动数据指针：
	a = []int{1, 2, 3}
	a = append(a[:0], a[1:]...) // 删除开头1个元素
	a = append(a[:0], a[N:]...) // 删除开头N个元素

}
