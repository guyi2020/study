package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// 1. 字符串读取器
	s := strings.NewReader("this is a test")
	b, err := ioutil.ReadAll(s)
	if err != nil {
		log.Fatalf("err: %s \n", err)
	}
	fmt.Printf("b: %s \n", b)

	// 2. 读取文件夹路径
	files, res := ioutil.ReadDir("D:\\go_project\\study")
	if res != nil {
		log.Fatalf("files: %s \n", res)
	}

	for _, file := range files {
		fmt.Printf("file name %s \n", file.Name())
		fmt.Printf("file IsDir %t \n", file.IsDir())
		fmt.Printf("file Mode %s \n", file.Mode())
		fmt.Printf("file Size %d \n", file.Size())
		fmt.Printf("file Sys %v \n", file.Sys())
	}

	// 3. 读取文件内容
	data, err := ioutil.ReadFile("test.log")
	if err != nil {
		log.Fatalf("ReadFile err: %s \n", err)
	}
	fmt.Printf("%s", data)

	// 4. 创建临时文件目录
	dir, err := ioutil.TempDir("D:\\go_project\\study\\package\\ioutil", "test")
	if err != nil {
		log.Fatalf("TempDir err: %s \n", err)
	}
	fmt.Printf("dir: %v ", dir)

	// 5. 创建临时文件
	file, err := ioutil.TempFile("D:\\go_project\\study\\package\\ioutil", "test")
	defer file.Close()
	if err != nil {
		log.Fatalf("TempFile err: %s \n", err)
	}
	file.WriteString("this is a content")

}
