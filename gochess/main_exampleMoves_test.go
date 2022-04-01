package main

import (
	"testing"
)

func TestIllegalMove1(t *testing.T) {
	// given
	board := "bbnnrkrq/ppp1pppp/8/3p4/8/P7/1PPPPPPP/BBNNRKRQ"
	var moves = []string{}

	// when
	result := GetBestPlay(moves, board, true)

	// then
	notExpected := "a1a2"
	assertNot(t, result.Format(), notExpected)
}
func TestIllegalMove2(t *testing.T) {
	// given
	board := "rnqknbbr/pppppppp/8/8/8/8/PPPPPPPP/RNQKNBBR"
	var moves = []string{}

	// when
	result := GetBestPlay(moves, board, true)

	// then
	notExpected := "a1a2"
	assertNot(t, result.Format(), notExpected)
}
func TestIllegalMove3(t *testing.T) {
	// given
	board := "8/8/8/8/8/8/8/B7"
	var moves = []string{}

	// when
	result := GetBestPlay(moves, board, true)

	// then
	notExpected := "a1a2"
	assertNot(t, result.Format(), notExpected)
}
func TestIllegalMove5(t *testing.T) {
	// given
	board := "rbknr1bq/p1pppppp/6n1/1p6/8/P7/BPPPPPPP/R1KNRNBQ"
	var moves = []string{}

	// when
	result := GetBestPlay(moves, board, true)

	// then
	notExpected := "a2b5"
	assertNot(t, result.Format(), notExpected)
}
func TestIllegalMove4(t *testing.T) {
	// given
	board := "2b2n1q/5r2/2pr1p2/p1b2k1p/Pp2p1p1/1B4P1/1PP3PP/1RBKQ1RN"
	var moves = []string{}

	// when
	result := GetBestPlay(moves, board, true)

	// then
	notExpected := "e1e4"
	assertNot(t, result.Format(), notExpected)
}
func TestIllegalMove4OppositeTakeKing(t *testing.T) {
	// given
	board := "2b2n1q/5r2/2pr1p2/p1b2k1p/Pp2p1p1/1B4P1/1PP3PP/1RBKQ1RN"
	var moves = []string{}

	// when
	result := GetBestPlay(moves, board, false)

	// then
	expected := "d6d1"
	assertEquals(t, result.Format(), expected)
}
func TestIllegalMove6(t *testing.T) {
	// given
	board := "Bk6/6q1/5p2/p5p1/P5Pp/R1PN1P2/4P2R/1r1KN3"
	var moves = []string{}

	// when
	result := GetBestPlay(moves, board, true)

	// then
	notExpected := "h2h4"
	assertNot(t, result.Format(), notExpected)
}