package main

import (
	"image/color"

	"github.com/funatsufumiya/ebiten_fade/fade"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	optional "github.com/moznion/go-optional"
)

var fader *fade.NonInteractiveFader

type Game struct{}

func (g *Game) Update() error {
	// Left mouse click restarts fade
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if fader.IsFinished() {
			fader.Start()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	fader.Delta(100, func(delta float32) {
	    col := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	    vector.DrawFilledCircle(screen, 220, 100+delta, 50, col, false)
	})
	ebitenutil.DebugPrintAt(screen, "click to restart", 400, 100)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 400, 300
}

func main() {
	fader = fade.NewNonInteractiveFader(1.0, 1.0, optional.Some[float32](0.5))
	fader.Start()
	
	ebiten.SetWindowSize(400, 300)
	ebiten.SetWindowTitle("fade non-interactive example (ebiten)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}