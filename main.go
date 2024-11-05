// 第一个 Go
package main

import (
	"fmt"
)

// 全局定义数组
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3}

func main() {
	arr3 := []int{1, 2}
	arr4 := []int{3, 4}

	fmt.Println("hello Go")

	// 数组操作
	// len 方法获取数组长度
	fmt.Println("arra1 长度 ===>", len(arr1))
	fmt.Println("arra2 长度 ===>", len(arr2))

	// cap 获取数组容量, 只有在数组的底层数组是可寻址的情况下才能使用
	fmt.Println("arr2 容量 ===>", cap(arr2))

	// copy, 只能切片使用数组不能使用
	fmt.Println("copy ===>", copy(arr3, arr4))
	fmt.Println("arra3 长度 ===>", arr3)
	fmt.Println("arra4 长度 ===>", arr4)

	// 切片
	// 1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
	// 2. 切片的长度可以改变，因此，切片是一个可变的数组。
	// 3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
	// 4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
	// 5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
	// 6. 如果 slice == nil，那么 len、cap 结果都等于 0。

	// 定义一个切片
	var slice1 = []int{1, 2, 3}
	var slice2 = append(slice1, 4)
	var slice3 = arr1[1:]
	var slice4 = arr1[2:4]
	var slice5 = arr1[2:4:5]

	fmt.Println("slice1 ===>", slice1)
	fmt.Println("slice2 ===>", slice2)
	fmt.Println("slice3 ===>", slice3)
	fmt.Println("slice4 ===>", slice4, cap(slice4))
	fmt.Println("slice5 ===>", slice5, cap(slice5))

	// 每隔1秒执行一次
	// t := time.Tick(1 * time.Second)
	// for {
	// 	<-t
	// 	currentTime := localpackage.GetCurrentTime()
	// 	fmt.Println("当前时间：", currentTime)
	// }

}
