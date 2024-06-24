package main

type Piece struct {
	pieceID        int
	icon           int
	canBeEnPassant bool
	teamID         int
}

var (
	blackKing   = Piece{'K', '\u2654', false, 1}
	blackQueen  = Piece{'Q', '\u2655', false, 1}
	blackRook   = Piece{'R', '\u2656', false, 1}
	blackKnight = Piece{'N', '\u2658', false, 1}
	blackBishop = Piece{'B', '\u2657', false, 1}
	blackPawn   = Piece{'P', '\u2659', false, 1}
	whiteKing   = Piece{'K', '\u2654', false, 0}
	whiteQueen  = Piece{'Q', '\u2655', false, 0}
	whiteRook   = Piece{'R', '\u2656', false, 0}
	whiteKnight = Piece{'N', '\u2658', false, 0}
	whiteBishop = Piece{'B', '\u2657', false, 0}
	whitePawn   = Piece{'P', '\u2659', false, 0}
	emptySquare = Piece{'0', ' ', false, -1}
)

func pawnPromote() string {
	for {
		fmt.Println("What will you promote to?(Q, N, B, R)")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed. Aborting.")
		}
		if match, err := regexp.MatchString(`^[QNBR]\n$`, input); err == nil && match {
			return input
		} else {
			fmt.Println("Improper input.")
		}
	}
}
