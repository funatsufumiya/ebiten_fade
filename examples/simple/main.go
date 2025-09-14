package main

import (
	"fmt"
	"image/color"
	"time"

	math "github.com/chewxy/math32"

	"github.com/funatsufumiya/ebiten_fade/fade"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	optional "github.com/moznion/go-optional"
)

var startTime time.Time

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// loop time: 2.5 (sec)
	t := math.Mod(float32(time.Since(startTime).Seconds()), 2.5)

	// fadein: 0.5, static: 1.0, fadeout: 0.5
	fade.Alpha(t, 0.5, 1.0, 0.5, func(a float32) {
		col := color.NRGBA{R: 255, G: 0, B: 0, A: uint8(a * 255)}
		vector.DrawFilledCircle(screen, 100, 100, 50, col, false)
	})

	// Advanced usage: show rateEasing, rateTime, phase
	fade.Advanced(
		t, 0.5, 1.0, optional.Some[float32](0.5),
		func(rateEasing, rateTime float32, phase fade.Phase) {
			ebitenutil.DebugPrintAt(screen, "rateEasing: "+formatFloat(rateEasing), 200, 50)
			ebitenutil.DebugPrintAt(screen, "rateTime: "+formatFloat(rateTime), 200, 100)
			ebitenutil.DebugPrintAt(screen, "phase: "+fade.PhaseToString(phase), 200, 150)

			// Indicator bars
			barColor := color.NRGBA{100, 100, 100, 255}
			activeColor := color.NRGBA{255, 255, 255, 255}
			vector.DrawFilledRect(screen, 200, 70, 100, 4, barColor, false)
			vector.DrawFilledRect(screen, 200, 70, float32(100*rateEasing), 4, activeColor, false)
			vector.DrawFilledRect(screen, 200, 120, 100, 4, barColor, false)
			vector.DrawFilledRect(screen, 200, 120, float32(100*rateTime), 4, activeColor, false)
		},
		fade.Cubic, fade.In, fade.Cubic, fade.Out,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 400, 200
}

func formatFloat(f float32) string {
	return fmt.Sprintf("%.2f", f)
}

func main() {
	startTime = time.Now()
	ebiten.SetWindowSize(400, 200)
	ebiten.SetWindowTitle("fade example (ebiten)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
