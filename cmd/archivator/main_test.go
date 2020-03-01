package main

import (
	"testing"
)

func Benchmark_sequenceArch(b *testing.B) {
	for i:=0; i< b.N; i++{
		sequenceArch([]string{
			"first.txt",
			"second.txt",
			"third.txt",
			"first.txt",
			"second.txt",
			"third.txt",
			"first.txt",
			"second.txt",
			"third.txt",
			"first.txt",
			"second.txt",
			"third.txt",
			"first.txt",
			"second.txt",
			"third.txt",
			"first.txt",
			"second.txt",
			"third.txt",
			"first.txt",
			"second.txt",
			"third.txt",

		})
	}
}

func Benchmark_conArch(b *testing.B) {
	for i:=0; i< b.N; i++{
		conArch([]string{
			"first.txt",
			"second.txt",
			"third.txt",
			"first.txt",
			"second.txt",
			"third.txt",
			"first.txt",
			"second.txt",
			"third.txt",
			"first.txt",
			"second.txt",
			"third.txt",
			"first.txt",
			"second.txt",
			"third.txt",
			"first.txt",
			"second.txt",
			"third.txt",
			"first.txt",
			"second.txt",
			"third.txt",
		})
	}
}