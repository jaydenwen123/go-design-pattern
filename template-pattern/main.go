package main

import tp "github.com/jaydenwen123/go-design-pattern/template-pattern/core"

func main() {
	var game *tp.Game
	game = tp.NewGame(tp.NewFootball())
	game.Play()
	game=tp.NewGame(&tp.Basketball{})
	game.Play()
}
