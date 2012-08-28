package main

import (
	"fmt"
)

func counts(n int) int {
	// note: we can short circuit the count is n is even but it doesn't matter
	sum := 0
	max := uint16(1 << uint(n))

	for target := 2; target <= n; target += 2 {
		for i := uint16(0); i < max; i++ {
			if countBits(i) != target {
				continue
			}
			sum += numPartition(n, countBits(i), i)
		}
	}
	return sum
}

func numPartition(n int, n2 int, chosen uint16) int {
	mapping := make([]int, 0, n2) //mapping takes indicies 0 to n2 into bit positions from chosen
	for j := uint(0); j < uint(n); j++ {
		if chosen&(1<<j) > 0 {
			mapping = append(mapping, int(j))
		}
	}
	sum := 0
	max := uint16(1 << uint(n2))
	for i := uint16(0); i < max; i++ {
		if countBits(i) != n2>>1 {
			continue
		}
		//i represents which indicies go on the left
		left, right := split(mapping, i)
		left2, right2 := split(mapping, i)
		if isBad(left, right) || isBad(right2, left2) {
			continue
		}
		sum++
	}
	return sum / 2
}

func isBad(left, right []int) bool {
	if len(left) == 0 {
		return true
	}
	//try to find an item in right that works
	for i, v := range right {
		if v > 0 && v > left[0] {
			right[i] = 0
			if isBad(left[1:], right) {
				return true
			}
			right[i] = v
		}
	}
	return false
}

//split puts the items in the mapping with the bit set in the left
func split(mapping []int, i uint16) (left, right []int) {
	left, right = make([]int, 0, len(mapping)/2), make([]int, 0, len(mapping)/2)
	for j, v := range mapping {
		if i&(1<<uint(j)) > 0 {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return
}

func countBits(n uint16) (num int) {
	for i := uint(0); i < 16; i++ {
		num += int((n & (1 << i)) >> i)
	}
	return
}

func main() {
	fmt.Println(counts(12))
}
