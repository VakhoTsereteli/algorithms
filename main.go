package main

import (
	"fmt"

	countingsort "algorithms/countingSort"
)

func test() (int, int, int) {
	return 1, 2, 3
}

func main() {
	nums := []int{4, 2, 2, 8, 3, 3, 1}

	sortArr := countingsort.CountingSor(nums)

	fmt.Println(sortArr)
}
