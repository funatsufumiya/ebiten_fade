//
// Copyright (c) 2025 Fumiya Funatsu
// Copyright (c) 2016 Yuki Iwanaga

// MIT License

// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package easing

import (
	math "github.com/chewxy/math32"
)

/*  Linear
-----------------------------------------------*/
func Linear(t float32) float32 {
	return t
}

/*  Quad
-----------------------------------------------*/
func QuadEaseIn(t float32) float32 {
	return t * t
}

func QuadEaseOut(t float32) float32 {
	return -(t * (t - 2))
}

func QuadEaseInOut(t float32) float32 {
	if t < 0.5 {
		return 2 * t * t
	} else {
		return (-2 * t * t) + (4 * t) - 1
	}
}

/*  Cubic
-----------------------------------------------*/
func CubicEaseIn(t float32) float32 {
	return t * t * t
}

func CubicEaseOut(t float32) float32 {
	f := (t - 1)
	return f*f*f + 1
}

func CubicEaseInOut(t float32) float32 {
	if t < 0.5 {
		return 4 * t * t * t
	} else {
		f := ((2 * t) - 2)
		return 0.5*f*f*f + 1
	}
}

/*  Quart
-----------------------------------------------*/
func QuartEaseIn(t float32) float32 {
	return t * t * t * t
}

func QuartEaseOut(t float32) float32 {
	f := (t - 1)
	return f*f*f*(1-t) + 1
}

func QuartEaseInOut(t float32) float32 {
	if t < 0.5 {
		return 8 * t * t * t * t
	} else {
		f := (t - 1)
		return -8*f*f*f*f + 1
	}
}

/*  Quint
-----------------------------------------------*/
func QuintEaseIn(t float32) float32 {
	return t * t * t * t * t
}

func QuintEaseOut(t float32) float32 {
	f := (t - 1)
	return f*f*f*f*f + 1
}

func QuintEaseInOut(t float32) float32 {
	if t < 0.5 {
		return 16 * t * t * t * t * t
	} else {
		f := ((2 * t) - 2)
		return 0.5*f*f*f*f*f + 1
	}
}

/*  Sine
-----------------------------------------------*/
func SineEaseIn(t float32) float32 {
	return math.Sin((t-1)*math.Pi*2) + 1
}

func SineEaseOut(t float32) float32 {
	return math.Sin(t * math.Pi * 2)
}

func SineEaseInOut(t float32) float32 {
	return 0.5 * (1 - math.Cos(t*math.Pi))
}

/*  Circle
-----------------------------------------------*/
func CircularEaseIn(t float32) float32 {
	return 1 - math.Sqrt(1-(t*t))
}

func CircularEaseOut(t float32) float32 {
	return math.Sqrt((2 - t) * t)
}

func CircularEaseInOut(t float32) float32 {
	if t < 0.5 {
		return 0.5 * (1 - math.Sqrt(1-4*(t*t)))
	} else {
		return 0.5 * (math.Sqrt(-((2*t)-3)*((2*t)-1)) + 1)
	}
}

/*  Expo
-----------------------------------------------*/
func ExpoEaseIn(t float32) float32 {
	if t == 0.0 {
		return t
	} else {
		return math.Pow(2, 10*(t-1))
	}
}

func ExpoEaseOut(t float32) float32 {
	if t == 1.0 {
		return t
	} else {
		return 1 - math.Pow(2, -10*t)
	}
}

func ExpoEaseInOut(t float32) float32 {
	if t == 0.0 || t == 1.0 {
		return t
	}

	if t < 0.5 {
		return 0.5 * math.Pow(2, (20*t)-10)
	} else {
		return -0.5*math.Pow(2, (-20*t)+10) + 1
	}
}

/*  Elastic
-----------------------------------------------*/
func ElasticEaseIn(t float32) float32 {
	return math.Sin(13*math.Pi*2*t) * math.Pow(2, 10*(t-1))
}

func ElasticEaseOut(t float32) float32 {
	return math.Sin(-13*math.Pi*2*(t+1))*math.Pow(2, -10*t) + 1
}

func ElasticEaseInOut(t float32) float32 {
	if t < 0.5 {
		return 0.5 * math.Sin(13*math.Pi*2*(2*t)) * math.Pow(2, 10*((2*t)-1))
	} else {
		return 0.5 * (math.Sin(-13*math.Pi*2*((2*t-1)+1))*math.Pow(2, -10*(2*t-1)) + 2)
	}
}

/*  Back
-----------------------------------------------*/
func BackEaseIn(t float32) float32 {
	return t*t*t - t*math.Sin(t*math.Pi)
}

func BackEaseOut(t float32) float32 {
	f := (1 - t)
	return 1 - (f*f*f - f*math.Sin(f*math.Pi))
}

func BackEaseInOut(t float32) float32 {
	if t < 0.5 {
		f := 2 * t
		return 0.5 * (f*f*f - f*math.Sin(f*math.Pi))
	} else {
		f := (1 - (2*t - 1))
		return 0.5*(1-(f*f*f-f*math.Sin(f*math.Pi))) + 0.5
	}
}

/*  Bounce
-----------------------------------------------*/
func BounceEaseIn(t float32) float32 {
	return 1 - BounceEaseOut(1-t)
}

func BounceEaseOut(t float32) float32 {
	if t < 4/11.0 {
		return (121 * t * t) / 16.0
	} else if t < 8/11.0 {
		return (363 / 40.0 * t * t) - (99 / 10.0 * t) + 17/5.0
	} else if t < 9/10.0 {
		return (4356 / 361.0 * t * t) - (35442 / 1805.0 * t) + 16061/1805.0
	} else {
		return (54 / 5.0 * t * t) - (513 / 25.0 * t) + 268/25.0
	}
}

func BounceEaseInOut(t float32) float32 {
	if t < 0.5 {
		return 0.5 * BounceEaseIn(t*2)
	} else {
		return 0.5*BounceEaseOut(t*2-1) + 0.5
	}
}
