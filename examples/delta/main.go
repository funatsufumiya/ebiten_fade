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

	// fadein: 0.5, static: 1.0, fadeout: 0.5, delta: 100
	fade.Delta(t, 0.5, 1.0, 0.5, 100, func(delta float32) {
		col := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
		vector.DrawFilledCircle(screen, 100, 100+delta, 50, col, false)
	})

	// fadein: 0.5, static: 1.0, fadeout: 0.5, delta: 100
	fade.DeltaMore(t, 0.5, 1.0, optional.Some[float32](0.5), 100, func(delta, alpha, rateEasing, rateTime float32, phase fade.Phase) {
		col := color.NRGBA{R: 255, G: 0, B: 0, A: uint8(alpha * 255)}
		vector.DrawFilledCircle(screen, 220, 100+delta, 50, col, false)
	}, fade.Linear, fade.Out, fade.Linear, fade.Out)

	// fadein: 0.5, static: 1.0, fadeout: 0.5, delta: (30, 100)
	deltaVec := [2]float32{30, 100}
	fade.DeltaMore(t, 0.5, 1.0, optional.Some[float32](0.5), 1, func(delta, alpha, rateEasing, rateTime float32, phase fade.Phase) {
		col := color.NRGBA{R: 255, G: 0, B: 0, A: uint8(alpha * 255)}
		vector.DrawFilledCircle(screen, 340+deltaVec[0], 100+deltaVec[1], 50, col, false)
	}, fade.Linear, fade.Out, fade.Linear, fade.Out)

	// Advanced usage: show rateEasing, rateTime, phase
	fade.Advanced(
		t, 0.5, 1.0, optional.Some[float32](0.5),
		func(rateEasing, rateTime float32, phase fade.Phase) {
			x := float32(460)
			ebitenutil.DebugPrintAt(screen, "rateEasing: "+formatFloat(rateEasing), int(x), 50)
			ebitenutil.DebugPrintAt(screen, "rateTime: "+formatFloat(rateTime), int(x), 100)
			ebitenutil.DebugPrintAt(screen, "phase: "+fade.PhaseToString(phase), int(x), 150)

			barColor := color.NRGBA{100, 100, 100, 255}
			activeColor := color.NRGBA{255, 255, 255, 255}
			vector.DrawFilledRect(screen, x, 70, 100, 4, barColor, false)
			vector.DrawFilledRect(screen, x, 70, 100*rateEasing, 4, activeColor, false)
			vector.DrawFilledRect(screen, x, 120, 100, 4, barColor, false)
			vector.DrawFilledRect(screen, x, 120, 100*rateTime, 4, activeColor, false)
		},
		fade.Cubic, fade.In, fade.Cubic, fade.Out,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 600, 300
}

func formatFloat(f float32) string {
	return fmt.Sprintf("%.2f", f)
}

func main() {
	startTime = time.Now()
	ebiten.SetWindowSize(600, 300)
	ebiten.SetWindowTitle("fade delta example (ebiten)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
