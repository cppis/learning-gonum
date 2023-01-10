package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat/distuv"
)

func main() {
	// N: 시행 10 번
	// P: 확률 0.1
	// x: 성공 횟수
	b := distuv.Binomial{N: 10, P: 0.1}
	got := b.Prob(1)
	fmt.Printf("succeed with value got %+v, binomial %+v\n", got, b)
}
