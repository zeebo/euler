package main

import (
	"fmt"
	"math/big"
)

const (
	primesize = 10
	rabinwork = 10
)

var buf [primesize]byte

func bmap(x int) byte {
	return byte(x) + '0'
}

//lets just assume all 9's for now
func build9(d, x, p int) string {
	var bd, bx = bmap(d), bmap(x)
	for i := 0; i < primesize; i++ {
		buf[i] = bd
	}
	buf[p] = bx

	return string(buf[:])
}

func build8(d, x1, p1, x2, p2 int) string {
	var bd, bx1, bx2 = bmap(d), bmap(x1), bmap(x2)
	for i := 0; i < primesize; i++ {
		buf[i] = bd
	}
	buf[p1] = bx1
	buf[p2] = bx2

	return string(buf[:])
}

func main() {
	var (
		ntotal = new(big.Int)
		x      = new(big.Int)
	)
	var misses = []int{}

	fmt.Printf("Working for %d digit primes\n", primesize)
	fmt.Println("dig\tn\tpsum")

	//all these padding digits have soluitons with 9
	for digit := 0; digit < 10; digit++ {
		var (
			psum = new(big.Int)
			n    int
		)
		for j := 0; j < 10; j++ {
			for pos := 0; pos < primesize; pos++ {
				x.SetString(build9(digit, j, pos), 0)
				if len(x.String()) != primesize {
					continue
				}

				if big.ProbablyPrime(x, rabinwork) { //we have a probable prime
					ntotal.Add(x, ntotal)
					psum.Add(x, psum)
					n++
				}
			}
		}
		fmt.Printf("%d\t%d\t%d\n", digit, n, psum)
		if n == 0 {
			misses = append(misses, digit)
		}
	}

	fmt.Println("(n-1)s done...")
	fmt.Println("dig\tn\tpsum")

	//now the misses
	for _, digit := range misses {
		var (
			psum = new(big.Int)
			n    int
		)

		for j1 := 0; j1 < 10; j1++ {
			for pos1 := 0; pos1 < primesize; pos1++ {
				for j2 := 0; j2 < 10; j2++ {
					for pos2 := pos1 + 1; pos2 < primesize; pos2++ {
						x.SetString(build8(digit, j1, pos1, j2, pos2), 0)
						if len(x.String()) != primesize {
							continue
						}

						if big.ProbablyPrime(x, rabinwork) { //we have a probable prime
							ntotal.Add(x, ntotal)
							psum.Add(x, psum)
							n++
						}
					}
				}
			}
		}
		fmt.Printf("%d\t%d\t%d\n", digit, n, psum)
		if n == 0 {
			panic("Won't get the correct result.")
		}
	}

	fmt.Println("(n-2)s done...")
	fmt.Println(ntotal)
}
