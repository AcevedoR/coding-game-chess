package main

import (
	"testing"
	"time"
)

func BenchmarkOneRound(b *testing.B) {
	board := "8/2q5/5b2/4B3/3K4/8/8/8"

	GetBestPlay(moves, board, true)
}

func TestTimeout1(t *testing.T) {
	board := "5RN1/3k2r1/p3p3/6pB/2n4p/8/3r2PP/R3K1NQ"
	var moves = []string{}

	start := time.Now()
	GetBestPlay(moves, board, true)
	end := time.Now()

	result := end.Sub(start).Milliseconds() 
	expected := 20
	if result >= int64(expected) {
		t.Errorf("code took %d, but it should have been below %d", result, expected)
	}
}