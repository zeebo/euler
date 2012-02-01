package main

import (
	"fmt"
	"math/big"
)

const turns = 15

func main() {
	var out, mask, n int32
	total, odds := new(big.Rat), new(big.Rat)
	for out = 0; out < 1<<turns; out++ {
		mask, n = 1, 0
		odds.SetInt64(1)
		for i := 0; i < turns; i, mask = i+1, mask<<1 {
			if out&mask > 0 {
				n++
				odds.Mul(odds, big.NewRat(1, int64(i+2)))
			} else {
				odds.Mul(odds, big.NewRat(int64(i+1), int64(i+2)))
			}
		}
		if n > turns/2 {
			total.Add(total, odds)
		}
	}
	fmt.Println("odds:", total)

	//invert it and find floor
	total.Inv(total)
	fmt.Println("take floor:", total.FloatString(1))
}
