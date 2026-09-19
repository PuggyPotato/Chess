package main

import "fmt"

type PieceType int

const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

type Color int

const (
	White Color = iota
	Black
)

type Piece struct {
	Type  PieceType
	Color Color
}

type Board struct {
	Squares [8][8]*Piece
}

func main() {
	
}