# Terminal Chess (MVP — archived, incomplete)
 
A chess rules engine and terminal interface written in Go, using coordinate
notation input (`e2e4`, not SAN). Built as a learning project. Archived
partway through — see **Status** below for exactly what works and what
doesn't.
 
## What this is
 
- Full board representation (`[8][8]*Piece`, nil = empty square)
- Legal move generation (ignoring check) for all six piece types:
  Pawn, Knight, Bishop, Rook, Queen, King
- Terminal input loop reading `e2e4`-style coordinate notation
- Move validation against generated legal moves, applied to the board
- Plain-text board printing (unicode chess glyphs, `.` for empty squares)
## What this is NOT (yet)
 
This was scoped as an MVP with a clear build order. It stops partway
through step 3. Specifically missing:
 
- **Turn alternation is not enforced.** Nothing currently stops you from
  moving the same color twice in a row, or moving out of turn. Every
  move function trusts whatever color the piece on the source square
  happens to be.
- **Check detection** — not implemented. Kings can currently move into
  check, and no piece's "legality" accounts for whether it would expose
  the king.
- **Checkmate / stalemate detection** — not implemented (depends on
  check detection above).
- **Castling, en passant, promotion** — not implemented. These were
  scoped as "if time allows" from the start.
- **Web version (Go backend + JS frontend)** — not started. Was planned
  to come after the terminal version's logic was fully proven correct.
## What's actually solid
 
- All six pieces' raw movement patterns are correctly implemented,
  including sliding-piece blocking/capture logic (rook/bishop/queen
  correctly stop at the first blocker and can capture but not slide
  past an enemy piece).
- Pawn logic correctly handles direction-by-color, the starting-rank
  double-move (checking both the intermediate and destination squares
  are empty), and diagonal-only captures (a pawn cannot move diagonally
  onto an empty square).
- Rook/Bishop/Queen share a single `slidingMoves(from, color, directions)`
  helper — direction vectors are passed in per piece, walking/blocking
  logic lives in one place.
## Known rough edges (non-blocking, cosmetic/dead-code)
 
- In each `*Move` function, the `!isOnBoard(to.X, to.Y)` check logs a
  warning but doesn't `return`/`continue` — harmless in practice since
  `main()` already rejects off-board input before dispatch, but it's
  unreachable-ish code.
- `slidingMoves` has a redundant double `isOnBoard` check.
- Terminal output is unicode glyphs + `.`, unaligned in some fonts/
  terminals — ugly-but-correct was the explicit bar for this MVP.
## Running it
 
```
go run main.go
```
 
Enter moves in `e2e4` format (from-square immediately followed by
to-square, no space). The board reprints after every accepted move.
Illegal moves are rejected with a log message; the board state doesn't
change.
 
## If picked back up later
 
The natural next steps, in order, would be:
 
1. Add a `currentTurn Color` variable in `main()`, check it before
   dispatching to any `*Move` function, flip it after a successful move.
2. Check detection: for a given color, generate all opponent legal
   moves and see if any lands on that color's king square.
3. Legal-move filtering: for each candidate move, simulate it, run (2)
   against the resulting board, revert if it leaves the mover's own
   king in check.
4. Checkmate/stalemate: `AllLegalMoves(color)` == empty, combined with
   whether that color is currently in check.
5. Castling / en passant / promotion as scoped extras.
6. Web layer: `GET /state`, `POST /move` over the same board/move-gen
   code, no game-logic changes needed.