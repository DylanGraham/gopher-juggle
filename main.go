package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/mp3"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

var (
	audioContext    *audio.Context
	kickPlayer      *audio.Player
	beatLoop        *audio.InfiniteLoop
	beatPlayer      *audio.Player
	gravity         float64
	ball            *ebiten.Image
	gopher          *ebiten.Image
	gopherKick      *ebiten.Image
	arcadeFont      font.Face
	smallArcadeFont font.Face
	gopherDisplay   *ebiten.Image
	signals         = make(chan *ebiten.Image, 1)
)

func init() {
	var err error

	// Audio
	audioContext, _ = audio.NewContext(44100)

	beat, err := mp3.Decode(audioContext, audio.BytesReadSeekCloser(beat_mp3))
	if err != nil {
		log.Fatal(err)
	}
	beatLoop = audio.NewInfiniteLoop(audio.ReadSeekCloser(beat), beat.Length()*4*44100)
	beatPlayer, err = audio.NewPlayer(audioContext, beatLoop)
	beatPlayer.SetVolume(.5)

	kickD, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(kick_wav))
	if err != nil {
		log.Fatal(err)
	}
	kickPlayer, err = audio.NewPlayer(audioContext, kickD)
	if err != nil {
		log.Fatal(err)
	}
	kickPlayer.SetVolume(.4)

	// Images
	img, _, err := image.Decode(bytes.NewReader(gopher_png))
	if err != nil {
		log.Fatal(err)
	}
	gopher, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	gopherDisplay = gopher

	img, _, err = image.Decode(bytes.NewReader(ball_png))
	if err != nil {
		log.Fatal(err)
	}
	ball, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	img, _, err = image.Decode(bytes.NewReader(gopherKick_png))
	if err != nil {
		log.Fatal(err)
	}
	gopherKick, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

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
	g.vx = 2
	g.vy = 0
	g.score = 0
	g.radial = 0
}

// Update the game state
func (g *Game) Update(screen *ebiten.Image) error {
	switch g.mode {
	case modeTitle:
		if kickBall(g) {
			g.mode = modeGame
		}
	case modeGame:
		g.vy += gravity

		if g.y > 400 && kickBall(g) {
			kickPlayer.Rewind()
			kickPlayer.Play()
			g.vy = -g.vy + gravity
			g.score++
		}

		if g.y > 600 {
			g.mode = modeGameOver
			return nil
		}

		if g.x > 400 || g.x < 30 {
			g.vx = -g.vx
		}

		g.x += int(math.Round(g.vx))
		g.y += int(math.Round(g.vy))

	case modeGameOver:
		if kickBall(g) {
			g.mode = modeTitle
			g.reset()
		}
	}
	return nil
}

func kickTimer() {
	signals <- gopherKick
	time.Sleep(150 * time.Millisecond)
	signals <- gopher
}

func kickBall(g *Game) bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.mode == modeGame {
			go kickTimer()
		}
		return true
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if g.mode == modeGame {
			go kickTimer()
		}
		return true
	}
	if len(inpututil.JustPressedTouchIDs()) > 0 {
		if g.mode == modeGame {
			go kickTimer()
		}
		return true
	}
	return false
}

func drawGopher(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(80, 30)
	screen.DrawImage(gopher, op)
}

// Draw the current game state
func (g *Game) Draw(screen *ebiten.Image) {
	var texts []string
	switch g.mode {
	case modeTitle:
		texts = []string{"", "", "", "", "", "Gopher Juggle", "", "", "PRESS SPACE KEY", "OR TOUCH"}
		drawGopher(screen)
	case modeGame:
		screen.Fill(color.RGBA{0x66, 0x66, 0x66, 0xff})
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(80, 30)

		select {
		case gopherDisplay = <-signals:
		default:
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

		// ebitenutil.DebugPrint(screen, fmt.Sprintf("%02.f", g.radial))
		// ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
	case modeGameOver:
		texts = []string{"", "", "", "", "", "GAME OVER!"}
		drawGopher(screen)
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
	g := &Game{}
	g.reset()
	g.mode = modeTitle
	ebiten.SetWindowSize(500, 600)
	ebiten.SetWindowTitle("Gopher Juggle")
	beatPlayer.Play()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
