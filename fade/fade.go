package fade

import (
	"time"

	"github.com/funatsufumiya/ebiten_fade/fade/easing"
	optional "github.com/moznion/go-optional"
)

// EasingFunction represents the type of easing function.
type EasingFunction int
const (
	Linear EasingFunction = iota
	Quad
	Cubic
	Quart
	Quint
	Sine
	Expo
	Circular
	Back
	Elastic
	Bounce
)

// EasingType represents the direction/type of easing.
type EasingType int
const (
	In EasingType = iota
	Out
	InOut
)

// Phase represents the fade phase.
type Phase int
const (
	FadeIn Phase = iota
	Static
	FadeOut
)

// Use optional.Optional[float32] for optional fadeOut

// Advanced applies a fade effect with full control (rateEasing, rateTime, phase) and custom callback.
func Advanced(
	t, fadeIn, static float32,
	fadeOut optional.Option[float32],
	fn func(rateEasing, rateTime float32, phase Phase),
	easingFuncIn EasingFunction, easingTypeIn EasingType, easingFuncOut EasingFunction, easingTypeOut EasingType,
) {
	total := fadeIn + static
	if fadeOut.IsSome() {
		v, _ := fadeOut.Take()
		total += v
	}

	var (
		rateEasing, rateTime float32
		phase Phase
	)

	switch {
	case t < 0:
		rateEasing = 0
		rateTime = 0
		phase = FadeIn
	case t < fadeIn:
		rateTime = t / fadeIn
		phase = FadeIn
		rateEasing = applyEasing(easingFuncIn, easingTypeIn, rateTime)
	case t < fadeIn+static:
		rateTime = 1
		rateEasing = 1
		phase = Static
	case fadeOut.IsSome() && t < total:
		v, _ := fadeOut.Take()
		rateTime = (t - fadeIn - static) / v
		phase = FadeOut
		rateEasing = applyEasing(easingFuncOut, easingTypeOut, rateTime)
	default:
		rateEasing = 1
		rateTime = 1
		phase = FadeOut
	}
	fn(rateEasing, rateTime, phase)
}

// Alpha applies a simple fade-in/out effect.
func Alpha(t, fadeIn, static, fadeOut float32, fn func(alpha float32)) {
	AlphaMore(t, fadeIn, static, optional.Some[float32](fadeOut), func(alpha, _rateEasing, _rateTime float32, _phase Phase) {
		fn(alpha)
	}, Linear, Out, Linear, Out)
}

// AlphaMore applies a fade effect with more control and callback details.
func AlphaMore(
	t, fadeIn, static float32,
	fadeOut optional.Option[float32],
	fn func(alpha, rateEasing, rateTime float32, phase Phase),
	easingFuncIn EasingFunction, easingTypeIn EasingType, easingFuncOut EasingFunction, easingTypeOut EasingType,
) {
	Advanced(t, fadeIn, static, fadeOut, func(rateEasing, rateTime float32, phase Phase) {
		var alpha float32
		switch phase {
		case FadeIn:
			alpha = rateEasing
		case Static:
			alpha = 1
		case FadeOut:
			alpha = 1 - rateEasing
		}
		fn(alpha, rateEasing, rateTime, phase)
	}, easingFuncIn, easingTypeIn, easingFuncOut, easingTypeOut)

}

func PhaseToString(p Phase) string {
	switch p {
	case FadeIn:
		return "FadeIn"
	case Static:
		return "Static"
	case FadeOut:
		return "FadeOut"
	}
	return "Unknown"
}

// applyEasing is a stub for go-easing function call. Replace with actual go-easing usage.
func applyEasing(funcType EasingFunction, typeType EasingType, t float32) float32 {
	switch funcType {
	case Linear:
		return t
	case Quad:
		switch typeType {
		case In:
			return easing.QuadEaseIn(t)
		case Out:
			return easing.QuadEaseOut(t)
		case InOut:
			return easing.QuadEaseInOut(t)
		}
	case Cubic:
		switch typeType {
		case In:
			return easing.CubicEaseIn(t)
		case Out:
			return easing.CubicEaseOut(t)
		case InOut:
			return easing.CubicEaseInOut(t)
		}
	case Quart:
		switch typeType {
		case In:
			return easing.QuartEaseIn(t)
		case Out:
			return easing.QuartEaseOut(t)
		case InOut:
			return easing.QuartEaseInOut(t)
		}
	case Quint:
		switch typeType {
		case In:
			return easing.QuintEaseIn(t)
		case Out:
			return easing.QuintEaseOut(t)
		case InOut:
			return easing.QuintEaseInOut(t)
		}
	case Sine:
		switch typeType {
		case In:
			return easing.SineEaseIn(t)
		case Out:
			return easing.SineEaseOut(t)
		case InOut:
			return easing.SineEaseInOut(t)
		}
	case Expo:
		switch typeType {
		case In:
			return easing.ExpoEaseIn(t)
		case Out:
			return easing.ExpoEaseOut(t)
		case InOut:
			return easing.ExpoEaseInOut(t)
		}
	case Circular:
		switch typeType {
		case In:
			return easing.CircularEaseIn(t)
		case Out:
			return easing.CircularEaseOut(t)
		case InOut:
			return easing.CircularEaseInOut(t)
		}
	case Back:
		switch typeType {
		case In:
			return easing.BackEaseIn(t)
		case Out:
			return easing.BackEaseOut(t)
		case InOut:
			return easing.BackEaseInOut(t)
		}
	case Elastic:
		switch typeType {
		case In:
			return easing.ElasticEaseIn(t)
		case Out:
			return easing.ElasticEaseOut(t)
		case InOut:
			return easing.ElasticEaseInOut(t)
		}
	case Bounce:
		switch typeType {
		case In:
			return easing.BounceEaseIn(t)
		case Out:
			return easing.BounceEaseOut(t)
		case InOut:
			return easing.BounceEaseInOut(t)
		}
	}
	return t // fallback: linear
}

// Delta applies a simple delta fade effect.
func Delta(t, fadeIn, static, fadeOut, delta float32, fn func(delta float32)) {
	DeltaMore(t, fadeIn, static, optional.Some[float32](fadeOut), delta, func(deltaVal, _alpha, _rateEasing, _rateTime float32, _phase Phase) {
		fn(deltaVal)
	}, Linear, Out, Linear, Out)
}

// DeltaMore applies a delta fade effect with more control and callback details.
func DeltaMore(
	t, fadeIn, static float32,
	fadeOut optional.Option[float32],
	delta float32,
	fn func(delta, alpha, rateEasing, rateTime float32, phase Phase),
	easingFuncIn EasingFunction, easingTypeIn EasingType, easingFuncOut EasingFunction, easingTypeOut EasingType,
) {
	Advanced(t, fadeIn, static, fadeOut, func(rateEasing, rateTime float32, phase Phase) {
		var alpha, deltaVal float32
		switch phase {
		case FadeIn:
			alpha = rateEasing
			deltaVal = rateEasing * delta
		case Static:
			alpha = 1
			deltaVal = delta
		case FadeOut:
			alpha = 1 - rateEasing
			deltaVal = alpha * delta
		}
		fn(deltaVal, alpha, rateEasing, rateTime, phase)
	}, easingFuncIn, easingTypeIn, easingFuncOut, easingTypeOut)
}

// --- Fader types ---
func (f *InteractiveFader) Stop() {
	f.started = false
	f.startTime = optional.None[float64]()
	f.fadeoutStartedTime = optional.None[float64]()
}

// InteractiveFader provides fade-in/fade-out timer for interactive usage.
type InteractiveFader struct {
	FadeInSec  float32
	FadeOutSec optional.Option[float32]
	started    bool
	startTime  optional.Option[float64]
	fadeoutStartedTime optional.Option[float64]
}

func NewInteractiveFader(fadeInSec float32, fadeOutSec optional.Option[float32]) *InteractiveFader {
	return &InteractiveFader{
		FadeInSec:  fadeInSec,
		FadeOutSec: fadeOutSec,
	}
}

func (f *InteractiveFader) Start() {
	f.started = true
	f.startTime = optional.Some[float64](float64(time.Now().UnixNano()) / 1e9)
	f.fadeoutStartedTime = optional.None[float64]()
}

func (f *InteractiveFader) FadeOut(immediate bool) {
       if !f.started || !f.FadeOutSec.IsSome() {
	       return
       }
       now := float64(time.Now().UnixNano()) / 1e9
       if !f.startTime.IsSome() {
	       f.Start()
       }
       f.fadeoutStartedTime = optional.Some[float64](now)
       if immediate {
		startValRaw, _ := f.startTime.Value()
		startVal := startValRaw.(float64)
		elapsed := now - startVal
	       if elapsed < float64(f.FadeInSec) {
		       diffIn := float64(f.FadeInSec) - elapsed
		       diffInRate := diffIn / float64(f.FadeInSec)
		       fadeOutVal, _ := f.FadeOutSec.Take()
		       diffOut := diffInRate * float64(fadeOutVal)
					   f.startTime = optional.Some[float64](startVal - (diffIn + diffOut))
					   fadeoutValRaw, _ := f.fadeoutStartedTime.Value()
					   fadeoutVal := fadeoutValRaw.(float64)
					   f.fadeoutStartedTime = optional.Some[float64](fadeoutVal - (diffIn + diffOut))
	       }
       }
}

func (f *InteractiveFader) IsStarted() bool {
	return f.started && f.startTime.IsSome()
}

func (f *InteractiveFader) IsFadeOutStarted() bool {
	return f.fadeoutStartedTime.IsSome()
}

func (f *InteractiveFader) IsFinished() bool {
       if !f.started || !f.startTime.IsSome() {
	       return false
       }
	startValRaw, _ := f.startTime.Value()
	startVal := startValRaw.(float64)
	elapsed := float32(float64(time.Now().UnixNano())/1e9 - startVal)
       if f.fadeoutStartedTime.IsSome() && f.FadeOutSec.IsSome() {
	       fadeOutVal, _ := f.FadeOutSec.Take()
	       return elapsed > f.FadeInSec + float32(fadeOutVal)
       }
       return false
}

func (f *InteractiveFader) Alpha(fn func(alpha float32)) {
       if !f.startTime.IsSome() {
	       fn(0)
	       return
       }
	startValRaw, _ := f.startTime.Value()
	startVal := startValRaw.(float64)
	t := float32(float64(time.Now().UnixNano())/1e9 - startVal)
       var staticSec float32 = 0
       var fadeOut float32 = 0
       if f.fadeoutStartedTime.IsSome() && f.FadeOutSec.IsSome() {
	       fadeOutVal, _ := f.FadeOutSec.Take()
			   fadeoutValRaw, _ := f.fadeoutStartedTime.Value()
			   fadeoutVal := fadeoutValRaw.(float64)
			   staticSec = float32(fadeoutVal - startVal - float64(f.FadeInSec))
	       if staticSec < 0 {
		       staticSec = 0
	       }
	       fadeOut = fadeOutVal
       }
       Alpha(t, f.FadeInSec, staticSec, fadeOut, fn)
}

func (f *InteractiveFader) AlphaMore(fn func(alpha, rateEasing, rateTime float32, phase Phase), easingFuncIn EasingFunction, easingTypeIn EasingType, easingFuncOut EasingFunction, easingTypeOut EasingType) {
       if !f.startTime.IsSome() {
	       fn(0, 0, 0, FadeIn)
	       return
       }
	startValRaw, _ := f.startTime.Value()
	startVal := startValRaw.(float64)
	t := float32(float64(time.Now().UnixNano())/1e9 - startVal)
       var staticSec float32 = 0
       var fadeOut float32 = 0
       if f.fadeoutStartedTime.IsSome() && f.FadeOutSec.IsSome() {
	       fadeOutVal, _ := f.FadeOutSec.Take()
			   fadeoutValRaw, _ := f.fadeoutStartedTime.Value()
			   fadeoutVal := fadeoutValRaw.(float64)
			   staticSec = float32(fadeoutVal - startVal - float64(f.FadeInSec))
	       if staticSec < 0 {
		       staticSec = 0
	       }
	       fadeOut = fadeOutVal
       }
       AlphaMore(t, f.FadeInSec, staticSec, optional.Some[float32](fadeOut), fn, easingFuncIn, easingTypeIn, easingFuncOut, easingTypeOut)
}

func (f *InteractiveFader) Delta(delta float32, fn func(delta float32)) {
       if !f.startTime.IsSome() {
	       fn(0)
	       return
       }
	startValRaw, _ := f.startTime.Value()
	startVal := startValRaw.(float64)
	t := float32(float64(time.Now().UnixNano())/1e9 - startVal)
       var staticSec float32 = 0
       var fadeOut float32 = 0
       if f.fadeoutStartedTime.IsSome() && f.FadeOutSec.IsSome() {
	       fadeOutVal, _ := f.FadeOutSec.Take()
			   fadeoutValRaw, _ := f.fadeoutStartedTime.Value()
			   fadeoutVal := fadeoutValRaw.(float64)
			   staticSec = float32(fadeoutVal - startVal - float64(f.FadeInSec))
	       if staticSec < 0 {
		       staticSec = 0
	       }
	       fadeOut = fadeOutVal
       }
       Delta(t, f.FadeInSec, staticSec, fadeOut, delta, fn)
}

func (f *InteractiveFader) DeltaMore(delta float32, fn func(delta, alpha, rateEasing, rateTime float32, phase Phase), easingFuncIn EasingFunction, easingTypeIn EasingType, easingFuncOut EasingFunction, easingTypeOut EasingType) {
       if !f.startTime.IsSome() {
	       fn(0, 0, 0, 0, FadeIn)
	       return
       }
	startValRaw, _ := f.startTime.Value()
	startVal := startValRaw.(float64)
	t := float32(float64(time.Now().UnixNano())/1e9 - startVal)
       var staticSec float32 = 0
       var fadeOut float32 = 0
       if f.fadeoutStartedTime.IsSome() && f.FadeOutSec.IsSome() {
	       fadeOutVal, _ := f.FadeOutSec.Take()
			   fadeoutValRaw, _ := f.fadeoutStartedTime.Value()
			   fadeoutVal := fadeoutValRaw.(float64)
			   staticSec = float32(fadeoutVal - startVal - float64(f.FadeInSec))
	       if staticSec < 0 {
		       staticSec = 0
	       }
	       fadeOut = fadeOutVal
       }
       DeltaMore(t, f.FadeInSec, staticSec, optional.Some[float32](fadeOut), delta, fn, easingFuncIn, easingTypeIn, easingFuncOut, easingTypeOut)
}

// NonInteractiveFader provides fade-in/static/fade-out timer for non-interactive usage.
type NonInteractiveFader struct {
	FadeInSec  float32
	StaticSec  float32
	FadeOutSec optional.Option[float32]
	started    bool
	startTime  time.Time
}

func NewNonInteractiveFader(fadeInSec, staticSec float32, fadeOutSec optional.Option[float32]) *NonInteractiveFader {
	return &NonInteractiveFader{
		FadeInSec:  fadeInSec,
		StaticSec:  staticSec,
		FadeOutSec: fadeOutSec,
	}
}

func (f *NonInteractiveFader) Start() {
	f.started = true
	f.startTime = time.Now()
}

func (f *NonInteractiveFader) IsStarted() bool {
	return f.started
}

func (f *NonInteractiveFader) IsFinished() bool {
	if !f.started {
		return false
	}
	total := f.FadeInSec + f.StaticSec
	if f.FadeOutSec.IsSome() {
		v, _ := f.FadeOutSec.Take()
		total += v
	}
	return float32(time.Since(f.startTime).Seconds()) > total
}

func (f *NonInteractiveFader) Delta(delta float32, fn func(delta float32)) {
	t := float32(time.Since(f.startTime).Seconds())
	var fadeOut float32
	if f.FadeOutSec.IsSome() {
		v, _ := f.FadeOutSec.Take()
		fadeOut = v
	} else {
		fadeOut = 0
	}
	Delta(t, f.FadeInSec, f.StaticSec, fadeOut, delta, func(d float32) {
		fn(d)
	})
}

func (f *NonInteractiveFader) DeltaMore(delta float32, fn func(delta, alpha, rateEasing, rateTime float32, phase Phase), easingFuncIn EasingFunction, easingTypeIn EasingType, easingFuncOut EasingFunction, easingTypeOut EasingType) {
	t := float32(time.Since(f.startTime).Seconds())
	var fadeOut float32
	if f.FadeOutSec.IsSome() {
		v, _ := f.FadeOutSec.Take()
		fadeOut = v
	} else {
		fadeOut = 0
	}
	DeltaMore(t, f.FadeInSec, f.StaticSec, optional.Some[float32](fadeOut), delta, fn, easingFuncIn, easingTypeIn, easingFuncOut, easingTypeOut)
}

func (f *NonInteractiveFader) Alpha(fn func(alpha float32)) {
	t := float32(time.Since(f.startTime).Seconds())
	var fadeOut float32
	if f.FadeOutSec.IsSome() {
		v, _ := f.FadeOutSec.Take()
		fadeOut = v
	} else {
		fadeOut = 0
	}
	Alpha(t, f.FadeInSec, f.StaticSec, fadeOut, fn)
}

func (f *NonInteractiveFader) AlphaMore(fn func(alpha, rateEasing, rateTime float32, phase Phase), easingFuncIn EasingFunction, easingTypeIn EasingType, easingFuncOut EasingFunction, easingTypeOut EasingType) {
	t := float32(time.Since(f.startTime).Seconds())
	var fadeOut float32
	if f.FadeOutSec.IsSome() {
		v, _ := f.FadeOutSec.Take()
		fadeOut = v
	} else {
		fadeOut = 0
	}
	AlphaMore(t, f.FadeInSec, f.StaticSec, optional.Some[float32](fadeOut), fn, easingFuncIn, easingTypeIn, easingFuncOut, easingTypeOut)
}
