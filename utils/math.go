package utils

func MaxFloatSlice(l []float64) float64 {
	var max float64
	for i, f := range l {
		if i == 0 || f > max {
			max = f
		}
	}

	return max
}
