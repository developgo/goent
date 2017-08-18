package goent_test

import (
	"math"
	"testing"

	"github.com/kzahedi/goent"
)

func TestEmperical1D(t *testing.T) {
	t.Log("Testing Emperical1D")
	d := []int64{0, 0, 1, 1, 2, 2, 3, 3}
	p := goent.Emperical1D(d)

	if len(p) != 4 {
		t.Errorf("Emperical1D should return a slice of length 4 and not %d", len(p))
	}

	for i, _ := range p {
		if math.Abs(p[i]-1.0/4.0) > 0.0000001 {
			t.Errorf("p[%d] should be 1/4 and not %f", i, p[i])
		}
	}
}

func TestEmperical2D(t *testing.T) {
	t.Log("Testing Emperical2D")

	d := [][]int64{
		{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
		{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4},
		{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4},
		{3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4}}

	p := goent.Emperical2D(d)

	if len(p) != 4 {
		t.Errorf("Emperical2D number of rows should be 4 but it is %d", len(p))
	}
	if len(p[0]) != 5 {
		t.Errorf("Emperical2D number of columns should be 5 but it is %d", len(p[0]))
	}

	for r := 0; r < 4; r++ {
		for c := 0; c < 5; c++ {
			if math.Abs(p[r][c]-1.0/20.0) > 0.0000001 {
				t.Errorf("p[%d][%d] should be 1/20 and not %f", r, c, p[r][c])
			}
		}
	}
}

func TestEmperical3D(t *testing.T) {
	t.Log("Testing Emperical2D")

	d := [][]int64{
		{0, 0, 0}, {0, 0, 1}, {0, 0, 2}, {0, 0, 3},
		{0, 1, 0}, {0, 1, 1}, {0, 1, 2}, {0, 1, 3},
		{0, 2, 0}, {0, 2, 1}, {0, 2, 2}, {0, 2, 3},
		{1, 0, 0}, {1, 0, 1}, {1, 0, 2}, {1, 0, 3},
		{1, 1, 0}, {1, 1, 1}, {1, 1, 2}, {1, 1, 3},
		{1, 2, 0}, {1, 2, 1}, {1, 2, 2}, {1, 2, 3}}

	p := goent.Emperical3D(d)

	if len(p) != 2 {
		t.Errorf("Emperical3D 1st dimension should be 2 but it is %d", len(p))
	}
	if len(p[0]) != 3 {
		t.Errorf("Emperical3D 2nd dimension should be 3 but it is %d", len(p[0]))
	}
	if len(p[0][0]) != 4 {
		t.Errorf("Emperical3D 3rd dimension should be 4 but it is %d", len(p[0][0]))
	}

	for a := 0; a < 2; a++ {
		for b := 0; b < 3; b++ {
			for c := 0; c < 4; c++ {
				if math.Abs(p[a][b][c]-1.0/24.0) > 0.0000001 {
					t.Errorf("p[%d][%d][%d] should be 1/24 and not %f", a, b, c, p[a][b][c])
				}
			}
		}
	}
}