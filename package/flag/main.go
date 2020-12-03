package main

import (
	"flag"
	"fmt"
)

var inputName = flag.String("name", "CHENJIAN", "Input Your Name.")
var inputAge = flag.Int("age", 18, "Input Your Age")
var inputSex = flag.Int("sex", 0, "Input Your Sex")
var inputFlagvar int

func Init() {
	flag.IntVar(&inputFlagvar, "flagname", 123, "Help")
}
func main() {
	Init()
	flag.Parse()
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("name=", *inputName)
	fmt.Println("age=", *inputAge)
	fmt.Println("sex=", *inputSex)
	fmt.Println("flagname=", inputFlagvar)
}
