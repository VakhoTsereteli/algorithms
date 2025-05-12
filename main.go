package main

import (
	"algorithms/bst"
	"fmt"
)

func test()(int,int,int){
	return 1,2,3
}

func main() {
	bt := bst.New()
	bt.Insert(22)
	bt.Insert(25)
	bt.Insert(-5)
	bt.Insert(-22)

	btMin, _ := bt.FindMin()
	fmt.Println(btMin)

}
