package main

import (
	"cmp"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"slices"
)

func genSample(sample []int) {
	for i := range sample {
		sample[i] = rand.Intn(2)
	}
}

func count(sample, pattern []int) int {
	score := 0
	for i := 0; i < len(sample)-len(pattern)+1; i++ {
		extracted := sample[i : i+len(pattern)]
		if slices.Equal(extracted, pattern) {
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
	hhPattern := []int{0, 0}
	htPattern := []int{0, 1}
	n := 1024
	hh := make([]int, n)
	ht := make([]int, n)
	diff := make(map[int]int)
	sample := make([]int, n)
	numSamples := n * 1024
	for i := 0; i < numSamples; i++ {
		genSample(sample)
		hhc := count(sample, hhPattern)
		htc := count(sample, htPattern)
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
