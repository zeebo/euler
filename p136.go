package main

import (
	"fmt"
	"math"
)

//loop bounds based on
//http://www.wolframalpha.com/input/?i=x%5E2+-+%28x-n%29%5E2+-+%28x-2n%29%5E2+%3E+0%2C+x%5E2+-+%28x-n%29%5E2+-+%28x-2n%29%5E2+%3C%3D+50000000%2C+x+%3E+2n%2C+n+%3E+0

//50 million 1-index based ints!
var counts [50e6 + 1]int

func f(x, m int64) int64 {
	return x*x - (x-m)*(x-m) - (x-2*m)*(x-2*m)
}

func lower(m int64) int64 {
	s := 2 * math.Sqrt(float64(m*m-1.25e7))
	return int64(math.Ceil(s)) + 3*m
}

func upper(m int64) int64 {
	s := 2 * math.Sqrt(float64(m*m-1.25e7))
	return 3*m - int64(s)
}

func main() {
	var m, x int64

	//easy pass
	for m = 3535; m > 0; m-- {
		for x = 2*m + 1; x < 5*m; x++ {
			counts[f(x, m)]++
		}
	}

	//middle pass
	for m = 3536; m <= 4082; m++ {
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
	for m = 4083; m < 50e6/4+1; m++ {
		for x = lower(m); x < 5*m; x++ {
			counts[f(x, m)]++
		}
	}

	//total it up
	total := 0
	for _, v := range counts {
		if v == 1 {
			total++
		}
	}

	fmt.Println(total) //2544559
}
