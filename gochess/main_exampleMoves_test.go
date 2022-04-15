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
		"Illegal move 9": {
		 	inputBoard: "1kbbnrqn/1ppppppp/8/8/1r6/3P2P1/PpP1PPNP/RK1B1RQN",
		 	notExpectedMove: "b1b2",
		 	isBlack: true,
		 },
	 	"Illegal move 10": {
		 	inputBoard: "1Q6/3kp3/p7/2pB1pP1/8/1P2P3/Pn1P1P2/BRNKRN2",
		 	notExpectedMove: "g5g6",
		 },
		 "Illegal move 10 bis": {
		 	inputBoard: "8/3k1r2/8/3B4/8/8/1n6/1R1K4",
		 	notExpectedMove: "d5f7",
		 },
		 "Illegal move 12": {
		 	inputBoard: "b2n1rr1/pp1pbpkp/2p1pqp1/8/7P/5PP1/PPPnP3/BBNNRKR1",
		 	notExpectedMove: "c2c3",
		 }, 
		 "Queen under attack should move": {
		 	inputBoard: "b2n1rr1/pp1pbpkp/2pnpqp1/8/4Q2P/6P1/PPPPPPR1/BBNNRK2",
		 	notExpectedMove: "g2g1",
		 },
		//  "Should take easy pawn and prevent nearby checkmate": {
		//  	inputBoard: "r2k2r1/p3p1P1/2pbq2p/1p1p4/3n2P1/2P5/PP1PPP2/R1BK1NRN",
		//  	notExpectedMove: "g1g2",
		//  },
		//  "Queen should prevent incoming checkmate": {
		//  	inputBoard: "rnknb1r1/ppppppbp/1q4p1/3N4/8/2N5/PPPPPPPP/R1K1BBRQ",
		//  	notExpectedMove: "b6d4",
		// 	 isBlack: true,
		//  },
		//  "Queen should prevent incoming checkmate 2": {
		//  	inputBoard: "3r4/1pk2p2/p1n4p/8/8/2QP4/PPP1Pb2/RN1KBq2",
		//  	notExpectedMove: "c3f6",
		//  },
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// when
			result := GetBestPlay(test.inputBoard, !test.isBlack, false, &[]Move{})

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
	result := GetBestPlay(board, false, false, &[]Move{})

	// then
	expected := "d6d1"
	assertEquals(t, result.Format(), expected)
}
