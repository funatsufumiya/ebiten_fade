package fade

import (
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
