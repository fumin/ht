package main

import (
	"cmp"
	"flag"
	"fmt"
	"log"
	"slices"
)

var (
	seqLen = flag.Int("n", 20, "sequence length")
)

type BitInt struct {
	I int
	N int
}

func count(roll, pattern BitInt) int {
	score := 0
	for i := 0; i < roll.N-pattern.N+1; i++ {
		mask := (1 << pattern.N) - 1
		shift := roll.N - pattern.N - i
		extracted := (roll.I & (mask << shift)) >> shift

		if extracted == pattern.I {
			score++
		}
	}
	return score
}

type scoreFrequency struct {
	score     int
	frequency int
}

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)
	if err := mainWithErr(); err != nil {
		log.Fatalf("%+v", err)
	}
}

func mainWithErr() error {
	hhPattern := BitInt{I: 0, N: 2}
	htPattern := BitInt{I: 1, N: 2}
	hh := make([]int, *seqLen)
	ht := make([]int, *seqLen)
	diff := make(map[int]int)
	for i := 0; i < (1 << *seqLen); i++ {
		roll := BitInt{I: i, N: *seqLen}
		hhc := count(roll, hhPattern)
		htc := count(roll, htPattern)
		hh[hhc]++
		ht[htc]++
		diff[hhc-htc] += 1
	}

	fmt.Printf("score,hh,ht\n")
	for score, hhc := range hh {
		htc := ht[score]
		fmt.Printf("%d,%d,%d\n", score, hhc, htc)
	}

	diffs := make([]scoreFrequency, 0, len(diff))
	for d, cnt := range diff {
		diffs = append(diffs, scoreFrequency{score: d, frequency: cnt})
	}
	slices.SortFunc(diffs, func(a, b scoreFrequency) int { return cmp.Compare(a.score, b.score) })
	fmt.Printf("score,hh-ht\n")
	for _, d := range diffs {
		fmt.Printf("%d,%d\n", d.score, d.frequency)
	}

	return nil
}
