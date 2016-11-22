package mymath

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Min(xs []float64) float64 {
	lowest := xs[0]
	for _, x := range xs[1:] {
		if x < lowest {
			lowest = x
		}
	}
	return lowest
}

func Max(xs []float64) float64 {
	greatest := xs[0]
	for _, x := range xs[1:] {
		if x > greatest {
			greatest = x
		}
	}
	return greatest
}
