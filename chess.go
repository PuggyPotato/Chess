package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Board struct {
	Square [8][8]*Piece
}

type Piece struct {
	Type PieceType
	Color Color
}

type Color int 

const (
	White Color = iota
	Black
)

type PieceType int 

const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

func (p *Piece) String() string {

	switch p.Color {
		case 0:
			symbols := [6]string{"♙", "♘", "♗", "♖", "♕", "♔"}
			return symbols[p.Type]

		case 1:
			symbols := [6]string{"♟", "♞", "♝", "♜", "♛", "♚"}
			return symbols[p.Type]
		
		default:
			return " "
	}

}

//Print the board
func PrintBoard(board Board) {
	for _,row := range board.Square {
		for _,col := range row {
			if col != nil {
				fmt.Print(col.String())
			} else {
				fmt.Print(".")
			}
			fmt.Print(" ") 
		}
		fmt.Println()
	}
}

func isOnBoard(x int, y int) bool {
	return x >= 0 && x <= 7 && y <= 7 && y >= 0
}

type Position struct{
	X int
	Y int
}

func (b *Board) KnightMove(from Position, to Position, color Color) {
	//Valid Knight Moves
	possiblePositions := [8]Position{
						{X: from.X + 1, Y: from.Y + 2},
						{X: from.X - 1, Y: from.Y + 2},
						{X: from.X + 1, Y: from.Y - 2},
						{X: from.X - 1, Y: from.Y - 2},

						{X: from.X + 2, Y: from.Y + 1},
						{X: from.X - 2, Y: from.Y + 1},
						{X: from.X + 2, Y: from.Y - 1},
						{X: from.X - 2, Y: from.Y - 1},
					}	

	valid := make(map[Position]struct{})

	for _, position := range possiblePositions {
		if isOnBoard(position.X, position.Y) {
			if b.Square[position.Y][position.X] == nil || b.Square[position.Y][position.X].Color != Color(color) {
				valid[position] = struct{}{}
			}
		}
	}

	if !isOnBoard(to.X, to.Y) {
		log.Printf("Error: coord not on board, %v", to) // Log print for now, will change to return fmt.Error later
	}

	if _, ok := valid[to]; !ok {
		log.Print("Invalid Move")
		log.Print(valid)
		return
	}

	b.Square[from.Y][from.X] = nil
	b.Square[to.Y][to.X] = &Piece{Type: Knight, Color: Color(color)}
}

//Todo: Add Castle mechanism
func (b *Board) KingMove(from Position, to Position, color Color) {
	//Valid King Moves
	possiblePositions := [8]Position{
						{X: from.X + 1, Y: from.Y + 1},
						{X: from.X + 1, Y: from.Y},
						{X: from.X + 1, Y: from.Y - 1},

						{X: from.X, Y: from.Y + 1},
						{X: from.X, Y: from.Y - 1},

						{X: from.X -1, Y: from.Y + 1},
						{X: from.X -1, Y: from.Y - 1},
						{X: from.X -1, Y: from.Y},
					}	

	valid := make(map[Position]struct{})

	for _, position := range possiblePositions {
		if isOnBoard(position.X, position.Y) {
			if b.Square[position.Y][position.X] == nil || b.Square[position.Y][position.X].Color != Color(color) {
				valid[position] = struct{}{}
			}
		}
	}

	if !isOnBoard(to.X, to.Y) {
		log.Printf("Error: coord not on board, %v", to) // Log print for now, will change to return fmt.Error later
	}

	if _, ok := valid[to]; !ok {
		log.Print("Invalid Move")
		log.Print(valid)
		return
	}

	b.Square[from.Y][from.X] = nil
	b.Square[to.Y][to.X] = &Piece{Type: King, Color: Color(color)}
}

func (b *Board) RookMove(from Position, to Position, color Color) {
	rookDirections := []Position{
		{X:1, Y:0},
		{X:-1, Y:0},
		{X:0, Y:1},
		{X:0, Y:-1},
	}

	valid := b.slidingMoves(from, color, rookDirections)

	if !isOnBoard(to.X, to.Y) {
		log.Printf("Error: coord not on board, %v", to) // Log print for now, will change to return fmt.Error later
	}

	if _, ok := valid[to]; !ok {
		log.Print("Invalid Move")
		log.Print(valid)
		return
	}

	b.Square[from.Y][from.X] = nil
	b.Square[to.Y][to.X] = &Piece{Type: Rook, Color: Color(color)}

}

func (b *Board) slidingMoves(from Position, color Color, directions []Position) map[Position]struct{} {
	valid := make(map[Position]struct{})

	for _, direction := range directions {
		for step := 1; step <= 7; step++ {
			pos := Position{
				X: from.X + direction.X*step,
				Y: from.Y + direction.Y*step,
			}

			if !isOnBoard(pos.X, pos.Y) {
				break
			}

			if isOnBoard(pos.X, pos.Y) {
				if b.Square[pos.Y][pos.X] == nil{
					valid[pos] = struct{}{}
				} else if b.Square[pos.Y][pos.X].Color != Color(color){
					valid[pos] = struct{}{}
					break
				} else {
					break
				}
			}
		}
	}

	return valid
} 

func (b *Board) BishopMove(from Position, to Position, color Color) {
	rookDirections := []Position{
		{X:1, Y:1},
		{X:-1, Y:-1},
		{X:1, Y:-1},
		{X:-1, Y:1},
	}

	valid := b.slidingMoves(from, color, rookDirections)

	if !isOnBoard(to.X, to.Y) {
		log.Printf("Error: coord not on board, %v", to) // Log print for now, will change to return fmt.Error later
	}

	if _, ok := valid[to]; !ok {
		log.Print("Invalid Move")
		log.Print(valid)
		return
	}

	b.Square[from.Y][from.X] = nil
	b.Square[to.Y][to.X] = &Piece{Type: Bishop, Color: Color(color)}

}

func (b *Board) QueenMove(from Position, to Position, color Color) {
	rookDirections := []Position{
		{X:1, Y:1},
		{X:-1, Y:-1},
		{X:1, Y:-1},
		{X:-1, Y:1},
		{X:1, Y:0},
		{X:-1, Y:0},
		{X:0, Y:1},
		{X:0, Y:-1},
	}

	valid := b.slidingMoves(from, color, rookDirections)

	if !isOnBoard(to.X, to.Y) {
		log.Printf("Error: coord not on board, %v", to) // Log print for now, will change to return fmt.Error later
	}

	if _, ok := valid[to]; !ok {
		log.Print("Invalid Move")
		log.Print(valid)
		return
	}

	b.Square[from.Y][from.X] = nil
	b.Square[to.Y][to.X] = &Piece{Type: Queen, Color: Color(color)}

}


var mapToXCoord = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
}

func main() {
	board := Board{
		[8][8]*Piece{ 
			0: {&Piece{Rook, White}, &Piece{Knight, White}, &Piece{Bishop, White}, &Piece{Queen, White}, &Piece{King, White}, &Piece{Bishop, White}, &Piece{Knight, White}, &Piece{Rook, White}},
			//1: {&Piece{Pawn, White}, &Piece{Pawn, White}, &Piece{Pawn, White}, &Piece{Pawn, White}, &Piece{Pawn, White}, &Piece{Pawn, White}, &Piece{Pawn, White}, &Piece{Pawn, White}},
			6: {&Piece{Pawn, Black}, &Piece{Pawn, Black}, &Piece{Pawn, Black}, &Piece{Pawn, Black}, &Piece{Pawn, Black}, &Piece{Pawn, Black}, &Piece{Pawn, Black}, &Piece{Pawn, Black}},
			7: {&Piece{Rook, Black}, &Piece{Knight, Black}, &Piece{Bishop, Black}, &Piece{Queen, Black}, &Piece{King, Black}, &Piece{Bishop, Black}, &Piece{Knight, Black}, &Piece{Rook, Black}},
		},
	}

	PrintBoard(board)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) != 4 {
			fmt.Printf("Bad notation. want %d, got %d", 4, len(input))
			continue
		}

		fromY, err := strconv.Atoi(string(input[1]))
		if err != nil {
			log.Fatalf("Something went wrong with converting from.Y coordinates: %v", err)
		}
		from := Position{X: mapToXCoord[string(input[0])], Y: fromY - 1}

		toY, err := strconv.Atoi(string(input[3]))
		if err != nil {
			log.Fatalf("Something went wrong with converting to.Y coordinates: %v", err)
		}
		to := Position{X: mapToXCoord[string(input[2])], Y: toY - 1}

		if !isOnBoard(from.X, from.Y) || !isOnBoard(to.X, to.Y) {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		piece := board.Square[from.Y][from.X]
		if piece == nil {
			fmt.Print("Error: No Piece Exist on this square\n")
			continue
		}
		switch piece.Type {
			case 1: // Knight
				board.KnightMove(from, to, piece.Color)
			case 2: // Bishop
				board.BishopMove(from, to, piece.Color)
			case 3: // Rook
				board.RookMove(from, to, piece.Color)
			case 4: // Queen
				board.QueenMove(from, to, piece.Color)
			case 5: // King
				board.KingMove(from, to, piece.Color)
		}

		PrintBoard(board)
	}
}
