package main

import (
	"fmt"

	"github.com/air-3/go-structure/linked"
)

func main() {

	req := []interface{}{1, 23, 4, 5, 1}
	head := linked.New(req)
	// fmt.Println(head)
	fmt.Println(head.Len())

	// var data interface{} = 10000
	// head.Insert(1, data)

	err := head.Delete(9)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(head.Len())

	head.PrintfNode()
}
