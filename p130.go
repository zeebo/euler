package main

import (
	"fmt"
	"math"
	"time"
)

func sieve() {
	primeSieve[0], primeSieve[1] = 1, 1 //not primes

	bound := int(math.Ceil(math.Sqrt(size)))
	for p := 2; p < bound; {
		for j := 2 * p; j < len(primeSieve); j += p {
			primeSieve[j] = 1 //j is divisible by p
		}
		for p++; primeSieve[p] == 1 && p < bound; p++ {
		}
	}
}

const size = 100000

var primeSieve = make([]int, size)

func init() {
	sieve()
}

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

func sum(x []int) (r int64) {
	for _, v := range x {
		r += int64(v)
	}
	return
}

func main() {
	s := time.Now()
	var comps []int
	for i := 91; len(comps) < 25; i += 2 {
		if i%5 == 0 || primeSieve[i] == 0 || (i-1)%A(i) != 0 {
			continue
		}
		comps = append(comps, i)
	}
	fmt.Println(comps, sum(comps), time.Since(s))
}
