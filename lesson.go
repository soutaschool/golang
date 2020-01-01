package main

import (
	"fmt"
)

type MyInt int 
// 自分のものに置き換える
func (i MyInt) Double() int {
	return int(i * 2)
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}