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

func TestAllyCheckMate(t *testing.T) {
	board := "k7/2q5/5b2/4B3/3K4/8/8/8"

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
