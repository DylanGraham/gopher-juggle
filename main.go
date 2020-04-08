package main

import (
	"fmt"
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
	gopher, _, err = ebitenutil.NewImageFromFile("gopher.png", ebiten.FilterDefault)
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

// Update the game state
func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

// Draw the current game state
func (g *Game) Draw(screen *ebiten.Image) {
	g.vy += gravity
	if g.y > 450 {
		g.vy = -g.vy + gravity
	}

	// g.x += int(math.Round(g.vx))
	g.y += int(math.Round(g.vy))

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(80, 30)
	screen.DrawImage(gopher, op)

	op.GeoM.Reset()
	op.GeoM.Translate(float64(g.x), float64(g.y))
	// op.GeoM.Rotate()
	screen.DrawImage(ball, op)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\nvy: %0.2f",
		ebiten.CurrentTPS(), ebiten.CurrentFPS(), g.vy))
}

// Layout accepts the outside size (e.g., window size), and
// returns the game screen size.
// The game screen scale is automatically adjusted.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 500, 600
}

func main() {
	g := &Game{x: 125, y: 100}
	gravity = .3
	ebiten.SetWindowSize(500, 600)
	ebiten.SetWindowTitle("Gopher Juggle")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
