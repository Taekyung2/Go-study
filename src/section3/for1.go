package main

import "fmt"

func main() {
	// 예제 1
	for i:= 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}
	/*
	for {
		fmt.Println("ex2 : Hello, Golang!")
		fmt.Println("ex2 : Infinite loop!") 
	}
	*/
	loc := []string{"Seoul", "Busan", "Incheon"}
	for _, name := range loc {
		fmt.Println("ex3 : ", name)
	}
}