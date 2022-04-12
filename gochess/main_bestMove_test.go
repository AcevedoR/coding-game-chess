package main

import (
	"testing"
)

func TestPawn(t *testing.T) {
	// given
	board := "8/8/8/q1r5/3P4/P7/8/8"

	// when
	move := GetBestPlay(board, true)

	// then
	result := move.Format()
	expected := "d4c5"
	assertEquals(t, result, expected)
}

func TestPawnPromotionInTheFuture(t *testing.T){
	board := "8/7P/8/8/8/8/8/8"

	result := GetBestPlay(board, true)

	expected := moveOff('h', '7', 'h', '8')
	expected.PromotionPiece = 'q'
	if result.Format() != expected.Format() {
		t.Errorf("expected: %s, but got %s", expected.Format(), result.Format())
	}
}

func TestBlackPawnPromotionInTheFuture(t *testing.T){
	board := "8/8/8/8/8/8/p7/8"

	result := GetBestPlay(board, false)

	expected := moveOff('a', '2', 'a', '1')
	expected.PromotionPiece = 'q'
	if result.Format() != expected.Format() {
		t.Errorf("expected: %s, but got %s", expected.Format(), result.Format())
	}
}

func xTestCheckMate(t *testing.T) {
	// given
	board := "nqbrkbnr/pppppppp/8/k7/1P6/8/PPPP4/NQBRK3"

	// when
	result := GetBestPlay(board, true)

	// then
	expected := "b4a5"
	assertEquals(t, result.Format(), expected)
}

func TestAllyCheckMate(t *testing.T) {
	board := "k7/2q5/5b2/4B3/3K4/8/8/8"

	result := GetBestPlay(board, true)

	assertEquals(t, result.Format(), "e5f6")
}

func TestMinMax(t *testing.T) {
	board := ParseBoardInput("8/8/8/2r5/2n1b3/3B4/8/8")
	
	result := GetBestMoveMinMax(board, true, 3)
	
	expectedScore := -50
	if result.Score != expectedScore {
		t.Errorf("got %d instead of %d", result, expectedScore)
	}
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
