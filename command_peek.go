package main

func commandPeek(cfg *config, args []string) error {
	cfg.gameBoard.peek()
	return nil
}