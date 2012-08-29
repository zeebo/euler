package main

import (
	"fmt"
)

var r []int

func layers(a, b, c int, n int) []int {
	diff := 4 * (a + b + c)
	layer := 2 * (a*b + b*c + c*a)
	r = r[:0]
	for layer < n {
		r = append(r, layer)
		layer, diff = layer+diff, diff+8
	}
	return r
}

//estimated guess
const top = 20000

var C = make([]int, top)

func main() {
	for a := 1; a < top/4; a++ {
		for b := a; b < top/4; b++ {
			for c := b; 2*(a*b+b*c+c*a) < top; c++ {
				for _, v := range layers(a, b, c, top) {
					C[v]++
				}
			}
		}
	}
	for k, v := range C {
		if v == 1000 {
			fmt.Println(k)
			return
		}
	}
}
