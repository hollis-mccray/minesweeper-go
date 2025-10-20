package main

func main() {
	gameBoard := NewBoard(9, 9, 10)
	cfg := &config{
		gameBoard: gameBoard,
		gamestate: StateStart,
	}

	startRepl(cfg)
}
