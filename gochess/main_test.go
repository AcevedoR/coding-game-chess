package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var boardInput = "nqbrkbnr/pppppppp/8/8/8/8/PPPPPPPP/NQBRKBNR"
var isWhite = true
var moves = []string{"a2a3", "a2a4", "b2b3", "b2b4", "c2c3", "c2c4", "d1c3", "d1e3", "d2d3", "d2d4", "e1d3", "e1f3", "e2e3", "e2e4", "f2f3", "f2f4", "g2g3", "g2g4", "h2h3", "h2h4"}
var moveHistory = make([]Move, 0)

// in = "brknnqrb/pppppppp/8/8/8/8/PPPPPPPP/BRKNNQRB w BGbg - 0 1"

func TestNominalCase(t *testing.T) {
	// given

	// when
	Play(boardInput, isWhite, moves, &moveHistory)

	// then
	assert.Len(t, moveHistory, 1)
}
func TestParseMoves(t *testing.T) {
	// when
	result := ParseMoves(moves)

	// then
	assert.Equal(t, "a2a3", result[0].Format())
}
func TestParseBoardInput(t *testing.T) {
	// when
	result := ParseBoardInput(boardInput).Grid

	// then
	assert.Equal(t, "n", string(result[0][7]))
	assert.Equal(t, "R", string(result[7][0]))
	assert.Equal(t, "P", string(result[3][1]))
}
func TestParseBoardInputWithPlayedPieces(t *testing.T) {
	// given
	b := "rnknqrbb/ppppppp1/8/8/6pP/8/PPPPPP2/RNKNQRBB"
	// when
	result := ParseBoardInput(b).Grid
	// then
	assert.Equal(t, "r", string(result[0][7]))
	assert.Equal(t, "B", string(result[7][0]))
	assert.Equal(t, "P", string(result[3][1]))
	assert.Equal(t, "\x00", string(result[7][1]))
	assert.Equal(t, "p", string(result[6][3]))
	assert.Equal(t, "P", string(result[7][3]))
}

func TestBoardMove(t *testing.T) {
	board := ParseBoardInput("8/8/8/8/3P4/8/8/8")

	result := board.Move(moveOff('d','4','d','5'))

	piece := result.get('d', 5)
	expected := byte('P')
	if piece != expected {
		t.Errorf("got %b instead of %b", result, expected)
	}
}