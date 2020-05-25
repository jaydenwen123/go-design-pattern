package core

type Gamer interface {
	initialize()
	startPlay()
	stopPlay()
}

type Game struct {
	gamer Gamer
}

func NewGame(gamer Gamer) *Game {
	return &Game{gamer: gamer}
}

func (g *Game) Play() {
	g.gamer.initialize()
	g.gamer.startPlay()
	g.gamer.stopPlay()
}
