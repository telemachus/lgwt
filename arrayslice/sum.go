package arrayslice

func Sum(ns ...int) int {
	var total int
	for _, n := range ns {
		total += n
	}

	return total
}

func SumAll(nss ...[]int) []int {
	l := len(nss)
	sums := make([]int, l, l)

	for i, ns := range nss {
		sums[i] = Sum(ns...)
	}
	return sums
}
