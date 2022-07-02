package main

import (
	"fmt"
	"reflect"
)

//func main() {
//
//	str1 := "golang"
//	str2 := "go语言"
//
//	//在go语言中，英文字符占一个字节，中文字符占三个字节
//	fmt.Println(reflect.TypeOf(str2[2]).Kind())
//	fmt.Println(str1[2], string(str1[2]))
//	//这里指的是下标为2处的字节，不是字符，%c表示的是输出单个字符
//	fmt.Printf("%d %c\n", str2[2], str2[2])
//
//	fmt.Println("len(Str2),", len(str2))
//}

/*
reflect.TypeOf().Kind() 可以知道某个变量的类型，我们可以看到，字符串是以 byte 数组形式保存的，类型是 uint8，占1个 byte，打印时需要用 string 进行类型转换，否则打印的是编码值。
因为字符串是以 byte 数组的形式存储的，所以，str2[2] 的值并不等于语。str2 的长度 len(str2) 也不是 4，而是 8（ Go 占 2 byte，语言占 6 byte）。

*/

/*
下面来看看中文字符串的正确处理方式
*/

func main() {

	str1 := "golang"
	str2 := "go语言"

	//在go语言中，英文字符占一个字节，中文字符占三个字节
	fmt.Println(reflect.TypeOf(str2[2]).Kind())
	fmt.Println(str1[2], string(str1[2]))
	//这里指的是下标为2处的字节，不是字符，%c表示的是输出单个字符
	fmt.Printf("%d %c\n", str2[2], str2[2])

	fmt.Println("len(Str2),", len(str2))

	runeArr := []rune(str2)

	fmt.Println(reflect.TypeOf(runeArr[2]).Kind())

	fmt.Println(runeArr[2], string(runeArr[2]))

	fmt.Println("len(runeARR),", len(runeArr))
}

/*

转换成 []rune 类型后，字符串中的每个字符，无论占多少个字节都用 int32 来表示，因而可以正确处理中文

*/
