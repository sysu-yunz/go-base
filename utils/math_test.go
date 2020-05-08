package utils

import (
	"math"
	"testing"
)

func TestMaxFloatSlice(t *testing.T) {
	var maxFloatTests = []struct {
		in []float64
		out float64
	}{
		{[]float64{0, 2.0, 3.0, 4.0, 5.0}, 5.0},
		{[]float64{math.Pi, math.E, 3}, math.Pi},
	}

	for _, floatSliceTest := range maxFloatTests {
		if MaxFloatSlice(floatSliceTest.in) != floatSliceTest.out {
			t.Errorf("Max of %v not %v", floatSliceTest.in, floatSliceTest.out)
		}
	}
}
