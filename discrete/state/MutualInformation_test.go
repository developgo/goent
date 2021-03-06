package state_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/kzahedi/goent/discrete"
	"github.com/kzahedi/goent/discrete/state"
)

func TestMutualInformation(t *testing.T) {
	t.Log("Testing Mutual Information")
	data := make([][]int, 100)
	for j := 0; j < 100; j++ {
		for i := 0; i < 100; i++ {
			data[i] = make([]int, 2, 2)
			data[i][0] = int(rand.Int63n(100))
			data[i][1] = int(rand.Int63n(100))
		}

		r := state.MutualInformationBase2(data)
		p := discrete.Empirical2D(data)
		mi := discrete.MutualInformationBase2(p)

		s := 0.0
		for i := 0; i < len(r); i++ {
			s += r[i]
		}

		s /= float64(len(r))

		if math.Abs(mi-s) > 0.00001 {
			t.Errorf("Mutual information should be equal %f = %f", s, mi)
		}
	}
}

func TestMutualInformationBaseE(t *testing.T) {
	t.Log("Testing Mutual Information")
	data := make([][]int, 100)
	for j := 0; j < 100; j++ {
		for i := 0; i < 100; i++ {
			data[i] = make([]int, 2, 2)
			data[i][0] = int(rand.Int63n(100))
			data[i][1] = int(rand.Int63n(100))
		}

		r := state.MutualInformationBaseE(data)
		p := discrete.Empirical2D(data)
		mi := discrete.MutualInformationBaseE(p)

		s := 0.0
		for i := 0; i < len(r); i++ {
			s += r[i]
		}

		s /= float64(len(r))

		if math.Abs(mi-s) > 0.00001 {
			t.Errorf("Mutual information (Base E) should be equal %f = %f", s, mi)
		}
	}
}
