package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	width  = 500
	height = 500
)

var (
	gopher1 *ebiten.Image
	gopher2 *ebiten.Image
	gopher3 *ebiten.Image
)

func init() {
	gopher1, _, err := ebitenutil.NewImageFromFile("ball.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	gopher2, _, err := ebitenutil.NewImageFromFile("ball.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	gopher3, _, err := ebitenutil.NewImageFromFile("ball.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	println(gopher1)
	println(gopher2)
	println(gopher3)
}

// Game struct
type Game struct{}

// Update the game state
func (g *Game) Update(screen *ebiten.Image) error {

	return nil
}

// Draw the current game state
func (g *Game) Draw(screen *ebiten.Image) {
}

// Layout accepts the outside size (e.g., window size), and
// returns the game screen size.
// The game screen scale is automatically adjusted.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	g := &Game{}
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
