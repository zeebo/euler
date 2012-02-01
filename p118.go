package main

import (
	"fmt"
	"math/big"
	"sort"
	"strings"
	"sync"
)

var buf [20]byte //ill never permute more than 20 things hehe

type Pool []int

func (p Pool) Reset() {
	for i := range p {
		p[i] = i
	}
}

func (p Pool) Next() {
	//degenerate case
	if len(p) == 1 {
		return
	}

	var i, j int
	for i = len(p) - 1; p[i-1] > p[i]; i-- {
		if i == 1 {
			p.Reset()
			return
		}
	}

	for j = len(p); p[j-1] < p[i-1]; j-- {
	}

	p[i-1], p[j-1] = p[j-1], p[i-1]
	for i, j = i+1, len(p); i < j; i, j = i+1, j-1 {
		p[i-1], p[j-1] = p[j-1], p[i-1]
	}
}

func (p Pool) NumCycles() (x int) {
	x = 1
	for i := len(p); i > 1; i-- {
		x *= i
	}
	return
}

func (p Pool) On(f []byte) string {
	defer p.Next()
	for i, v := range p {
		buf[i] = f[v]
	}
	return string(buf[:len(p)])
}

type SubSet int64

func (s *SubSet) Next() {
	*s = *s + 1
}

func (s *SubSet) On(f []byte) []byte {
	defer s.Next()
	var mask int64 = 1
	j := 0
	for i := range f {
		if int64(*s)&mask > 0 {
			buf[j] = f[i]
			j++
		}
		mask <<= 1
	}

	ret := make([]byte, j)
	copy(ret, buf[:j])
	return ret
}

func main() {
	primes := make(map[int][]string)
	x := new(big.Int)

	var y SubSet
	//ignore empty subset
	y.Next()

	data := []byte("123456789")
	for i := 0; i < 1<<9-1; i++ {
		subset := y.On(data)
		pool := make(Pool, len(subset))
		pool.Reset()

		//loop over every permutation
		for j := 0; j < pool.NumCycles(); j++ {
			x.SetString(pool.On(subset), 10)

			//2 throws out the same amount as 100.
			if big.ProbablyPrime(x, 2) {
				primes[len(subset)] = append(primes[len(subset)], x.String())
			}
		}
	}

	fmt.Println("Primes generated.")

	//30 partitions of 9 so just list them!
	partitions := [][]int{
		{9}, {8, 1}, {7, 2}, {7, 1, 1}, {6, 3}, {6, 2, 1}, {6, 1, 1, 1}, {5, 4},
		{5, 3, 1}, {5, 2, 2}, {5, 2, 1, 1}, {5, 1, 1, 1, 1}, {4, 4, 1}, {4, 3, 2},
		{4, 3, 1, 1}, {4, 2, 2, 1}, {4, 2, 1, 1, 1}, {4, 1, 1, 1, 1, 1}, {3, 3, 3},
		{3, 3, 2, 1}, {3, 3, 1, 1, 1}, {3, 2, 2, 2}, {3, 2, 2, 1, 1},
		{3, 2, 1, 1, 1, 1}, {3, 1, 1, 1, 1, 1, 1}, {2, 2, 2, 2, 1},
		{2, 2, 2, 1, 1, 1}, {2, 2, 1, 1, 1, 1, 1}, {2, 1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

	//check for typos
	if len(partitions) != 30 {
		panic("Too many/few")
	}

	for _, p := range partitions {
		var acc int
		for _, v := range p {
			acc += v
		}
		if acc != 9 {
			panic("One doesn't sum to 30")
		}
	}

	fmt.Println("Partitions generated.")

	group := new(sync.WaitGroup)
	group.Add(len(partitions))

	//for each partition, we have to find all the sets of primes that satisfy
	for _, p := range partitions {
		//now we have a partition! lets get the number of sets
		go numSets(primes, p, "", nil, group)
	}
	sy := make(chan bool)

	//send a message when the group is finished
	go func() {
		group.Wait()
		sy <- true
	}()

	items := make(map[string]struct{})

collecting:
	for {
		select {
		case set := <-sets:
			//turn the sets into a key and put it into our map
			sort.Strings(set)
			items[strings.Join(set, ",")] = struct{}{}
		case <-sy:
			break collecting
		}
	}

	fmt.Println(len(items))
}

var sets = make(chan []string)

func numSets(primes map[int][]string, p []int, curr string, n []string, g *sync.WaitGroup) {
	//simple base case
	if len(p) == 0 {
		sets <- n
		return
	}

	size := p[0]
	for _, prime := range primes[size] {
		if strings.IndexAny(curr, prime) == -1 {
			numSets(primes, p[1:], curr+string(prime), append(n, prime), nil)
		}
	}

	if g != nil {
		fmt.Println(p)
		g.Done()
	}
}
