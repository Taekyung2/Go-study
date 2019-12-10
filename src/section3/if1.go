package main

import "fmt"

func main() {
	// 제어문
	// if문 -> 1, 0 사용 불가, 자동 형 변환 불가
	// 소괄호 사용 x
	
	var a int = 20
	b := 20

	// 예제 1
	if a >= 15 {
		fmt.Println("15이상")
	}
	// 예제 2
	if b >= 25 {
		fmt.Println("25이상")
	}
	if c:= 40; c >= 35 {
		fmt.Println("35이상")
	}
	
}