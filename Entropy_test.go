package goent_test

import (
	"math"
	"math/rand"
	"testing"

	"github.com/kzahedi/goent"
)

func TestEntropy(t *testing.T) {
	t.Log("Testing Entropy")
	p1 := []float64{0.5, 0.5, 0.5, 0.5}

	if r := goent.Entropy2(p1); r != 2.0 {
		t.Errorf("Entropy of four state uniform distribution should be 2.0 but it is ", r)
	}

	p2 := []float64{1.0, 0.0, 0.0, 0.0}

	if r := goent.Entropy2(p2); r != 0.0 {
		t.Errorf("Entropy of deterministic distribution should be 0.0 but it is ", r)
	}
}

func TestEntropyChaoShen(t *testing.T) {
	t.Log("Testing Chao-Shen Entropy")
	r := 0.0
	for i := 0; i < 100; i++ {
		h := make([]int64, 5000, 5000)
		for j := 0; j < 5000; j++ {
			h[j] = rand.Int63n(100)
		}
		r += goent.EntropyChaoShen(h)
	}

	r /= 100.0

	if math.Abs(r-4.595091) > 0.1 {
		t.Errorf("Entropy should be 4.595091 and not %f", r)
	}

}

func TestEntropyMLBC(t *testing.T) {
	t.Log("Testing Maximum Likelihood Bias Corrected")
	r := 0.0
	for i := 0; i < 100; i++ {
		h := make([]int64, 5000, 5000)
		for j := 0; j < 5000; j++ {
			h[j] = rand.Int63n(100)
		}
		r += goent.EntropyMLBC(h)
	}

	r /= 100.0

	if math.Abs(r-4.604982) > 0.1 {
		t.Errorf("Entropy should be 4.604982 and not %f", r)
	}

}
