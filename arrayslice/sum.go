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

func SumAllTails(nss ...[]int) []int {
	l := len(nss)
	sums := make([]int, l, l)

	for i, ns := range nss {
		l := len(ns)
		switch l {
		case 0:
			sums[i] = 0
		default:
			sums[i] = Sum(ns[1:]...)
		}
	}
	return sums
}
