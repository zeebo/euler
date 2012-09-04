package main

import (
	"fmt"
	"math"
)

func sieve() {
	for i := range primeSieve {
		primeSieve[i] = 1
	}
	primeSieve[0], primeSieve[1] = 0, 0 //not primes

	bound := int(math.Ceil(math.Sqrt(size)))
	for p := 2; p < bound; {
		primes = append(primes, p)
		for j := 2 * p; j < len(primeSieve); j += p {
			primeSieve[j] = 0 //j is divisible by p
		}
		for p++; primeSieve[p] == 0 && p < bound; p++ {
		}
	}
	for p := primes[len(primes)-1] + 1; p < len(primeSieve); p++ {
		if primeSieve[p] == 1 {
			primes = append(primes, p)
		}
	}
}

func countPrimes(x []int) int {
	count := 0
	for _, v := range x {
		if primeSieve[v] == 1 {
			count++
		}
	}
	return count
}

const size = 900000

var primeSieve = make([]int, size)
var primes = make([]int, 0, size)

func init() {
	sieve()
}

func d(i int) int {
	return 6 * (i - 1)
}

func n(i int) int64 {
	x := int64(i)
	return 3*x*x - 3*x + 2
}

func diffs(i int, x int) (int64, []int) {
	if x == 0 {
		return n(i), []int{d(i+1) + 1, d(i + 1), d(i+2) + d(i+1) - 1, d(i+1) - 1, d(i), 1}
	}
	return n(i+1) - 1, []int{d(i+1) - 1, d(i + 2), d(i+2) - 1, 1, d(i + 1), d(i+1) + d(i) - 1}
}

func main() {
	items := []int64{1, 2}
	for ring := 2; len(items) < 2000; ring++ {
		for i := 0; i < 2; i++ {
			n, vals := diffs(ring, i)
			if countPrimes(vals) == 3 {
				items = append(items, n)
			}
		}
	}
	fmt.Println(items[9], items[1999])
}
