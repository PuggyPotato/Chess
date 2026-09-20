package main

import (
	"fmt"

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
			symbols := [...]string{"♙", "♘", "♗", "♖", "♕", "♔"}
			return symbols[p.Type]

		case 1:
			symbols := [...]string{"♟", "♞", "♝", "♜", "♛", "♚"}
			return symbols[p.Type]
		
		default:
			return " "
	}

}


func main() {
	board := Board{
		[8][8]*Piece{ 
			0: {&Piece{Rook, Black}, &Piece{Knight, Black}, &Piece{Bishop, Black}, &Piece{Queen, Black}, &Piece{King, Black}, &Piece{Bishop, Black}, &Piece{Knight, Black}, &Piece{Rook, Black}},
			1: {&Piece{Pawn, Black}, &Piece{Pawn, Black}, &Piece{Pawn, Black}, &Piece{Pawn, Black}, &Piece{Pawn, Black}, &Piece{Pawn, Black}, &Piece{Pawn, Black}, &Piece{Pawn, Black}},
			6: {&Piece{Pawn, White}, &Piece{Pawn, White}, &Piece{Pawn, White}, &Piece{Pawn, White}, &Piece{Pawn, White}, &Piece{Pawn, White}, &Piece{Pawn, White}, &Piece{Pawn, White}},
			7: {&Piece{Rook, White}, &Piece{Knight, White}, &Piece{Bishop, White}, &Piece{Queen, White}, &Piece{King, White}, &Piece{Bishop, White}, &Piece{Knight, White}, &Piece{Rook, White}},
		},
	}


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