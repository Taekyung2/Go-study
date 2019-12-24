//라이브러리 접근 제어
package lib

import "fmt"


func init() {
	fmt.Println("lib init start!")
}


func CheckNum(c int32) bool {
	return c > 10
}
