package main

import "fmt"

func main() {

	//slice1 := make([]float64, 0)
	slice2 := make([]float64, 3, 5)

	fmt.Printf("slice2 len:%v cap:%v", len(slice2), cap(slice2))
	fmt.Println()
	slice2 = append(slice2, 1, 2, 3, 4)

	fmt.Printf("slice2 len:%v cap:%v", len(slice2), cap(slice2))
}

/*
当前长度3，容量5，加四个元素，期望长度为7，超过当前容量发生扩容，根据golang扩容策略当前容量翻倍为10
为提高内存利用率，扩容后的切片还要进行内存对齐，根据你最后的结果可以推测你的电脑是64位系统，10容量下的float32需要40字节，内存要向上对齐到48字节，
最后真实扩容容量为48/4=12
*/

/*

什么是64位系统？
八个字节就是8个bytes. 计算机存储量的计量单位是字节，一个字节是2进制8位，八个字节等于 8*8=64 ， 即2进制64位。
例如，双精度浮点数，double 型变量，就是八个字节长度，2进制64位。
long long int 整型变量 也是八个字节长度，2进制64位。
1 byte == 8 bits. 中文里 bit / bits 叫字元。中文里 byte / bytes叫字节。
八个字节，可以存放数据，也可以存放指令。指令为八个字节64位的计算机，称为64位机，64位系统。


*/
