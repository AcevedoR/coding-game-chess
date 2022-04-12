package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "starting game")
	var moveHistory []Move

	var constantsCount int
	fmt.Scan(&constantsCount)

	for i := 0; i < constantsCount; i++ {
		var name, value string
		fmt.Scan(&name, &value)
		fmt.Fprintln(os.Stderr, "constant: ", name, " ", value)
	}

	fmt.Println("fen moves")
	firstRound := true

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
		Play(boardInput, isWhite, availableMoves, &moveHistory, firstRound)
		firstRound = false
	}
}

func Play(boardInput string, isWhite bool, moves []string, moveHistory *[]Move, firstRound bool) {

	move := GetBestPlay(boardInput, isWhite, firstRound)
	fmt.Println(move.Format()) // Write action to stdout
	//*moveHistory = append(*moveHistory, move)
}

func GetBestPlay(boardInput string, isWhite bool, firstRound bool) Move {
	board := ParseBoardInput(boardInput)
	depth := 3
	if firstRound {
		depth = 4
	}
	move := GetBestMoveMinMax(board, isWhite, depth)
	if isDebug() {
		fmt.Println("moves plan:")
		for i:= 0; i < len(move.History); i++ { 
			fmt.Println(move.History[i].Format())
		}
	}
	return move.Move
}

func GetBestMoveMinMax(board Board, isWhite bool, depth int) MinMaxScore {
	if depth == 0 {
		return MinMaxScore{Move{}, board.GetPositionalScore(), []Move{}}
	}
	printBoard(board.Grid)
	if isWhite {
		//max
		moves := GetAllAggressiveMoves(board, isWhite)
		value := MinMaxScore{Move{}, -9999, []Move{}}
		for i := 0; i < len(moves); i++ {
			move := moves[i]
			currentBoard := board.Move(move)
			if isCheckMate(board, move, isWhite) {
				value = MinMaxScore{move, 222222, []Move{}}
				printWithDepth(move, value, depth, isWhite)
				break
			} else {
				curMax := GetBestMoveMinMax(currentBoard, false, depth-1)
				printWithDepth(move, curMax, depth, isWhite)
				if curMax.Score >= value.Score {
					if isDebug() && depth == 3 {
						fmt.Printf("new max, old: %d, new: %d \n", value.Score, curMax.Score)
						fmt.Println(move.Format())
						fmt.Println()
					}
					value = MinMaxScore{move, curMax.Score, append(curMax.History, curMax.Move)}
				}
			}
		}
		return value
	} else {
		//min
		moves := GetAllAggressiveMoves(board, isWhite)
		value := MinMaxScore{Move{}, 9999, []Move{}}
		for i := 0; i < len(moves); i++ {
			move := moves[i]
			currentBoard := board.Move(move)
			if isCheckMate(board, move, isWhite) {
				value = MinMaxScore{move, -222222, []Move{}}
				printWithDepth(move, value, depth, isWhite)
				break
			} else {
				curMin := GetBestMoveMinMax(currentBoard, true, depth-1)
				printWithDepth(move, curMin, depth, isWhite)
				if curMin.Score <= value.Score {
					if isDebug() && depth == 2 {
						fmt.Printf("\t\tnew min, old: %d, new: %d \n\t\t", value.Score, curMin.Score)
						fmt.Println(move.Format())
						fmt.Println()
					}
					value = MinMaxScore{move, curMin.Score, append(curMin.History, curMin.Move)}
					value.Move = move
				}
			}
		}
		return value
	}
}
func printWithDepth(move Move, score MinMaxScore, depth int, isWhite bool){
	if isDebug(){
		s := ""
		if depth ==2 {
			s = "\t"
		}
		if depth == 1 {
			s="\t\t"
		}
		fmt.Println(s , move.Format(), " ", isWhite, " ", score.Score)
	}
}
func isCheckMate(b Board, m Move, isWhite bool) bool {
	if isWhite && b.Grid[m.End.x][m.End.y] == 'k' || (!isWhite && b.Grid[m.End.x][m.End.y] == 'K') {
		return true
	}
	return false
}

type MinMaxScore struct {
	Move    Move
	Score   int
	History []Move
}

func GetAllAggressiveMoves(board Board, isWhite bool) []Move {
	vmod := 1
	whites, blacks := board.GetPieces()
	var colorPieces []Piece = whites
	if !isWhite {
		colorPieces = blacks
		vmod = -1
	}

	moves := make([]Move, 0, 20)
	for i := 0; i < len(colorPieces); i++ {
		p := colorPieces[i]
		if p.Value == 'P' || p.Value == 'p' {
			moves = append(moves, getPawnMoves(board.Grid, isWhite, vmod, p.Position, getWeightOfPiece(p))...)
		} else if p.Value == 'R' || p.Value == 'r' {
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, 1, 0, getWeightOfPiece(p))...)
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, 0, 1, getWeightOfPiece(p))...)
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, -1, 0, getWeightOfPiece(p))...)
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, -0, -1, getWeightOfPiece(p))...)
		} else if p.Value == 'B' || p.Value == 'b' {
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, 1, 1, getWeightOfPiece(p))...)
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, -1, -1, getWeightOfPiece(p))...)
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, -1, 1, getWeightOfPiece(p))...)
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, 1, -1, getWeightOfPiece(p))...)
		} else if p.Value == 'Q' || p.Value == 'q' {
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, 1, 0, getWeightOfPiece(p))...)
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, 0, 1, getWeightOfPiece(p))...)
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, -1, 0, getWeightOfPiece(p))...)
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, -0, -1, getWeightOfPiece(p))...)
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, 1, 1, getWeightOfPiece(p))...)
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, -1, -1, getWeightOfPiece(p))...)
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, 1, -1, getWeightOfPiece(p))...)
			moves = append(moves, getDiagonalMoves(board, p.Position, isWhite, -1, 1, getWeightOfPiece(p))...)
		} else if p.Value == 'N' || p.Value == 'n' {
			moves = append(moves, getKnightMoves(board, p.Position, isWhite, getWeightOfPiece(p))...)
		} else if p.Value == 'K' || p.Value == 'k' {
			moves = append(moves, getKingMoves(board, p.Position, isWhite, getWeightOfPiece(p))...)
		}
	}
	return moves
}
func getAvailableMoves(board Board, origin Position, isWhite bool, horizontalDirection int, verticalDirection int, weight int) []Move {
	moves := make([]Move, 0, 48)
	horizontalSign := horizontalDirection
	verticalSign := verticalDirection
	if horizontalDirection == 0 {
		horizontalSign = 1
	}
	if verticalDirection == 0 {
		verticalSign = 1
	}

	for x := origin.x + horizontalDirection; x >= 0 && x <= 7; x += horizontalSign {
		for y := origin.y + verticalDirection; y >= 0 && y <= 7; y += verticalSign {

			if x == origin.x && y == origin.y {
				return moves
			}
			target := board.Grid[x][y]
			if target == 0 {
				moves = append(moves, moveOf(origin.x, origin.y, x, y))
			} else if determineIfWhite(target) != isWhite {
				moves = append(moves, moveWithTakeOf(origin.x, origin.y, x, y, weight, getWeight(target)))
				return moves
			} else {
				return moves
			}
		}
	}
	return moves
}
func getDiagonalMoves(board Board, origin Position, isWhite bool, horizontalDirection int, verticalDirection int, weight int) []Move {
	x := origin.x + horizontalDirection
	y := origin.y + verticalDirection
	moves := make([]Move, 0, 16)
	for x >= 0 && x <= 7 && y >= 0 && y <= 7 {
		target := board.Grid[x][y]
		if target == 0 {
			moves = append(moves, moveOf(origin.x, origin.y, x, y))
		} else if determineIfWhite(target) != isWhite {
			moves = append(moves, moveWithTakeOf(origin.x, origin.y, x, y, weight, getWeight(target)))
			return moves
		} else {
			return moves
		}
		x += horizontalDirection
		y += verticalDirection
	}
	return moves
}
func getPawnMoves(grid [8][8]byte, isWhite bool, vmod int, origin Position, weight int) []Move {
	moves := make([]Move, 0, 48)
	var promotion byte = 0
	if isWhite && origin.y == 6 || (!isWhite && origin.y == 1) {
		promotion = 'q'
	}
	if grid[origin.x][origin.y+vmod] == 0 {
		move := moveOf(origin.x, origin.y, origin.x, origin.y+vmod)
		move.PromotionPiece = promotion
		moves = append(moves, move)
	}
	if origin.x < 7 && grid[origin.x+1][origin.y+vmod] != 0 && determineIfWhite(grid[origin.x+1][origin.y+vmod]) != isWhite {
		move := moveWithTakeOf(origin.x, origin.y, origin.x+1, origin.y+vmod, weight, getWeight(grid[origin.x+1][origin.y+vmod]))
		move.PromotionPiece = promotion
		moves = append(moves, move)
	}
	if origin.x > 0 && grid[origin.x-1][origin.y+vmod] != 0 && determineIfWhite(grid[origin.x-1][origin.y+vmod]) != isWhite {
		move := moveWithTakeOf(origin.x, origin.y, origin.x-1, origin.y+vmod, weight, getWeight(grid[origin.x-1][origin.y+vmod]))
		move.PromotionPiece = promotion
		moves = append(moves, move)
	}
	return moves
}
func getKnightMoves(board Board, origin Position, isWhite bool, weight int) []Move {
	moves := make([]Move, 0, 8)
	theoryMoves := []Position{
		{origin.x + 1, origin.y + 2},
		{origin.x + 2, origin.y + 1},
		{origin.x - 1, origin.y + 2},
		{origin.x - 2, origin.y + 1},
		{origin.x + 1, origin.y - 2},
		{origin.x + 2, origin.y - 1},
		{origin.x - 1, origin.y - 2},
		{origin.x - 2, origin.y - 1},
	}
	for i := 0; i < len(theoryMoves); i++ {
		m := theoryMoves[i]
		if !(m.x >= 0 && m.x <= 7 && m.y >= 0 && m.y <= 7) {
			
		} else {
			target := board.Grid[m.x][m.y]
			if target == 0 {
				moves = append(moves, moveOf(origin.x, origin.y, m.x, m.y))
			} else if determineIfWhite(target) != isWhite {
				moves = append(moves, moveWithTakeOf(origin.x, origin.y, m.x, m.y, weight, getWeight(target)))
			}
		}
	}
	return moves
}
func getKingMoves(board Board, origin Position, isWhite bool, weight int) []Move {
	moves := make([]Move, 0, 8)
	theoryMoves := []Position{
		{origin.x + 1, origin.y + 0},
		{origin.x + 0, origin.y + 1},
		{origin.x - 1, origin.y + 0},
		{origin.x + 0, origin.y - 1},
		{origin.x + 1, origin.y + 1},
		{origin.x - 1, origin.y - 1},
		{origin.x - 1, origin.y + 1},
		{origin.x + 1, origin.y - 1},
	}
	for i := 0; i < len(theoryMoves); i++ {
		m := theoryMoves[i]
		if m.x >= 0 && m.x <= 7 && m.y >= 0 && m.y <= 7 {
			target := board.Grid[m.x][m.y]
			if target == 0 {
				moves = append(moves, moveOf(origin.x, origin.y, m.x, m.y))
			} else if determineIfWhite(target) != isWhite {
				moves = append(moves, moveWithTakeOf(origin.x, origin.y, m.x, m.y, weight, getWeight(target)))
			}
		}
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
	parsedMoves := make([]Move, 0, 48)
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
	return Board{Grid: board}
}

type Board struct {
	Grid [8][8]byte
}

func (b Board) get(column byte, line int) byte {
	return b.Grid[column-97][line-1]
}
func (b Board) GetPieces() (whitePieces []Piece, blackPieces []Piece) {
	whitePieces = make([]Piece, 0, 16)
	blackPieces = make([]Piece, 0, 16)
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			square := b.Grid[x][y]
			if square != 0 {
				piece := Piece{square, Position{x, y}}
				if determineIfWhite(byte(square)) {
					whitePieces = append(whitePieces, piece)
				} else {
					blackPieces = append(blackPieces, piece)
				}
			}
		}
	}
	return whitePieces, blackPieces
}
func (b Board) Move(move Move) Board {
	piece := b.Grid[move.Begin.x][move.Begin.y]

	if move.PromotionPiece != 0 {
		piece = move.PromotionPiece
	}
	b.Grid[move.Begin.x][move.Begin.y] = 0
	b.Grid[move.End.x][move.End.y] = piece
	return b
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
	Begin          Position
	End            Position
	Value          int
	PromotionPiece byte
}

func (b Board) GetPositionalScore() int {
	whiteScore := 0
	blackScore := 0
	whites, blacks := b.GetPieces()
	for i := 0; i < len(whites); i++ {
		whiteScore += getWeightOfPiece(whites[i])
	}
	for i := 0; i < len(blacks); i++ {
		blackScore += getWeightOfPiece(blacks[i])
	}
	return whiteScore - blackScore
}
func getWeightOfPiece(p Piece) int {
	return getWeight(p.Value)
}
func getWeight(p byte) int {
	if p == 'P' || p == 'p' {
		return 10
	
	} else if p == 'B' || p == 'b' {
		return 30
	} else if p == 'N' || p == 'n' {
		return 30
	} else if p == 'R' || p == 'r' {
		return 50
	} else if p == 'Q' || p == 'q' {
		return 90
	} else if p == 'K' || p == 'k' {
		return 900
	} else {
		return 0
	}

}

func moveOf(ox int, oy int, tx int, ty int) Move {
	return Move{Begin: Position{x: ox, y: oy}, End: Position{x: tx, y: ty}}
}
func moveWithTakeOf(ox int, oy int, tx int, ty int, oWeight int, tWeight int) Move {
	return Move{Begin: Position{x: ox, y: oy}, End: Position{x: tx, y: ty}, Value: tWeight - oWeight}
}
func (m Move) Format() string {
	if m.PromotionPiece != 0 {
		return m.Begin.Format() + m.End.Format() + string(m.PromotionPiece)
	}
	return m.Begin.Format() + m.End.Format()
}
func printBoard(b [8][8]byte) {
	if isDebug() {
		data := [][]string{}

		for y := 7; y >= 0; y-- {
			var line []string
			line = append(line, fmt.Sprint(int(y+1)))
			for x := 0; x <= 7; x++ {
				line = append(line, string(b[x][y]))
			}
			data = append(data, line)
		}

		// table := tablewriter.NewWriter(os.Stdout)
		// table.SetFooter([]string{" ", "a", "b", "c", "d", "e", "f", "g", "h"})

		// for _, v := range data {
		// 	table.Append(v)
		// }
		// table.Render() // Send output
	}
}
func isDebug() bool {
	return false
}
