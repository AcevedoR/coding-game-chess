package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
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

	moves := make([]Move, 0)
	for i := 0; i < len(colorPieces); i++ {
		p := colorPieces[i]
		if p.Value == 'P' || p.Value == 'p' {
			moves = append(moves, getPawnMoves(board.Grid, isWhite, vmod, p.Position)...)
		} else if p.Value == 'R' || p.Value == 'r' {
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, 1, 0)...)
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, 0, 1)...)
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, -1, 0)...)
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, -0, -1)...)
		}else if p.Value == 'B' || p.Value == 'b' {
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, 1, 1)...)
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, -1, -1)...)
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, 1, -1)...)
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, -1, 1)...)
		}else if p.Value == 'Q' || p.Value == 'q' {
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, 1, 0)...)
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, 0, 1)...)
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, -1, 0)...)
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, -0, -1)...)
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, 1, 1)...)
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, -1, -1)...)
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, 1, -1)...)
			moves = append(moves, getAvailableMoves(board, p.Position, isWhite, -1, 1)...)
		}
	}
	return moves
}
func getAvailableMoves(board Board, origin Position, isWhite bool, horizontalDirection int, verticalDirection int) []Move {
	var auditBoard [8][8]byte
	var moves []Move = []Move{}
	horizontalSign := horizontalDirection
	verticalSign := verticalDirection
	if horizontalDirection == 0 {
		horizontalSign = 1
	}
	if verticalDirection == 0 {
		verticalSign = 1
	}

	for x := max(0, min(7, (origin.x)+horizontalDirection)); x >= 0 && x <= 7; x += horizontalSign {
		for y := max(0, min(7, (origin.y)+verticalDirection)); y >= 0 && y <= 7; y += verticalSign {
			auditBoard[x][y] = '-'

			if x == origin.x && y == origin.y {
				printBoard(auditBoard)
				return moves
			}
			target := board.Grid[x][y]
			if target == 0 {
				auditBoard[x][y] = 'x'
				moves = append(moves, Move{Begin: Position{x: origin.x, y: origin.y}, End: Position{x: x, y: y}})
			} else if determineIfWhite(target) != isWhite {
				auditBoard[x][y] = 'o'
				moves = append(moves, Move{Begin: Position{x: origin.x, y: origin.y}, End: Position{x: x, y: y}})
				printBoard(auditBoard)
				return moves
			} else {
				return moves
			}
		}
	}
	printBoard(auditBoard)

	return moves
}
func getPawnMoves(grid [8][8]byte, isWhite bool, vmod int, origin Position) []Move {
	var moves []Move = []Move{}
	if origin.y+vmod < 1 || origin.y+vmod > 6 {
		return moves
	}
	if grid[origin.x][origin.y+vmod] == 0 {
		moves = append(moves, moveOf(origin.x, origin.y, origin.x, origin.y+vmod))
	}
	if origin.x < 7 && grid[origin.x+1][origin.y+vmod] != 0 && determineIfWhite(grid[origin.x+1][origin.y+vmod]) != isWhite {
		moves = append(moves, moveWithTakeOf(origin.x, origin.y, origin.x+1, origin.y+vmod))
	}
	if origin.x > 0 && grid[origin.x-1][origin.y+vmod] != 0 && determineIfWhite(grid[origin.x-1][origin.y+vmod]) != isWhite {
		moves = append(moves, moveWithTakeOf(origin.x, origin.y, origin.x-1, origin.y+vmod))
	}
	return moves
}
func appendMoveIfPresent(moves *[]Move, move Move) {
	if (move != Move{}) {
		*moves = append(*moves, move)
	}
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
	printBoard(board)
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

func moveOf(ox int, oy int, tx int, ty int) Move {
	return Move{Begin: Position{x: ox, y: oy}, End: Position{x: tx, y: ty}}
}
func moveWithTakeOf(ox int, oy int, tx int, ty int) Move {
	return Move{Begin: Position{x: ox, y: oy}, End: Position{x: tx, y: ty}}
}
func (m Move) Format() string {
	return m.Begin.Format() + m.End.Format()
}
func printBoard(b [8][8]byte) {
	if isDebug() {
		data := [][]string{}
		println("FSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSSS")

		for y := 7; y >= 0; y-- {
			var line []string
			line = append(line, fmt.Sprint(int(y+1)))
			for x := 0; x <= 7; x++ {
				line = append(line, string(b[x][y]))
			}
			data = append(data, line)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetFooter([]string{" ", "a", "b", "c", "d", "e", "f", "g", "h"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render() // Send output
	}
}
func isDebug() bool {
	return true
}
