package main

func main() {
	gameBoard := NewBoard(9, 9, 10)
	cfg := &config{
		board: gameBoard,
		state: StateStart,
	}

	startRepl(cfg)
}
