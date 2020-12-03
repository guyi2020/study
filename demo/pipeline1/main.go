package main

import "fmt"

// 管道

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	/*
		go func() {
			for x :=0 ; ;x++ {
				naturals <- x
			}
		}()

	*/

	/*
		情况1
			go func() {
				for {
					x := <- naturals
					squares <- x * x
				}
			}()
	*/

	/*
		情况2
			go func() {
				for {
					x, ok := <- naturals
					if !ok {
						break // 通道关闭并且读完
					}
					squares <- x * x
				}
				close(squares)
			}()
	*/

	// 情况3
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	//for {
	//	fmt.Println(<-squares)
	//}

	for x := range squares {
		fmt.Println(x)
	}
}
