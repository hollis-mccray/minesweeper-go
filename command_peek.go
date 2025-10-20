package main

func commandPeek(cfg *config, args []string) error {
	cfg.board.peek()
	return nil
}