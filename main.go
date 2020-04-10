package main

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"
	"math"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

var (
	gravity         float64
	ball            *ebiten.Image
	gopher          *ebiten.Image
	gopherKick      *ebiten.Image
	arcadeFont      font.Face
	smallArcadeFont font.Face
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
	gopherKick, _, err = ebitenutil.NewImageFromFile("gopher-kick.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	tt, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	smallArcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    smallFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

type mode int

const (
	modeTitle mode = iota
	modeGame
	modeGameOver
)

const (
	screenWidth   = 500
	screenHeight  = 600
	fontSize      = 32
	smallFontSize = fontSize / 2
	circle        = math.Pi * 2
)

// Game struct
type Game struct {
	mode   mode
	x      int
	y      int
	vx     float64
	vy     float64
	score  int
	radial float64
}

func (g *Game) reset() {
	g.x = 125
	g.y = 100
	g.vy = 0
	g.score = 0
	g.radial = 0
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
			g.score++
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
	var texts []string
	switch g.mode {
	case modeTitle:
		texts = []string{"Gopher Juggle", "", "", "", "", "", "", "", "", "PRESS SPACE KEY"}
		screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(80, 30)
		screen.DrawImage(gopher, op)
	case modeGame:
		screen.Fill(color.RGBA{0x66, 0x66, 0x66, 0xff})
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(80, 30)

		var gopherDisplay *ebiten.Image
		if kickBall() {
			gopherDisplay = gopherKick
		} else {
			gopherDisplay = gopher
		}
		screen.DrawImage(gopherDisplay, op)

		op.GeoM.Reset()
		w, h := ball.Size()
		g.radial += circle / 80
		if g.radial >= circle {
			g.radial = -circle
		}
		op.GeoM.Translate(-float64(w)/2.0, -float64(h)/2.0)
		op.GeoM.Rotate(g.radial)
		op.GeoM.Translate(float64(w)/2.0, float64(h)/2.0)
		op.GeoM.Translate(float64(g.x), float64(g.y))
		screen.DrawImage(ball, op)

		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
	case modeGameOver:
		texts = []string{"", "GAME OVER!"}
		g.reset()
		g.mode = modeTitle
	}

	for i, l := range texts {
		x := (screenWidth - len(l)*fontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*fontSize, color.Black)
	}
	scoreStr := fmt.Sprintf("%04d", g.score)
	text.Draw(screen, scoreStr, arcadeFont, screenWidth-len(scoreStr)*fontSize, fontSize, color.Black)
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
