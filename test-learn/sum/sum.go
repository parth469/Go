package sum

func Sum(list []int) int {
	sum := 0

	for _, v := range list {
		sum += v
	}

	return sum
}
