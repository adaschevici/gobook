package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
)

const TESTLENGTH = 100000

type CPUHog struct {
	mHog []MemoryHog
}
type MemoryHog struct {
	a, b, c, d, e, f, g int64
	h, i, j, k, l, m, n float64
	longByte            []byte
}

func makeMemoryHog() []MemoryHog {

	memoryHogs := make([]MemoryHog, TESTLENGTH)

	for i := 0; i < TESTLENGTH; i++ {
		m := MemoryHog{}
		_ = append(memoryHogs, m)
	}

	return memoryHogs
}

var profile = flag.String("cpuprofile", "", "output pprof data to file")

func main() {
	var CPUHogs []CPUHog

	flag.Parse()
	if *profile != "" {
		flag, err := os.Create(*profile)
		if err != nil {
			fmt.Println("Could not create profile", err)
		}
		pprof.StartCPUProfile(flag)
		defer pprof.StopCPUProfile()

	}

	for i := 0; i < TESTLENGTH; i++ {
		fmt.Println("Appending to memory HOG", i)
		hog := CPUHog{}
		hog.mHog = makeMemoryHog()
		_ = append(CPUHogs, hog)
	}
}
