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
		Move{Begin: *PositionOf('d', '4'), End: *PositionOf('c', '5')},
		Move{Begin: *PositionOf('d', '4'), End: *PositionOf('e', '5')},
	)
}

func TestPawnTakeWithBlacks(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/1PPPPP2/1PPPPP2/1PPpPP2/1PPPPP2/1PPPPP2/8")

	// when
	result := GetAllAggressiveMoves(board, false)

	// then
	assert.Contains(
		t,
		result,
		Move{Begin: *PositionOf('d', '4'), End: *PositionOf('c', '3')},
		Move{Begin: *PositionOf('d', '4'), End: *PositionOf('e', '3')},
	)
}

func TestPawnShouldNotTakeFriendly(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/1PPPPP2/1PPPPP2/1PPPPP2/1PPPPP2/1PPPPP2/8")

	// when
	result := GetAllAggressiveMoves(board, true)

	// then
	isResultEmpty := true
	for _, v := range result {
		if (v != Move{}){
			isResultEmpty = false
		}
	}

	assert.True(
		t,
		isResultEmpty,
	)
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
		Move{Begin: *PositionOf('c', '6'), End: *PositionOf('d', '7')},
	)
}
func xTestRockTake(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/1ppppp2/1ppppp2/1ppRpp2/1ppppp2/1ppppp2/8")

	// when
	result := GetAllAggressiveMoves(board, true)

	// then
	assert.Contains(
		t,
		result,
		Move{Begin: *PositionOf('d', '4'), End: *PositionOf('d', '5')},
		Move{Begin: *PositionOf('d', '4'), End: *PositionOf('d', '3')},
		Move{Begin: *PositionOf('d', '4'), End: *PositionOf('e', '4')},
		Move{Begin: *PositionOf('d', '4'), End: *PositionOf('c', '4')},
	)
}