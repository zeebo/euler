package main

import (
	"fmt"
	"math"
)

func sieve(x []int) []int {
	primes := make([]int, 0, len(x))
	for i := range x {
		x[i] = 1
	}
	x[0], x[1] = 0, 0 //not primes

	bound := int(math.Ceil(math.Sqrt(float64(len(x)))))
	for p := 2; p < bound; {
		primes = append(primes, p)
		for j := 2 * p; j < len(x); j += p {
			x[j] = 0 //j is divisible by p
		}
		for p++; x[p] == 0 && p < bound; p++ {
		}
	}
	for p := primes[len(primes)-1] + 1; p < len(x); p++ {
		if x[p] == 1 {
			primes = append(primes, p)
		}
	}
	return primes
}

const size = 120000

var primeSieve = make([]int, size)
var rcache = make([]uint64, size)
var primes []int

func init() {
	primes = sieve(primeSieve)
	fmt.Println("primes sieved")

	//precompute radicals
	for i := 1; i < size; i++ {
		radical(rcache, i)
	}
	fmt.Println("radicals precomputed")
}

func radical(rcache []uint64, x int) {
	y := x

	product := 1
	for i := 0; i < len(primes) && primes[i] <= x; i++ {
		p := primes[i]
		if x%p == 0 {
			product *= p
			for x%p == 0 {
				x /= p
			}
		}
	}

	rcache[y] = uint64(product)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	sum, n := uint64(0), uint64(0)
	for c := 1; c < size; c++ {
		if 2*rcache[c] >= uint64(c) {
			continue
		}
		for a := 1; a <= c/2; a++ {
			b := c - a
			if gcd(a, b) == 1 && rcache[a]*rcache[b]*rcache[c] < uint64(c) {
				// fmt.Println(a, b, c, rcache[a], rcache[b], rcache[c])
				n++
				sum += uint64(c)
			}
		}
		if c%10 == 0 {
			fmt.Println(c)
		}
	}
	fmt.Println(n, sum)
}
