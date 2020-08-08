package nzproj

import (
	"math"
	"testing"
)

/**
NZGD2000	41° 43' 44.6973" S	172° 29' 58.2847" E
NZTM2000	5,380,181.71 mN		1,558,376.32 mE
NZGD2000	44° 40' 24.6238" S	167° 55' 26.6376" E
NZTM2000	5,040,771.40 mN		1,197,666.98 mE
**/

func TestTransverseMercator(t *testing.T) {
	nztm := NewNZTM2000()

	var tests = []struct {
		x   float64
		y   float64
		lat float64
		lon float64
	}{
		{
			1558376.32,
			5380181.71,
			-41.0 - 43.0/60.0 - 44.6973/3600.0,
			172.0 + 29.0/60.0 + 58.2847/3600.0,
		},
		{
			1197666.98,
			5040771.40,
			-44.0 - 40.0/60.0 - 24.6238/3600.0,
			167.0 + 55.0/60.0 + 26.6376/3600.0,
		},
		{
			1270291.48,
			4970217.78,
			-45.343956,
			168.791399,
		},
	}
	t.Run("inverse raw conversion", func(t *testing.T) {
		for _, v := range tests {
			lon, lat := nztm.Inverse(v.x, v.y)
			if d := math.Abs(v.lat - lat); d > 0.00001 {
				t.Errorf("latitude error: expected %f, got %f (difference %g)", v.lat, lat, d)
			}
			if d := math.Abs(v.lon - lon); d > 0.00001 {
				t.Errorf("longitude error: expected %f, got %f (difference %g)", v.lon, lon, d)
			}

		}
	})
	t.Run("inverse conversion", func(t *testing.T) {
		for _, v := range tests {
			lon, lat := nztm.Inverse(v.x, v.y)
			if d := math.Abs(v.lat - lat); d > 0.00001 {
				t.Errorf("latitude error: expected %f, got %f (difference %g)", v.lat, lat, d)
			}
			if d := math.Abs(v.lon - lon); d > 0.00001 {
				t.Errorf("longitude error: expected %f, got %f (difference %g)", v.lon, lon, d)
			}
		}
	})
	t.Run("forward raw conversion", func(t *testing.T) {
		for _, v := range tests {
			x, y := nztm.Forward(v.lon, v.lat)
			if d := math.Abs(v.x - x); d > 0.5 {
				t.Errorf("x error: expected %g, got %g (difference %g)", v.x, x, d)
			}
			if d := math.Abs(v.y - y); d > 0.5 {
				t.Errorf("y error: expected %g, got %g (difference %g)", v.y, y, d)
			}
		}
	})

	t.Run("forward conversion", func(t *testing.T) {
		for _, v := range tests {
			x, y := nztm.Forward(v.lon, v.lat)
			if d := math.Abs(v.x - x); d > 0.5 {
				t.Errorf("x error: expected %g, got %g (difference %g)", v.x, x, d)
			}
			if d := math.Abs(v.y - y); d > 0.5 {
				t.Errorf("y error: expected %g, got %g (difference %g)", v.y, y, d)
			}
		}
	})
}
