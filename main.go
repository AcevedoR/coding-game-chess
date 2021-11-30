package gochess

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
	aggressiveMoves := getAllAggressiveMoves(board, color)
	if len(aggressiveMoves) > 0 {
		return aggressiveMoves[0]
	}
	return legalMoves[len(legalMoves)-1]
}

func getAllAggressiveMoves(board Board, color string) []Move {
	vmod := 1
	if color != "w" {
		vmod = -1
	}
	aggressiveMoves := make([]Move, 0)
	for i := 0; i < len(board.pieces); i++ {
		p := board.pieces[i]
		if p.value == 'P' {
			frontRight := addPieceMovesIfValid(board, p.position.x, p.position.y, p.position.x + 1, p.position.y + vmod)
			if(frontRight != Move{}){
				aggressiveMoves = append(aggressiveMoves, frontRight)
			}
			frontLeft := addPieceMovesIfValid(board, p.position.x, p.position.y, p.position.x - 1, p.position.y + vmod)
			if(frontLeft != Move{}){
				aggressiveMoves = append(aggressiveMoves, frontLeft)
			}
		}
	}
	return aggressiveMoves
}

func addPieceMovesIfValid(board Board, originX int, originY, x int, y int) Move {
	if x >= 0 && x < 7 && y >= 0 && y < 7 {
		target := board.grid[x][y]
		if target != 0 {
			return Move{begin: Position{x:originX, y: originY}, end: Position{x:x, y: y}}
		}
	}
	return Move{}
}


func parseMoves(strMoves []string) []Move {
	var parsedMoves []Move
	for i := 0; i < len(strMoves); i++ {
		m := strMoves[i]
		parsedMoves = append(
			parsedMoves,
			Move{
				begin: *PositionOf(
					m[0],
					m[1],
				),
				end: *PositionOf(
					m[2],
					m[3],
				),
			},
		)
	}
	return parsedMoves
}

func parseBoardInput(boardInput string) Board {
	var board [8][8]byte
	var pieces [32]Piece
	var pieceIndex int
	var y = 7
	var x = 0
	for i := 0; i < len(boardInput); i++ {
		input := boardInput[i]
		// fmt.Println(i, ", ", string(input)," (",int(input), "), ", x, ":", y)

		if input == '/' {
			y--
			x = 0
		} else if input >= 49 && input <= 56 {
			x = int(input)-48+x
		} else {
			board[x][y] = byte(input)
			pieces[pieceIndex] = Piece{value: byte(input), position: Position{x: x, y: y}}
			pieceIndex++
			x++
		}
	}

	// fmt.Fprintln(os.Stderr, "---------")

	// for y := 7; y >= 0; y-- {
	// 	var str string
	// 	for x := 0; x <= 7; x++ {
	// 		str += string(board[x][y])
	// 	}
	// 	fmt.Fprintln(os.Stderr, str)
	// }
	// fmt.Fprintln(os.Stderr, "---------")

	return Board{grid: board, pieces: pieces}
}

type Board struct {
	grid   [8][8]byte
	pieces [32]Piece
}

func (b Board) get(column byte, line int) byte {
	return b.grid[column-97][line-1]
}

type Piece struct {
	value    byte
	position Position
}

type Position struct {
	x int
	y int
}

func PositionOf(column byte, line byte) *Position {
	return &Position{x: int(column) - 97, y: int(line) - 48 - 1}
}

func (p Position) format() string {
	return fmt.Sprintf("%s%d", string(byte(p.x+97)), p.y+1)
}

type Move struct {
	begin Position
	end   Position
}

func (m Move) format() string {
	return m.begin.format() + m.end.format()
}
