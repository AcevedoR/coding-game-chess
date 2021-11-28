package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "starting game")
	var moveHistory = make([]Move, 100)

	var constantsCount int
	fmt.Scan(&constantsCount)

	for i := 0; i < constantsCount; i++ {
		var name, value string
		fmt.Scan(&name, &value)
		fmt.Fprintln(os.Stderr, "constant: ", name, " ", value)
	}

	fmt.Println("fen moves")

	for {
		fmt.Fprintln(os.Stderr, "starting turn")

		var boardInput, color, castling, enPassant string
		var halfMoveClock, fullMove int
		fmt.Scanln(&boardInput, &color, &castling, &enPassant, &halfMoveClock, &fullMove)
		fmt.Fprintln(os.Stderr, "fen: ", boardInput, color, castling, enPassant, halfMoveClock, fullMove)

		var availableMovesCount int
		fmt.Scanln(&availableMovesCount)
		var availableMoves = make([]string, availableMovesCount)

		for i := 0; i < availableMovesCount; i++ {
			var move string
			fmt.Scanln(&move)
			availableMoves[i] = move
		}
		fmt.Fprintln(os.Stderr, "moves: ", availableMoves)

		play(boardInput, color, availableMoves, &moveHistory)
	}
}

func play(boardInput string, color string, moves []string, moveHistory *[]Move) {
	legalMoves := parseMoves(moves)
	move := getBestMove(parseBoardInput(boardInput), color, legalMoves)
	fmt.Println(move.format()) // Write action to stdout
	*moveHistory = append(*moveHistory, move)
}

func getBestMove(board Board, color string, legalMoves []Move) Move {
    
	return legalMoves[len(legalMoves)-1]
}

func parseMoves(strMoves []string) []Move {
	var parsedMoves []Move
	for i := 0; i < len(strMoves); i++ {
		m := strMoves[i]
		parsedMoves = append(
			parsedMoves,
			Move{
				begin: Position{
					column: m[0],
					line:   int(m[1]),
				},
				end: Position{
					column: m[2],
					line:   int(m[3]),
				},
			},
		)
	}
	return parsedMoves
}

func parseBoardInput(boardInput string) Board {
	var board [8][8]byte
	var y = 7
	var x = 0
	for i := 0; i < len(boardInput); i++ {
		input := boardInput[i]
		fmt.Println(i, ", ", string(input)," (",int(input), "), ", x, ":", y)

		if input == '/' {
			y--
			x = 0
		} else if input >= 49 && input <= 56 {
            fmt.Println( "---xxx: ", x)
			x = int(input - 48)
		} else {
			board[x][y] = byte(input)
			x++
		}
	}

	for i := 0; i < len(board); i++ {
        fmt.Fprintln(os.Stderr, board[i])
    }

	return Board{grid: board}
}
type Board struct {
    grid [8][8]byte
}

type Position struct {
	column byte
	line   int
}

func (p Position) format() string {
	return string(p.column) + string(p.line)
}

type Move struct {
	begin Position
	end   Position
}

func (m Move) format() string {
	return m.begin.format() + m.end.format()
}

