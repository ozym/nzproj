package nzproj

import (
	"math"
	"testing"
)

// NZGD2000	41° 43' 44.6973" S	172° 29' 58.2847" E
// NZCS2000	6,919,056.80 mN	2,958,434.27 mE
// NZGD2000	44° 40' 24.6238" S	167° 55' 26.6376" E
// NZCS2000	6,580,692.05 mN	2,597,661.53 mE

func TestLambertConformal(t *testing.T) {
	nzcs := NewNZCS2000()

	var tests = []struct {
		x   float64
		y   float64
		lat float64
		lon float64
	}{
		{
			2958434.27,
			6919056.80,
			-41.0 - 43.0/60.0 - 44.6973/3600.0,
			172.0 + 29.0/60.0 + 58.2847/3600.0,
		},
		{
			2597661.53,
			6580692.05,
			-44.0 - 40.0/60.0 - 24.6238/3600.0,
			167.0 + 55.0/60.0 + 26.6376/3600.0,
		},
		{
			2669944.56,
			6509882.69,
			-45.343956,
			168.791399,
		},
	}
	t.Run("inverse raw conversion", func(t *testing.T) {
		for _, v := range tests {
			lon, lat := nzcs.Inverse(v.x, v.y)

			if d := math.Abs(v.lat - lat); d > 0.0000001 {
				t.Errorf("latitude error: expected %g, got %g (difference %g)", v.lat, lat, d)
			}
			if d := math.Abs(v.lon - lon); d > 0.0000001 {
				t.Errorf("longitude error: expected %g, got %g (difference %g)", v.lon, lon, d)
			}

		}
	})
	t.Run("forward raw conversion", func(t *testing.T) {
		for _, v := range tests {
			x, y := nzcs.Forward(v.lon, v.lat)
			if d := math.Abs(v.x - x); d > 0.01 {
				t.Errorf("x error: expected %g, got %g (difference %g)", v.x, x, d)
			}
			if d := math.Abs(v.y - y); d > 0.01 {
				t.Errorf("y error: expected %g, got %g (difference %g)", v.y, y, d)
			}

		}
	})

	t.Run("inverse conversion", func(t *testing.T) {
		for _, v := range tests {
			lon, lat := nzcs.Inverse(v.x, v.y)

			if d := math.Abs(v.lat - lat); d > 0.0000001 {
				t.Errorf("latitude error: expected %g, got %g (difference %g)", v.lat, lat, d)
			}
			if d := math.Abs(v.lon - lon); d > 0.0000001 {
				t.Errorf("longitude error: expected %g, got %g (difference %g)", v.lon, lon, d)
			}

		}
	})

	t.Run("forward conversion", func(t *testing.T) {
		for _, v := range tests {
			x, y := nzcs.Forward(v.lon, v.lat)
			if d := math.Abs(v.x - x); d > 0.01 {
				t.Errorf("x error: expected %g, got %g (difference %g)", v.x, x, d)
			}
			if d := math.Abs(v.y - y); d > 0.01 {
				t.Errorf("y error: expected %g, got %g (difference %g)", v.y, y, d)
			}
		}
	})
}
