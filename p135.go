package main

import (
	"fmt"
	"math"
)

//a million 1-index based ints!
var counts [1e6 + 1]int

func f(x, m int64) int64 {
	return x*x - (x-m)*(x-m) - (x-2*m)*(x-2*m)
}

func lower(m int64) int64 {
	s := 2 * math.Sqrt(float64(m*m-500*500))
	return int64(math.Ceil(s)) + 3*m
}

func upper(m int64) int64 {
	s := 2 * math.Sqrt(float64(m*m-500*500))
	return 3*m - int64(s)
}

func main() {
	var m, x int64

	//easy pass
	for m = 500; m > 0; m-- {
		for x = 2*m + 1; x < 5*m; x++ {
			counts[f(x, m)]++
		}
	}

	//middle pass
	for m = 501; m <= 577; m++ {
		//first section
		for x = 2*m + 1; x < upper(m); x++ {
			counts[f(x, m)]++
		}

		//second section
		for x = lower(m); x < 5*m; x++ {
			counts[f(x, m)]++
		}
	}

	//upper pass
	for m = 578; m < 1e6/4+1; m++ {
		for x = lower(m); x < 5*m; x++ {
			counts[f(x, m)]++
		}
	}

	//total it up
	total := 0
	for _, v := range counts {
		if v == 10 {
			total++
		}
	}

	fmt.Println(total) //4989
}
