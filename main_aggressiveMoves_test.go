package gochess

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPawnTake(t *testing.T) {
	// given
	board := parseBoardInput("8/8/1ppppp2/1ppppp2/1ppPpp2/1ppppp2/1ppppp2/8")

	// when
	result := getAllAggressiveMoves(board, "w")

	// then
	assert.Contains(
		t, 
		result, 
		Move{begin: *PositionOf('d', '4'), end: *PositionOf('c', '5')}, 
		Move{begin: *PositionOf('d', '4'), end: *PositionOf('e', '5')},
	)
}
