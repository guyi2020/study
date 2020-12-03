package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 1. 字符串转整形
	i, err := strconv.Atoi("-10")
	fmt.Printf("i: %d , err : %v \n", i, err)

	// 2. 整形转字符串
	s := strconv.Itoa(100)
	fmt.Printf("s: %T, %s \n", s, s)

	// 3. 字符串转其他数据类型
	if b, err := strconv.ParseBool("false"); err == nil {
		fmt.Printf("b: %T, %t \n", b, b)
	}

	if f, err := strconv.ParseFloat("3.14159", 64); err == nil {
		fmt.Printf("f: %T, %v \n", f, f)
	}

	// 4. 将 s 转换为双引号字符串
	s = strconv.Quote(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)

	// 5. bool 转 字符串格式
	s = strconv.FormatBool(true)
	fmt.Printf("%T, %v\n", s, s)
}
