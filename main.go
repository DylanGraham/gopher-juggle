package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
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

type mode int

const (
	modeTitle mode = iota
	modeGame
	modeGameOver
)

// Game struct
type Game struct {
	mode mode
	x    int
	y    int
	vx   float64
	vy   float64
}

func (g *Game) reset() {
	g.x = 125
	g.y = 100
	g.vy = 0
}

// Update the game state
func (g *Game) Update(screen *ebiten.Image) error {
	switch g.mode {
	case modeTitle:
		if kickBall() {
			g.mode = modeGame
		}
	case modeGame:
		g.vy += gravity

		if g.y > 450 && kickBall() {
			g.vy = -g.vy + gravity
		}

		if g.y > 600 {
			g.mode = modeGameOver
		}

		// g.x += int(math.Round(g.vx))
		g.y += int(math.Round(g.vy))

	case modeGameOver:
		g.mode = modeTitle
	}
	return nil
}

func kickBall() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return true
	}
	return false
}

// Draw the current game state
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.mode {
	case modeTitle:
		screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(80, 30)
		screen.DrawImage(gopher, op)
	case modeGame:
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(80, 30)
		screen.DrawImage(gopher, op)

		op.GeoM.Reset()
		op.GeoM.Translate(float64(g.x), float64(g.y))
		// op.GeoM.Rotate()
		screen.DrawImage(ball, op)
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
	case modeGameOver:
		g.reset()
		g.mode = modeTitle
	}
}

// Layout accepts the outside size (e.g., window size), and
// returns the game screen size.
// The game screen scale is automatically adjusted.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 500, 600
}

func main() {
	gravity = .3
	g := &Game{x: 125, y: 100}
	g.mode = modeTitle
	ebiten.SetWindowSize(500, 600)
	ebiten.SetWindowTitle("Gopher Juggle")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
