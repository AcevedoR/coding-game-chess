package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func xTestPawn(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/8/q1r5/3P4/P7/8/8")
	var moves = []string{"a2a3", "a2a4", "b2b3", "b2b4", "c2c3", "c2c4", "d1c3", "d1e3", "d2d3", "d2d4", "e1d3", "e1f3", "e2e3", "e2e4", "f2f3", "f2f4", "g2g3", "g2g4", "h2h3", "h2h4"}

	// when
	result := GetBestMove(board, true, ParseMoves(moves))

	// then
	assert.Equal(t, "d4c5", result.Format())
}

func xTestCheckMate(t *testing.T) {
	// given
	board := ParseBoardInput("nqbrkbnr/pppppppp/8/k7/1P6/8/PPPP4/NQBRK3")
	var moves = []string{"a2a3", "a2a4", "b2b3", "b2b4", "c2c3", "c2c4", "d1c3", "d1e3", "d2d3", "d2d4", "e1d3", "e1f3", "e2e3", "e2e4", "f2f3", "f2f4", "g2g3", "g2g4", "h2h3", "h2h4"}

	// when
	result := GetBestMove(board, true, ParseMoves(moves))

	// then
	assert.Equal(t, "b4a5", result.Format())
}
