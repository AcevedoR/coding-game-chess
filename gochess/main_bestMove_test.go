package main

import (
	"testing"
)

func TestPawn(t *testing.T) {
	// given
	board := "8/8/8/q1r5/3P4/P7/8/8"
	var moves = []string{"a2a3", "a2a4", "b2b3", "b2b4", "c2c3", "c2c4", "d1c3", "d1e3", "d2d3", "d2d4", "e1d3", "e1f3", "e2e3", "e2e4", "f2f3", "f2f4", "g2g3", "g2g4", "h2h3", "h2h4"}

	// when
	move := GetBestPlay(moves, board, true)

	// then
	result := move.Format()
	expected := "d4c5"
	assertEquals(t, result, expected)
}

func xTestCheckMate(t *testing.T) {
	// given
	board := "nqbrkbnr/pppppppp/8/k7/1P6/8/PPPP4/NQBRK3"
	var moves = []string{"a2a3", "a2a4", "b2b3", "b2b4", "c2c3", "c2c4", "d1c3", "d1e3", "d2d3", "d2d4", "e1d3", "e1f3", "e2e3", "e2e4", "f2f3", "f2f4", "g2g3", "g2g4", "h2h3", "h2h4"}

	// when
	result := GetBestPlay(moves, board, true)

	// then
	expected := "b4a5"
	assertEquals(t, result.Format(), expected)
}

func TestIllegalMove1(t *testing.T) {
	// given
	board := "bbnnrkrq/ppp1pppp/8/3p4/8/P7/1PPPPPPP/BBNNRKRQ"
	var moves = []string{}

	// when
	result := GetBestPlay(moves, board, true)

	// then
	expected := "a1a2"
	assertNot(t, result.Format(), expected)
}
func TestIllegalMove2(t *testing.T) {
	// given
	board := "rnqknbbr/pppppppp/8/8/8/8/PPPPPPPP/RNQKNBBR"
	var moves = []string{}

	// when
	result := GetBestPlay(moves, board, true)

	// then
	expected := "a1a2"
	assertNot(t, result.Format(), expected)
}
func TestIllegalMove3(t *testing.T) {
	// given
	board := "8/8/8/8/8/8/8/B7"
	var moves = []string{}

	// when
	result := GetBestPlay(moves, board, true)

	// then
	expected := "a1a2"
	assertNot(t, result.Format(), expected)
}
func TestIllegalMove4(t *testing.T) {
	// given
	board := "rbknr1bq/p1pppppp/6n1/1p6/8/P7/BPPPPPPP/R1KNRNBQ"
	var moves = []string{}

	// when
	result := GetBestPlay(moves, board, true)

	// then
	expected := "a2b5"
	assertNot(t, result.Format(), expected)
}

func TestAllyCheckMate(t *testing.T) {
	board := "8/2q5/5b2/4B3/3K4/8/8/8"

	result := GetBestPlay(moves, board, true)

	assertEquals(t, result.Format(), "e5f6")
}

func assertEquals(t *testing.T, result string, expected string) {
	if result != expected {
		t.Errorf("got %s instead of %s", result, expected)
	}
}
func assertNot(t *testing.T, result string, notExpected string) {
	if result == notExpected {
		t.Errorf("got the only value we didn't wanted: %s", result)
	}
}
