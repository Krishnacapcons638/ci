package main

import "testing"

type input struct {
	num1 int
	num2 int
}

func TestAdd(t *testing.T) {
	testcases := []struct {
		name  string
		inp   input
		expec int
	}{
		{
			name: "success",
			inp: input{
				num1: 1,
				num2: 2,
			},
			expec: 3,
		}, {
			name: "success",
			inp: input{
				num1: 1,
				num2: 5,
			},
			expec: 6,
		},
	}

	for _, test := range testcases {
		exact := Add(test.inp.num1, test.inp.num2)
		t.Run(test.name, func(t *testing.T) {
			if exact != test.expec {
				t.Error("bekar hain bhaiya main toh tootha gaya")
			}
		})
	}
}
