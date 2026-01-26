package common

import (
	"fmt"
	"log"
	"os"
	"time"
)

func GetFile() *os.File {
	filename := GetFilename()
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func GetFilename() string {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}
	return os.Args[1]
}

func TimeIt[T any](f func() T) (T, time.Duration) {
	start := time.Now()
	res := f()
	dur := time.Since(start)
	return res, dur
}

func Assert(cond bool, msg string) {
	if !cond {
		panic("assertion failed: " + msg)
	}
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
