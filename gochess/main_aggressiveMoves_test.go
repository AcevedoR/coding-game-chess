package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPawnTake(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/1ppppp2/1ppppp2/1ppPpp2/1ppppp2/1ppppp2/8")

	// when
	result := GetAllAggressiveMoves(board, true)

	// then
	assert.Contains(
		t,
		result,
		moveOf('d', '4', 'c', '5'),
		moveOf('d', '4', 'e', '5'),
	)
}
func TestPawnTakeB(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/1PPPPP2/1PPPPP2/1PPpPP2/1PPPPP2/1PPPPP2/8")

	// when
	result := GetAllAggressiveMoves(board, false)

	// then
	assert.Contains(
		t,
		result,
		moveOf('d', '4', 'c', '3'),
	)
	assert.Contains(
		t,
		result,
		moveOf('d', '4', 'e', '3'),
	)
}

func TestPawnMove(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/8/8/3P4/8/8/8")

	// when
	result := GetAllAggressiveMoves(board, true)

	// then
	assert.Contains(
		t,
		result,
		moveOf('d', '4', 'd', '5'),
	)
}

func TestPawnMoveB(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/8/3p4/8/8/8/8")

	// when
	result := GetAllAggressiveMoves(board, false)

	// then
	assert.Contains(
		t,
		result,
		moveOf('d', '5', 'd', '4'),
	)
}

func TestPawnEdges(t *testing.T) {
	board := ParseBoardInput("P6P/8/8/8/8/8/8/p6p")

	result := GetAllAggressiveMoves(board, true)

	assertEmpty(t, result)
}

func TestPawnEdgesB(t *testing.T) {
	board := ParseBoardInput("P6P/8/8/8/8/8/8/p6p")

	result := GetAllAggressiveMoves(board, false)

	assertEmpty(t, result)
}

func xTestPawnShouldNotTakeFriendly(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/1PPPPP2/1PPPPP2/1PPPPP2/1PPPPP2/1PPPPP2/8")

	// when
	result := GetAllAggressiveMoves(board, true)

	// then
	assertEmpty(t, result)
}
func TestDoNotTryToPlayEnemy(t *testing.T) {
	// given
	// failure move: c6d7
	board := ParseBoardInput("rbkrq1bn/1p1ppppp/2p3n1/p7/2P5/6NP/PP1PPPP1/RBKRQNB1")

	// when
	result := GetAllAggressiveMoves(board, true)

	// then
	assert.NotContains(
		t,
		result,
		moveOf('c', '6', 'd', '7'),
	)
}
func xTestRockTakeClose(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/1ppppp2/1ppppp2/1ppRpp2/1ppppp2/1ppppp2/8")

	// when
	result := GetAllAggressiveMoves(board, true)

	// then
	assert.Contains(
		t,
		result,
		moveOf('d', '4', 'd', '5'),
	)
	assert.Contains(
		t,
		result,
		moveOf('d', '4', 'd', '3'),
	)
	assert.Contains(
		t,
		result,
		moveOf('d', '4', 'e', '4'),
	)
	assert.Contains(
		t,
		result,
		moveOf('d', '4', 'c', '4'),
	)

}

func xTestRockTakeFar(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/1ppppp2/1ppppp2/1ppRpp2/1ppppp2/1ppppp2/8")

	// when
	result := GetAllAggressiveMoves(board, true)

	// then
	assert.Contains(
		t,
		result,
		moveOf('d', '4', 'd', '8'),
	)
	assert.Contains(
		t,
		result,
		moveOf('d', '4', 'h', '4'),
	)
	assert.Contains(
		t,
		result,
		moveOf('d', '4', 'd', '2'),
	)
	assert.Contains(
		t,
		result,
		moveOf('d', '4', 'b', '4'),
	)
}

func moveOf(beginX byte, beginY byte, endX byte, endY byte) Move {
	return Move{Begin: *PositionOf(beginX, beginY), End: *PositionOf(endX, endY)}
}
func assertEmpty(t *testing.T, result []Move) {
	isResultEmpty := true

	for _, v := range result {
		if (v != Move{}) {
			isResultEmpty = false
		}
	}

	assert.True(
		t,
		isResultEmpty,
	)
}
