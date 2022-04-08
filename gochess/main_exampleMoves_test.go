package main

import (
	"testing"
)

func TestIllegalMoves(t *testing.T) {
	// given
	tests := map[string]struct {
		inputBoard      string
		notExpectedMove string
		isBlack bool
	}{
		"Illegal move 1": {
			inputBoard: "bbnnrkrq/ppp1pppp/8/3p4/8/P7/1PPPPPPP/BBNNRKRQ",
			notExpectedMove: "a1a2",
		},
		"Illegal move 2": {
			inputBoard: "rnqknbbr/pppppppp/8/8/8/8/PPPPPPPP/RNQKNBBR",
			notExpectedMove: "a1a2",
		},
		"Illegal move 3": {
			inputBoard: "rnqknbbr/pppppppp/8/8/8/8/PPPPPPPP/RNQKNBBR",
			notExpectedMove: "a1a2",
		},
		"Illegal move 4": {
			inputBoard: "rbknr1bq/p1pppppp/6n1/1p6/8/P7/BPPPPPPP/R1KNRNBQ",
			notExpectedMove: "a2b5",
		},
		"Illegal move 5": {
			inputBoard: "2b2n1q/5r2/2pr1p2/p1b2k1p/Pp2p1p1/1B4P1/1PP3PP/1RBKQ1RN",
			notExpectedMove: "e1e4",
		},
		"Illegal move 6": {
			inputBoard: "Bk6/6q1/5p2/p5p1/P5Pp/R1PN1P2/4P2R/1r1KN3",
			notExpectedMove: "e1e4",
		},
		"Illegal move 7": {
			inputBoard: "r1r5/k1p3Q1/p4N2/2PRp3/6N1/3P2P1/5P1P/q3KRB1",
			notExpectedMove: "d5e5",
		},
		"Illegal move 8": {
			inputBoard: "3br1bB/r1knpp1p/1ppp4/5Pn1/8/P2P2N1/p1P1P2q/1K1BR2N",
			notExpectedMove: "a3a4",
		},
		// "Illegal move 9": {
		// 	// TODO set the queen as new piece in minmax
		// 	inputBoard: "1kbbnrqn/1ppppppp/8/8/1r6/3P2P1/PpP1PPNP/RK1B1RQN",
		// 	notExpectedMove: "b1b2",
		// 	isBlack: true,
		// },
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// when
			result := GetBestPlay(test.inputBoard, !test.isBlack)

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

	// when
	result := GetBestPlay(board, false)

	// then
	expected := "d6d1"
	assertEquals(t, result.Format(), expected)
}
