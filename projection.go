package nzproj

type Projection interface {
	Forward(float64, float64) (float64, float64)
	Inverse(float64, float64) (float64, float64)
}

func Bounds(p Projection, bb []float64) []float64 {
	if !(len(bb) > 3) {
		return nil
	}
	var x0, y0, x1, y1 float64

	for n, xy := range [][]float64{{bb[0], bb[1]}, {bb[0], bb[3]}, {bb[2], bb[3]}, {bb[2], bb[1]}} {
		x, y := p.Forward(xy[0], xy[1])
		if n == 0 || x < x0 {
			x0 = x
		}
		if n == 0 || x > x1 {
			x1 = x
		}
		if n == 0 || y < y0 {
			y0 = y
		}
		if n == 0 || y > y1 {
			y1 = y
		}
	}

	return []float64{x0, y0, x1, y1}
}
