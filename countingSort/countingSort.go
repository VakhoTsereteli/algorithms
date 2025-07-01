package countingsort

func CountingSort(input []int) []int {
	minNum := input[0]
	maxNum := input[0]
	for _, num := range input {
		if num < minNum {
			minNum = num
		}
	}
	for _, num := range input {
		if num > maxNum {
			maxNum = num
		}
	}

	countArr := make([]int, (maxNum-minNum)+1)

	for i := range input {
		countArr[input[i]-minNum] += 1
	}

	for i := 1; i < len(countArr); i++ {
		countArr[i] = countArr[i] + countArr[i-1]
	}

	outputArr := make([]int, len(input))

	for i := len(input) - 1; i >= 0; i-- {
		currentNum := input[i]
		outputIndex := countArr[currentNum-minNum] - 1
		outputArr[outputIndex] = currentNum

		countArr[currentNum-minNum]--
	}

	return outputArr
}

