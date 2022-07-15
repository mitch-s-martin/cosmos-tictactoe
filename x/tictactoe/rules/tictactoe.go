package tictactoe

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

const (
	BOARD_DIM = 3
	PLAYER1   = "x"
	PLAYER2   = "o"
	NO_PIECE  = "*"
	ROW_SEP   = "|"
)

type Player struct {
	Symbol string
}

type Piece struct {
	Player Player
}

var PieceStrings = map[Player]string{
	FIRST_PLAYER:  PLAYER1,
	SECOND_PLAYER: PLAYER2,
	NO_PLAYER:     NO_PIECE,
}

var StringPieces = map[string]Piece{
	PLAYER1:  Piece{FIRST_PLAYER},
	PLAYER2:  Piece{SECOND_PLAYER},
	NO_PIECE: Piece{NO_PLAYER},
}

type Pos struct {
	X int
	Y int
}

var NO_POS = Pos{-1, -1}

var FIRST_PLAYER = Player{PLAYER1}
var SECOND_PLAYER = Player{PLAYER2}
var NO_PLAYER = Player{NO_PIECE}

var Players = map[string]Player{
	PLAYER1: FIRST_PLAYER,
	PLAYER2: SECOND_PLAYER,
}

var Opponents = map[Player]Player{
	FIRST_PLAYER:  SECOND_PLAYER,
	SECOND_PLAYER: FIRST_PLAYER,
}

var Usable = map[Pos]bool{}

func init() {
	// Initialize usable spaces
	for y := 0; y < BOARD_DIM; y++ {
		for x := 0; x < BOARD_DIM; x++ {
			Usable[Pos{X: x, Y: y}] = true
		}
	}
}

type Game struct {
	Pieces map[Pos]Piece
	Turn   Player
}

func NewGame() *Game {
	pieces := make(map[Pos]Piece)
	game := &Game{pieces, FIRST_PLAYER}
	return game
}

func (game *Game) PieceAt(pos Pos) bool {
	_, ok := game.Pieces[pos]
	return ok
}

func (game *Game) TurnIs(player Player) bool {
	return game.Turn == player
}

func (game *Game) Winner() Player {
	pieces := game.Pieces

	if pieces[Pos{0, 0}].Player.Symbol != NO_PIECE {
		//top row
		if pieces[Pos{0, 0}] == pieces[Pos{0, 1}] && pieces[Pos{0, 1}] == pieces[Pos{0, 2}] {
			return pieces[Pos{0, 0}].Player
		}

		//top L to bottom R diagonal
		if pieces[Pos{0, 0}] == pieces[Pos{1, 1}] && pieces[Pos{0, 0}] == pieces[Pos{2, 2}] {
			return pieces[Pos{0, 0}].Player
		}

		//left column
		if pieces[Pos{0, 0}] == pieces[Pos{1, 0}] && pieces[Pos{0, 0}] == pieces[Pos{2, 0}] {
			return pieces[Pos{0, 0}].Player
		}
	}

	if pieces[Pos{2, 2}].Player.Symbol != NO_PIECE {
		//bottom row
		if pieces[Pos{2, 2}] == pieces[Pos{2, 1}] && pieces[Pos{2, 2}] == pieces[Pos{2, 0}] {
			return pieces[Pos{2, 2}].Player

		}

		//right column
		if pieces[Pos{2, 2}] == pieces[Pos{1, 2}] && pieces[Pos{2, 2}] == pieces[Pos{0, 2}] {
			return pieces[Pos{2, 2}].Player
		}
	}

	//bottom L to top R diagonal
	if pieces[Pos{0, 2}].Player.Symbol != NO_PIECE && (pieces[Pos{0, 2}] == pieces[Pos{1, 1}] && pieces[Pos{1, 1}] == pieces[Pos{2, 0}]) {
		return pieces[Pos{0, 2}].Player
	}

	//middle row
	if pieces[Pos{0, 1}].Player.Symbol != NO_PIECE && (pieces[Pos{0, 1}] == pieces[Pos{1, 1}] && pieces[Pos{1, 1}] == pieces[Pos{2, 1}]) {
		return pieces[Pos{0, 1}].Player
	}

	//middle column
	if pieces[Pos{1, 0}].Player.Symbol != NO_PIECE && (pieces[Pos{1, 0}] == pieces[Pos{1, 1}] && pieces[Pos{1, 1}] == pieces[Pos{1, 2}]) {
		return pieces[Pos{1, 0}].Player
	}

	return NO_PLAYER
}

func (game *Game) ValidMove(loc Pos) bool {
	if game.PieceAt(loc) {
		return false
	}
	return true
}

func (game *Game) updateTurn() {
	opponent := Opponents[game.Turn]
	game.Turn = opponent
}

func (game *Game) Move(loc Pos) (captured Pos, err error) {
	err = nil
	//DETERMINE HOW TO DO TURNS
	// if !game.TurnIs(game.Pieces[src].Player) {
	// 	return NO_POS, errors.New(fmt.Sprintf("Not %v's turn", game.Pieces[src].Player))
	// }
	if !game.ValidMove(loc) {
		return NO_POS, errors.New(fmt.Sprintf("Invalid move: %v", loc))
	}

	game.Pieces[loc] = StringPieces[game.Turn.Symbol]

	game.updateTurn()
	return
}

func (game *Game) String() string {
	var buf bytes.Buffer
	for y := 0; y < BOARD_DIM; y++ {
		for x := 0; x < BOARD_DIM; x++ {
			pos := Pos{x, y}
			if game.PieceAt(pos) {
				piece := game.Pieces[pos]
				val := PieceStrings[piece.Player]
				buf.WriteString(val)
			} else {
				buf.WriteString(PieceStrings[NO_PLAYER])
			}
		}
		if y < (BOARD_DIM - 1) {
			buf.WriteString(ROW_SEP)
		}
	}
	return buf.String()
}

func ParsePiece(s string) (Piece, bool) {
	piece, ok := StringPieces[s]
	return piece, ok
}

func Parse(s string) (*Game, error) {
	if len(s) != BOARD_DIM*BOARD_DIM+(BOARD_DIM-1) {
		return nil, errors.New(fmt.Sprintf("invalid board string: %v", s))
	}
	var pieceCount = 0

	pieces := make(map[Pos]Piece)
	result := &Game{pieces, FIRST_PLAYER}
	for y, row := range strings.Split(s, ROW_SEP) {
		for x, c := range strings.Split(row, "") {
			if x >= BOARD_DIM || y >= BOARD_DIM {
				return nil, errors.New(fmt.Sprintf("invalid board, piece out of bounds: %v, %v", x, y))
			}
			if piece, ok := ParsePiece(c); !ok {
				return nil, errors.New(fmt.Sprintf("invalid board, invalid piece at %v, %v", x, y))
			} else if piece.Player.Symbol != NO_PIECE {
				result.Pieces[Pos{x, y}] = piece
				pieceCount++
			}
		}
	}
	if pieceCount%2 != 0 {
		result.Turn = SECOND_PLAYER
	}
	return result, nil
}
