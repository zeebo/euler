package main

//NOTE: for this problem, add a newline to the end of sets.txt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	Left = iota
	Right
	None
)

type Set struct {
	data []int32
	p    []int
}

//done when we reach none everywhere
func (s Set) Done() bool {
	for i := range s.p {
		if s.p[i] != None {
			return false
		}
	}
	return true
}

//add and carry!
func (s Set) Next() {
	var i int
	for i = 0; s.p[i] == None; i++ {
		s.p[i] = Left
	}
	s.p[i]++
}

func (s Set) Valid() bool {
	for !s.Done() && s.ValidPerm() {
		s.Next()
	}
	return s.Done()
}

func (s Set) ValidPerm() bool {
	//partials
	var left, right int32
	var lc, rc int

	for i, v := range s.p {
		switch v {
		case Left:
			lc++
			left += s.data[i]
		case Right:
			rc++
			right += s.data[i]
		}
	}

	switch {
	case left == right:
		return false
	case lc > rc && right > left:
		return false
	case rc > lc && left > right:
		return false
	}

	return true
}

func (s *Set) Load(r *bufio.Reader) (err error) {
	//reslice to length 0
	s.data = s.data[:0]

	line, err := r.ReadString('\n')
	if err != nil {
		return
	}

	chunks := strings.Split(line[:len(line)-2], ",")
	var nv int64
	for _, v := range chunks {
		nv, err = strconv.ParseInt(v, 10, 32)
		if err != nil {
			return
		}
		s.data = append(s.data, int32(nv))
	}

	//reset the partitons
	s.p = s.p[:len(s.data)]
	for i := range s.p {
		s.p[i] = 0
	}

	return
}

func (s Set) Sum() (total int32) {
	for _, v := range s.data {
		total += v
	}
	return
}

func (s Set) String() string {
	return fmt.Sprint(s.data)
}

const top = 1 << 12

func has(x int32, mask uint) bool {
	return x&(1<<mask) > 0
}

func main() {
	f, err := os.Open("sets.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := bufio.NewReader(f)

	s := Set{
		data: make([]int32, 0, 12),
		p:    make([]int, 0, 12),
	}

	var total int32
	for {
		err = s.Load(b)
		if err != nil {
			//expected
			if err == io.EOF {
				break
			}

			panic(err)
		}

		if s.Valid() {
			fmt.Printf("%v is valid.\n", s)
			total += s.Sum()
		} else {
			fmt.Printf("%v is invalid.\n", s)
		}
	}

	fmt.Println(total)
}
