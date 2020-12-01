package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world 你好 hello"
	// 1. 判断两个utf-8编码字符串（将unicode大写、小写、标题三种格式字符视为相同）是否相同
	bool := strings.EqualFold(str, "Hello World 您好")
	fmt.Printf("EqualFold: %v \n", bool)

	// 2. 判断s是否有前缀字符串prefix
	bool = strings.HasPrefix(str, "h")
	fmt.Printf("HasPrefix: %v \n", bool)

	// 3. 判断s是否有后缀字符串suffix。
	bool = strings.HasSuffix(str, "好")
	fmt.Printf("HasSuffix: %v \n", bool)

	// 4. 判断字符串s是否包含子串substr。
	bool = strings.Contains(str, "world")
	fmt.Printf("Contains: %v \n", bool)

	// 5. 判断字符串s是否包含utf-8码值r。
	strRune := 'a'
	bool = strings.ContainsRune(str, strRune)
	fmt.Printf("ContainsRune: %v \n", bool)

	// 6. 判断字符串s是否包含字符串chars中的任一字符。
	bool = strings.ContainsAny(str, "m & p")
	fmt.Printf("ContainsAny: %v \n", bool)

	// 7. 返回字符串s中有几个不重复的sep子串。
	fmt.Printf("Count: %v \n", strings.Count(str, "hello"))

	// 8. 子串sep在字符串s中第一次出现的位置，不存在则返回-1。
	fmt.Printf("Index: %v \n", strings.Index(str, "world"))

	// 9. 字符c在s中第一次出现的位置，不存在则返回-1。
	fmt.Printf("IndexByte: %v \n", strings.IndexByte(str, 'w'))

	// 10. 子串sep在字符串s中最后一次出现的位置，不存在则返回-1。
	fmt.Printf("LastIndex: %v \n", strings.LastIndex(str, "hello"))

	// 11. 返回将所有字母都转为对应的小写版本的拷贝。
	fmt.Printf("ToLower: %v \n", strings.ToLower("HELLo"))

	// 12. 返回将所有字母都转为对应的大写版本的拷贝。
	fmt.Printf("ToUpper: %v \n", strings.ToUpper("hello"))

	// 13. 返回count个s串联的字符串。
	fmt.Printf("Repeat: %v \n", strings.Repeat("hello ", 2))

	// 14. 返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
	fmt.Printf("Replace: %v \n", strings.Replace("hello hello", "o", "k", -1))

	// 15. 返回将s前后端所有cutset包含的utf-8码值都去掉的字符串。
	fmt.Printf("Trim: %v \n", strings.Trim("hellh", "h"))

	// 16. 返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串。
	fmt.Printf("TrimSpace: %v \n", strings.TrimSpace(" hello "))

	// 17. 返回将s前端所有cutset包含的utf-8码值都去掉的字符串。
	fmt.Printf("TrimLeft: %v \n", strings.TrimLeft("hello", "he"))

	// 18. 返回将s后端所有cutset包含的utf-8码值都去掉的字符串。
	fmt.Printf("TrimRight: %v \n", strings.TrimRight("hello", "lo"))


	// 19. 返回去除s可能的后缀suffix的字符串。
	fmt.Printf("TrimSuffix: %v \n", strings.TrimSuffix("Hello, goodbye, etc!", "etc"))

	// 20. 返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串。如果字符串全部是空白或者是空字符串的话，会返回空切片。
	fmt.Printf("Fields: %q\n", strings.Fields("  foo bar  baz   "))

	// 21. 用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。
	fmt.Printf("Split %q\n", strings.Split("a,b,c", ","))

	// 22. 将一系列字符串连接为一个字符串，之间用sep来分隔。
	s := []string{"foo", "bar", "baz"}
	fmt.Printf("Join: %s\n",strings.Join(s, ", "))

	// 23. 
}
