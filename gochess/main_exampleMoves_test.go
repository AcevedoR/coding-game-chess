package main

import (
	"testing"
)

func TestIllegalMoves(t *testing.T) {
	// given
	tests := map[string]struct {
		inputBoard      string
		notExpectedMove string
	}{
		"Illegal move 1": {
			"bbnnrkrq/ppp1pppp/8/3p4/8/P7/1PPPPPPP/BBNNRKRQ",
			"a1a2",
		},
		"Illegal move 2": {
			"rnqknbbr/pppppppp/8/8/8/8/PPPPPPPP/RNQKNBBR",
			"a1a2",
		},
		"Illegal move 3": {
			"rnqknbbr/pppppppp/8/8/8/8/PPPPPPPP/RNQKNBBR",
			"a1a2",
		},
		"Illegal move 4": {
			"rbknr1bq/p1pppppp/6n1/1p6/8/P7/BPPPPPPP/R1KNRNBQ",
			"a2b5",
		},
		"Illegal move 5": {
			"2b2n1q/5r2/2pr1p2/p1b2k1p/Pp2p1p1/1B4P1/1PP3PP/1RBKQ1RN",
			"e1e4",
		},
		"Illegal move 6": {
			"Bk6/6q1/5p2/p5p1/P5Pp/R1PN1P2/4P2R/1r1KN3",
			"e1e4",
		},
		"Illegal move 7": {
			"r1r5/k1p3Q1/p4N2/2PRp3/6N1/3P2P1/5P1P/q3KRB1",
			"d5e5",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// given
			var moves = []string{}
			// when
			result := GetBestPlay(moves, test.inputBoard, true)

			// then
			if result.Format() == test.notExpectedMove {
				t.Errorf("got the only value we didn't wanted: %s", test.notExpectedMove)
			}
		})
	}
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
