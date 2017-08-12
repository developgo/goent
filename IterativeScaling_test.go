package goent_test

import (
	"math"
	"testing"

	"github.com/kzahedi/goent"
	stat "gonum.org/v1/gonum/stat"
)

// alphabet         [][]int

func TestIterativeScalingAND(t *testing.T) {
	split := goent.IterativeScaling{}

	split.Nr_of_variables = 3
	split.Nr_of_states = 2

	split.P_target = make([]float64, 8, 8)
	split.P_target[0] = 0.5
	split.P_target[4] = 0.25
	split.P_target[6] = 0.25

	split.Features = make(map[string][]int)
	split.Features["X,Z"] = []int{0, 2}
	split.Features["Y,Z"] = []int{1, 2}
	split.Features["X,Y"] = []int{0, 1}

	split.Init()
	for i := 0; i < 100; i++ {
		split.Iterate()
	}

	r := stat.KullbackLeibler(split.P_target, split.P_estimate) / math.Log(2)

	if r > 0.0001 {
		t.Errorf("AND should be 0 but it is ", r)
	}
}

func TestIterativeScalingXOR(t *testing.T) {
	split := goent.IterativeScaling{}

	split.Nr_of_variables = 3
	split.Nr_of_states = 2

	split.P_target = make([]float64, 8, 8)
	split.P_target[0] = 0.25
	split.P_target[3] = 0.25
	split.P_target[5] = 0.25
	split.P_target[6] = 0.25

	split.Features = make(map[string][]int)
	split.Features["X,Z"] = []int{0, 2}
	split.Features["Y,Z"] = []int{1, 2}
	// split.Features["X,Y"] = []int{0, 1}

	split.Init()
	for i := 0; i < 1000; i++ {
		split.Iterate()
	}

	r := stat.KullbackLeibler(split.P_target, split.P_estimate) / math.Log(2)

	if math.Abs(r-1.0) > 0.00001 {
		t.Errorf("XOR should be 1 but it is ", r)
	}
}

func TestCheckFeature(t *testing.T) {
	alph := []int{10, 20, 30, 40, 50}
	a := []int{10, 20, 30, 40, 50}
	f := []int{1, 3}

	if goent.Check_feature_alphabet(f, a, alph) == false {
		t.Errorf("Check_feature_alphabet(", f, ",", a, ",", alph, ") should be true")
	}

	a = []int{100, 200, 300, 400, 500}
	g := []int{1, 2}
	if goent.Check_feature_alphabet(g, a, alph) == true {
		t.Errorf("Check_feature_alphabet(", g, ",", a, ",", alph, ") should be false")
	}
}

func TestGetAlphabetIndices(t *testing.T) {
	alphabet := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{1, 9, 3},
		{7, 8, 9},
		{1, 0, 3},
		{1, 5, 3}}
	f := []int{0, 2}
	index := 0

	indices := goent.Get_alphabet_indices(index, f, &alphabet)

	if len(indices) != 4 {
		t.Errorf("Get_alphabet_indices should return 4 values, but only has ", len(indices))
	}

	if indices[0] != 0 {
		t.Errorf("First index should be 0")
	}

	if indices[1] != 2 {
		t.Errorf("Second index should be 2")
	}

	if indices[2] != 4 {
		t.Errorf("Third index should be 4")
	}

	if indices[3] != 5 {
		t.Errorf("Fourth index should be 5")
	}

}
