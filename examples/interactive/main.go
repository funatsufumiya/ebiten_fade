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

var fader *fade.InteractiveFader

type Game struct{}

func (g *Game) Update() error {
	// Space or Enter toggles fade
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if fader.IsStarted() && !fader.IsFadeOutStarted() {
			fader.FadeOut(true)
		} else if fader.IsFinished() {
			fader.Start()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	fader.Alpha(func(a float32) {
		col := color.NRGBA{R: 255, G: 0, B: 0, A: uint8(a * 255)}
		vector.DrawFilledCircle(screen, 100, 100, 50, col, false)
	})
	ebitenutil.DebugPrintAt(screen, "click to toggle fade", 100, 100)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 300, 200
}

func main() {
	fader = fade.NewInteractiveFader(1.0, optional.Some[float32](1.0))
	fader.Start()
	
	ebiten.SetWindowSize(300, 200)
	ebiten.SetWindowTitle("fade interactive example (ebiten)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

// Input handling
func init() {
	ebiten.SetRunnableOnUnfocused(true)
}

func (g *Game) MousePressed(x, y int, button ebiten.MouseButton) {
	if fader.IsStarted() && !fader.IsFadeOutStarted() {
		fader.FadeOut(true)
	} else if fader.IsFinished() {
		fader.Start()
	}
}
