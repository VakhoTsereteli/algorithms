package main

import (
	"algorithms/bst"
)

func main() {
	bt := bst.New()

	bt.Insert(5)
	bt.Insert(10)
	bt.Insert(1)
	bt.Insert(-23)
	bt.Insert(122)

}
