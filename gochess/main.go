package main

import (
	"fmt"
	"math/rand"
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

		isWhite := color == "w"
		Play(boardInput, isWhite, availableMoves, &moveHistory)
	}
}

func Play(boardInput string, isWhite bool, moves []string, moveHistory *[]Move) {
	legalMoves := ParseMoves(moves)
	move := GetBestMove(ParseBoardInput(boardInput), isWhite, legalMoves)
	fmt.Println(move.Format()) // Write action to stdout
	*moveHistory = append(*moveHistory, move)
}

func GetBestMove(board Board, isWhite bool, legalMoves []Move) Move {
	aggressiveMoves := GetAllAggressiveMoves(board, isWhite)
	if len(aggressiveMoves) > 0 {
		return aggressiveMoves[0]
	}
	if len(legalMoves) > 1 {
		return legalMoves[rand.Intn(len(legalMoves)-1)]
	} else {
		return legalMoves[0]
	}
}

func GetAllAggressiveMoves(board Board, isWhite bool) []Move {
	vmod := 1
	var colorPieces *[30]Piece = &board.WhitePieces
	if !isWhite {
		colorPieces = &board.BlackPieces
		vmod = -1
	}

	aggressiveMoves := make([]Move, 0)
	for i := 0; i < len(colorPieces); i++ {
		p := colorPieces[i]
		if p.Value == 'P' || p.Value == 'p' {
			frontRight := checkLineTakes(board, isWhite, p.Position, 1, vmod, true)
			if (frontRight != Move{}) {
				aggressiveMoves = append(aggressiveMoves, frontRight)
			}
			frontLeft := checkLineTakes(board, isWhite, p.Position, -1, vmod, true)
			if (frontLeft != Move{}) {
				aggressiveMoves = append(aggressiveMoves, frontLeft)
			}
		} else if p.Value == 'P' || p.Value == 'p' {
			// up := AddPieceMovesIfValid(board, isWhite, p.Position.x, p.Position.y, p.Position.x, p.Position.y + 1)

		}
	}
	return aggressiveMoves
}
func checkLineTakes(board Board, isWhite bool, origin Position, horizontalDirection int, verticalDirection int, close bool) Move {
	xGoal := 7
	yGoal := 7
	if close {
		xGoal = max(0, min(7, origin.x + horizontalDirection))
		yGoal = max(0, min(7, origin.y + verticalDirection))
	}
	for x := max(0, min(7, origin.x + horizontalDirection)); x <= xGoal; x++ {
		for y := max(0, min(7, origin.y + verticalDirection)); y <= yGoal; y++ {
			target := board.Grid[x][y]
			if target != 0 && determineIfWhite(target) != isWhite {
				return Move{Begin: Position{x: origin.x, y: origin.y}, End: Position{x: x, y: y}}
			}
		}
	}
	return Move{}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func AddPieceMovesIfValid(board Board, isWhite bool, originX int, originY, x int, y int) Move {
	if x >= 0 && x < 7 && y >= 0 && y < 7 {
		target := board.Grid[x][y]
		if target != 0 && determineIfWhite(target) != isWhite {
			return Move{Begin: Position{x: originX, y: originY}, End: Position{x: x, y: y}}
		}
	}
	return Move{}
}
func determineIfWhite(piece byte) bool {
	return piece >= 'A' && piece <= 'Z'
}

func ParseMoves(strMoves []string) []Move {
	var parsedMoves []Move
	for i := 0; i < len(strMoves); i++ {
		m := strMoves[i]
		parsedMoves = append(
			parsedMoves,
			Move{
				Begin: *PositionOf(
					m[0],
					m[1],
				),
				End: *PositionOf(
					m[2],
					m[3],
				),
			},
		)
	}
	return parsedMoves
}

func ParseBoardInput(boardInput string) Board {
	var board [8][8]byte
	var whitePieces [30]Piece
	var blackPieces [30]Piece
	var whitePiecesIndex, blackPiecesIndex int
	var y = 7
	var x = 0
	for i := 0; i < len(boardInput); i++ {
		input := boardInput[i]
		// fmt.Println(i, ", ", string(input)," (",int(input), "), ", x, ":", y)

		if input == '/' {
			y--
			x = 0
		} else if input >= 49 && input <= 56 {
			x = int(input) - 48 + x
		} else {
			board[x][y] = byte(input)
			if determineIfWhite(byte(input)) {
				whitePieces[whitePiecesIndex] = Piece{Value: byte(input), Position: Position{x: x, y: y}}
				whitePiecesIndex++
			} else {
				blackPieces[blackPiecesIndex] = Piece{Value: byte(input), Position: Position{x: x, y: y}}
				blackPiecesIndex++
			}
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

	return Board{Grid: board, WhitePieces: whitePieces, BlackPieces: blackPieces}
}

type Board struct {
	Grid        [8][8]byte
	WhitePieces [30]Piece
	BlackPieces [30]Piece
}

func (b Board) get(column byte, line int) byte {
	return b.Grid[column-97][line-1]
}

type Piece struct {
	Value    byte
	Position Position
}

type Position struct {
	x int
	y int
}

func PositionOf(column byte, line byte) *Position {
	return &Position{x: int(column) - 97, y: int(line) - 48 - 1}
}

func (p Position) Format() string {
	return fmt.Sprintf("%s%d", string(byte(p.x+97)), p.y+1)
}

type Move struct {
	Begin Position
	End   Position
}

func (m Move) Format() string {
	return m.Begin.Format() + m.End.Format()
}
