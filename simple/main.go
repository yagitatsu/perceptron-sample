package main

import (
	"errors"
	"fmt"
)

var (
	//x, y, 1
	inputs = [][]float64{
		{5.0, 5.0, 1},
		{0.13, 0.22, 1},
		{1.2, 5.0, 1},
		{3.5, 1.37, 1},
		{8.0, 0.0, 1},
		{2.2, 1.0, 1},
		{6.1, -0.2, 1},
		{-1.2, 5.1, 1},
		{2.3, 5.1, 1},
		{3.4, 1.2, 1},
	}
	outputs = []float64{
		1.0,
		-1.0,
		1.0,
		-1.0,
		1.0,
		-1.0,
		1.0,
		-1.0,
		1.0,
		-1.0,
	}
)

func main() {

	//a, b, c
	//ax + by + c = 0
	weight := []float64{0.0, 0.0, 0.0}

	//even if same training data is used, weight can be changed until convergence.
	for n := 0; n < 10; n++ {
		for i := range inputs {
			w, err := train(weight, inputs[i], outputs[i])
			if err != nil {
				fmt.Println(err)
				return
			}
			weight = w
		}
	}

	fmt.Println(fmt.Sprintf("weight: %v", weight))
}

func train(w []float64, in []float64, out float64) ([]float64, error) {
	//use Stochastic Gradient Descent
	ret, err := classify(w, in)
	if err != nil {
		return []float64{}, err
	}
	if ret == out {
		return w, nil
	}
	const eta = 0.1 //TODO: what does it mean...?
	delta := multiply(eta*out, in)
	return add(w, delta)

}

func classify(w []float64, in []float64) (float64, error) {
	ret, err := scoring(w, in)
	if err != nil {
		return 0.0, err
	}
	if ret >= 0.0 {
		return 1.0, nil
	}
	return -1.0, nil
}

func scoring(w []float64, in []float64) (float64, error) {

	if len(w) != len(in) {
		return 0.0, errors.New("length of w is different from length of in")
	}

	f := 0.0
	for i := range w {
		f += w[i] * in[i]
	}
	return f, nil
}

func add(a []float64, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return []float64{}, errors.New("length of a is different from length of b")
	}
	c := []float64{}
	for i := range a {
		c = append(c, a[i]+b[i])
	}
	return c, nil
}

func multiply(a float64, x []float64) []float64 {
	y := []float64{}
	for i := range x {
		y = append(y, x[i]*a)
	}
	return y
}
