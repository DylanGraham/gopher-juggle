package main

import (
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	gravity float64
	ball    *ebiten.Image
	gopher  *ebiten.Image
)

func init() {
	var err error
	ball, _, err = ebitenutil.NewImageFromFile("ball.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	gopher, _, err = ebitenutil.NewImageFromFile("ball.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

// Game struct
type Game struct {
	x  int
	y  int
	vx float64
	vy float64
}

func (g *Game) init() {

}

// Update the game state
func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

// Draw the current game state
func (g *Game) Draw(screen *ebiten.Image) {
	g.vy += gravity
	g.x += int(math.Round(g.vx))
	g.y += int(math.Round(g.vy))

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.x), float64(g.y))
	op.GeoM.Scale(.2, .2)
	// op.GeoM.Rotate()
	screen.DrawImage(ball, op)
}

// Layout accepts the outside size (e.g., window size), and
// returns the game screen size.
// The game screen scale is automatically adjusted.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 768
}

func main() {
	g := &Game{x: 50, y: 75, vx: 2}
	gravity = .1
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Gopher Juggle")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
