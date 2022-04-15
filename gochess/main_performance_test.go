package main

import (
	"testing"
	"time"
)

var history []Move= []Move{}

func BenchmarkOneRound(b *testing.B) {
	board := "8/2q5/5b2/4B3/3K4/8/8/8"

	GetBestPlay(board, true, false, &history)
}

func TestTimeout1(t *testing.T) {
	board := "5RN1/3k2r1/p3p3/6pB/2n4p/8/3r2PP/R3K1NQ"

	start := time.Now()
	GetBestPlay(board, true, false, &history)
	end := time.Now()

	result := end.Sub(start).Milliseconds() 
	expected := 40
	if result >= int64(expected) {
		t.Errorf("code took %d, but it should have been below %d", result, expected)
	}
}
func TestTimeout2(t *testing.T) {
	board := "b2k1n2/p2prp2/1q5b/4P2P/P7/1Pr2NN1/K7/1Q3B2"

	start := time.Now()
	GetBestPlay(board, true, false, &history)
	end := time.Now()

	result := end.Sub(start).Milliseconds() 
	expected := 40
	if result >= int64(expected) {
		t.Errorf("code took %d, but it should have been below %d", result, expected)
	}
}
func TestTimeoutRound1(t *testing.T) {
	board := "nbrkbnqr/pppppppp/8/8/8/8/PPPPPPPP/NBKRBNQR"

	start := time.Now()
	GetBestPlay(board, true, true, &history)
	end := time.Now()

	result := end.Sub(start).Milliseconds() 
	expected := 900
	if result >= int64(expected) {
		t.Errorf("code took %d, but it should have been below %d", result, expected)
	}
}