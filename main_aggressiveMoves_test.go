package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPawnTake(t *testing.T) {
	// given
	board := parseBoardInput("8/8/1ppppp2/1ppppp2/1ppPpp2/1ppppp2/1ppppp2/8")

	// when
	result := getAllAggressiveMoves(board, true)

	// then
	assert.Contains(
		t,
		result,
		Move{begin: *PositionOf('d', '4'), end: *PositionOf('c', '5')},
		Move{begin: *PositionOf('d', '4'), end: *PositionOf('e', '5')},
	)
}

func TestPawnTakeWithBlacks(t *testing.T) {
	// given
	board := parseBoardInput("8/8/1PPPPP2/1PPPPP2/1PPpPP2/1PPPPP2/1PPPPP2/8")

	// when
	result := getAllAggressiveMoves(board, false)

	// then
	assert.Contains(
		t,
		result,
		Move{begin: *PositionOf('d', '4'), end: *PositionOf('c', '3')},
		Move{begin: *PositionOf('d', '4'), end: *PositionOf('e', '3')},
	)
}

func TestPawnShouldNotTakeFriendly(t *testing.T) {
	// given
	board := parseBoardInput("8/8/1PPPPP2/1PPPPP2/1PPPPP2/1PPPPP2/1PPPPP2/8")

	// when
	result := getAllAggressiveMoves(board, true)

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
func xTestDoNotTryToPlayEnemy(t *testing.T) {
	// given
	// failure move: c6d7
	board := parseBoardInput("rbkrq1bn/1p1ppppp/2p3n1/p7/2P5/6NP/PP1PPPP1/RBKRQNB1")

	// when
	result := getAllAggressiveMoves(board, true)

	// then
	assert.Contains(
		t,
		result,
		Move{begin: *PositionOf('d', '4'), end: *PositionOf('c', '4')},
	)
}
func xTestRockTake(t *testing.T) {
	// given
	board := parseBoardInput("8/8/1ppppp2/1ppppp2/1ppRpp2/1ppppp2/1ppppp2/8")

	// when
	result := getAllAggressiveMoves(board, true)

	// then
	assert.Contains(
		t,
		result,
		Move{begin: *PositionOf('d', '4'), end: *PositionOf('d', '5')},
		Move{begin: *PositionOf('d', '4'), end: *PositionOf('d', '3')},
		Move{begin: *PositionOf('d', '4'), end: *PositionOf('e', '4')},
		Move{begin: *PositionOf('d', '4'), end: *PositionOf('c', '4')},
	)
}