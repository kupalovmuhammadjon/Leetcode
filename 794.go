func validTicTacToe(board []string) bool {
	b := map[string]int{}
	for _, row := range board {
		for _, e := range row {
			b[string(e)]++
		}
	}
	if b["O"] > b["X"] {
		return false
	} else if b["O"] < b["X"] && b["O"] != b["X"]-1 {
		return false
	}
	if twoWinners(board, b) {
		return false
	}
	return true
}

func twoWinners(board []string, b map[string]int) bool {
	xcount := 0
	ocount := 0
	if board[0] == "XXX" {
		xcount++
	} else if board[0] == "OOO" {
		ocount++
	}
	if board[1] == "XXX" {
		xcount++
	} else if board[1] == "OOO" {
		ocount++
	}
	if board[2] == "XXX" {
		xcount++
	} else if board[2] == "OOO" {
		ocount++
	}
	if board[0][0] == board[1][0] && board[1][0] == board[2][0] && board[2][0] == 'X' {
		xcount++
	} else if board[0][0] == board[1][0] && board[1][0] == board[2][0] && board[2][0] == 'O' {
		ocount++
	}
	if board[0][1] == board[1][1] && board[1][1] == board[2][1] && board[2][1] == 'X' {
		xcount++
	} else if board[0][1] == board[1][1] && board[1][1] == board[2][1] && board[2][1] == 'O' {
		ocount++
	}
	if board[0][2] == board[1][2] && board[1][2] == board[2][2] && board[2][2] == 'X' {
		xcount++
	} else if board[0][2] == board[1][2] && board[1][2] == board[2][2] && board[2][2] == 'O' {
		ocount++
	}
	// Diagonals
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[2][2] == 'X' {
		xcount++
	} else if board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[2][2] == 'O' {
		ocount++
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[2][0] == 'X' {
		xcount++
	} else if board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[2][0] == 'O' {
		ocount++
	}
	if xcount > 0 {
		if b["O"] == b["X"] {
			return true
		}
	} else if ocount > 0 {
		if b["O"] != b["X"] {
			return true
		}
	}
	return xcount > 0 && ocount > 0
}
