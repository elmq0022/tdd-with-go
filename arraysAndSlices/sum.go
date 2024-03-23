package arraysandslices

func Sum(numbers []int) int {
	var s int
	for _, n := range numbers {
		s += n
	}
	return s
}

func SumAll(numbers ...[]int) []int {
	n := len(numbers)
	s := make([]int, n)
	for i, n := range numbers {
		s[i] = Sum(n)
	}
	return s
}

func SumAllTails(numbers ...[]int) []int {
	var s []int
	for _, n := range numbers {
		s = append(s, Sum(n[1:]))
	}
	return s
}
