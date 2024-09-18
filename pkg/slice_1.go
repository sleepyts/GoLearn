package pkg

import "fmt"

func SliceTest() {
	// 动态数组切片
	dyinamicArr := []int{1, 2, 3, 4, 5}
	PrintArray(dyinamicArr)
	slice1 := make([]int, 3)
	if slice1 == nil {
		print("slice1 is nil")
	} else {
		PrintArray(slice1)
	}
	slice2 := make([]int,3,5)
	slice2=append(slice2, 6, 7)
	PrintArray(slice2)
	// 容量超过后2倍扩容
	slice2=append(slice2, 8)
	PrintArray(slice2)
	// 切片截取 截取的是同一个内存地址
	slice3 := slice2[0:10]
	PrintArray(slice3)
	// 切片拷贝 深拷贝 如果长度不一样，会从头覆盖
	copy(slice3, []int{9, 10, 11})
	PrintArray(slice3)
}

// 引用传递 arr
func PrintArray(arr []int) {
	fmt.Printf("arr: %v, len: %d, cap: %d\n", arr, len(arr), cap(arr))
}