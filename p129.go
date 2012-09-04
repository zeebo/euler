package main

import (
	"fmt"
	"time"
)

func A(n int) (k int) {
	if n%2 == 0 || n%5 == 0 {
		panic("GCD(n, 10) != 0")
	}

	x := 1
	for z := 1; z != 0; k++ {
		x = (x * 10) % n
		z = (z + x) % n
		x = x % n
	}
	k++
	return
}

func main() {
	//A(i) > i for i > 1.
	s := time.Now()
	for i := 1000001; ; i += 2 {
		if i%5 == 0 {
			continue
		}
		if A(i) > 1000000 {
			fmt.Println(i, A(i), time.Since(s))
			return
		}
	}
}
