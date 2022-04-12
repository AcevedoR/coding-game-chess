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
	assertContainsOnly(
		t,
		result,
		moveOff('d', '4', 'c', '5'),
		moveOff('d', '4', 'e', '5'),
	)
}
func TestPawnTakeB(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/1PPPPP2/1PPPPP2/1PPpPP2/1PPPPP2/1PPPPP2/8")

	// when
	result := GetAllAggressiveMoves(board, false)

	// then
	assertContainsOnly(
		t,
		result,
		moveOff('d', '4', 'c', '3'),
		moveOff('d', '4', 'e', '3'),
	)
}

func TestPawnMove(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/8/8/3P4/8/8/8")

	// when
	result := GetAllAggressiveMoves(board, true)

	// then
	assertContainsOnly(
		t,
		result,
		moveOff('d', '4', 'd', '5'),
	)
}

func TestPawnMoveB(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/8/3p4/8/8/8/8")

	// when
	result := GetAllAggressiveMoves(board, false)

	// then
	assertContainsOnly(
		t,
		result,
		moveOff('d', '5', 'd', '4'),
	)
}

func TestPawnEdges(t *testing.T) {
	board := ParseBoardInput("8/p6p/8/8/8/8/P6P/8")

	result := GetAllAggressiveMoves(board, true)

	assertContainsOnly(
		t,
		result,
		moveOff('a', '2', 'a', '3'),
		moveOff('h', '2', 'h', '3'),
	)
}

func TestPawnEdgesB(t *testing.T) {
	board := ParseBoardInput("8/p6p/8/8/8/8/P6P/8")

	result := GetAllAggressiveMoves(board, false)

	assertContainsOnly(
		t,
		result,
		moveOff('a', '7', 'a', '6'),
		moveOff('h', '7', 'h', '6'),
	)
}

func TestPawnShouldNotTakeFriendly(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/1PPPPP2/1PPPPP2/1PPPPP2/1PPPPP2/1PPPPP2/8")

	// when
	result := GetAllAggressiveMoves(board, true)

	// then
	assert.NotContains(
		t,
		result,
		moveOff('d', '4', 'e', '1'),
		moveOff('d', '4', 'e', '2'),
		moveOff('d', '4', 'e', '3'),
		moveOff('d', '4', 'e', '4'),
		moveOff('d', '4', 'e', '5'),
		moveOff('d', '4', 'c', '1'),
		moveOff('d', '4', 'c', '2'),
		moveOff('d', '4', 'c', '3'),
		moveOff('d', '4', 'c', '4'),
		moveOff('d', '4', 'c', '5'),
		moveOff('d', '4', 'd', '5'),
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
		moveOff('c', '6', 'd', '7'),
	)
}
func TestRockTakeClose(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/1ppppp2/1ppppp2/1ppRpp2/1ppppp2/1ppppp2/8")

	// when
	result := GetAllAggressiveMoves(board, true)

	// then
	assertContains(
		t,
		result,
		moveOff('d', '4', 'd', '5'),
		moveOff('d', '4', 'd', '3'),
		moveOff('d', '4', 'e', '4'),
		moveOff('d', '4', 'c', '4'),
	)

}

func TestRockTakeFar(t *testing.T) {
	// given
	board := ParseBoardInput("3p4/8/8/8/1p1R3p/8/3p4/8")

	// when
	result := GetAllAggressiveMoves(board, true)

	// then
	assertContains(
		t,
		result,
		moveOff('d', '4', 'd', '8'),
		moveOff('d', '4', 'h', '4'),
		moveOff('d', '4', 'd', '2'),
		moveOff('d', '4', 'b', '4'),
	)
}
func TestBishopTakeClose(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/2PPP3/2PbP3/2PPP3/8/8/8")

	// when
	result := GetAllAggressiveMoves(board, false)

	// then
	assertContains(
		t,
		result,
		moveOff('d', '5', 'c', '4'),
		moveOff('d', '5', 'c', '6'),
		moveOff('d', '5', 'e', '4'),
		moveOff('d', '5', 'e', '6'),
	)
}
func TestKnightMove(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/8/8/8/8/1n6/8")

	// when
	result := GetAllAggressiveMoves(board, false)

	// then
	assertContains(
		t,
		result,
		moveOff('b', '2', 'a', '4'),
		moveOff('b', '2', 'c', '4'),
		moveOff('b', '2', 'd', '3'),
		moveOff('b', '2', 'd', '1'),
	)
}
func TestQueenTakeClose(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/2PPP3/2PqP3/2PPP3/8/8/8")

	// when
	result := GetAllAggressiveMoves(board, false)

	// then
	assertContains(
		t,
		result,
		moveOff('d', '5', 'c', '4'),
		moveOff('d', '5', 'c', '5'),
		moveOff('d', '5', 'c', '6'),
		moveOff('d', '5', 'e', '4'),
		moveOff('d', '5', 'e', '5'),
		moveOff('d', '5', 'e', '6'),
		moveOff('d', '5', 'd', '4'),
		moveOff('d', '5', 'd', '6'),
	)
}
func TestPawnPromotion(t *testing.T){
	board := ParseBoardInput("8/7P/8/8/8/8/p7/8")

	result := GetAllAggressiveMoves(board, true)

	expected := moveOff('h', '7', 'h', '8')
	expected.PromotionPiece = 'q'
	assertContains(
		t,
		result,
		expected,
	)
}

func xTestKingMove(t *testing.T) {
	// given
	board := ParseBoardInput("8/8/8/8/3K4/8/8/8")

	// when
	result := GetAllAggressiveMoves(board, false)

	// then
	assertContains(
		t,
		result,
		moveOff('d', '4', 'c', '4'),
		moveOff('d', '4', 'c', '5'),
		moveOff('d', '4', 'c', '6'),
		moveOff('d', '4', 'e', '4'),
		moveOff('d', '4', 'e', '5'),
		moveOff('d', '4', 'e', '6'),
		moveOff('d', '4', 'd', '4'),
		moveOff('d', '4', 'd', '6'),
	)
}

func moveOff(beginX byte, beginY byte, endX byte, endY byte) Move {
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
		result,
	)
}
func assertContains(t *testing.T, slice []Move, expectedSlice ...Move){
	var missing []Move = []Move{}
	for _, expected := range expectedSlice {
		found := false
		for _, move := range slice {
			if move.Begin == expected.Begin && move.End == expected.End && move.PromotionPiece == expected.PromotionPiece{
				found = true
			}
		}
		if !found {
			missing = append(missing, expected)
		}
	}
	if len(missing) > 0 {
		t.Errorf("\nInput: \n%+v \n\ndid not contain all of: \n%+v\n\n missing:\n %+v\n", slice, expectedSlice, missing)
	}
}
func assertContainsOnly(t *testing.T, slice []Move, expected ...Move){
	assertContains(t, slice , expected...)
	if len(slice) != len(expected){
		t.Errorf("\nExpected \n%+v \nwith size %d \n\ndid not contain ONLY all of:\n%+v\n of size %d\n", expected, slice,len(expected), len(slice))
	}
}