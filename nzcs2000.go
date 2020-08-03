package nzproj

import (
	"math"
	"sync"
)

/*
	New Zealand Continental Shelf Lambert Conformal 2000

	The extents of NZCS2000 are 160° E to 170° W and 25° S to 60° S.

	A Lambert conformal conic projection was chosen rather than the Transverse Mercator projection
	because the latter becomes excessively distorted when extended over large longitudinal ranges.

	https://www.linz.govt.nz/data/geodetic-system/datums-projections-heights/projections/new-zealand-continental-shelf-lambert
*/

type LambertConformal struct {
	a  float64
	f0 float64
	ϕ1 float64
	ϕ2 float64
	ϕ0 float64
	λ0 float64
	n0 float64
	e0 float64

	_e, _n, _F, _ρ0 struct {
		once  sync.Once
		value float64
	}
}

func (lc *LambertConformal) e() float64 {
	lc._e.once.Do(func() {
		lc._e.value = math.Sqrt(2.0*lc.f0 - math.Pow(lc.f0, 2.0))
	})
	return lc._e.value
}

func (lc *LambertConformal) m(l float64) float64 {
	return math.Cos(l) / math.Sqrt(1.0-math.Pow(lc.e(), 2.0)*math.Pow(math.Sin(l), 2.0))
}
func (lc *LambertConformal) t(l float64) float64 {
	return math.Tan((math.Pi/4.0)-(l/2.0)) / math.Pow((1.0-lc.e()*math.Sin(l))/(1.0+lc.e()*math.Sin(l)), lc.e()/2.0)
}
func (lc *LambertConformal) f() float64 {
	lc._F.once.Do(func() {
		lc._F.value = lc.m(lc.ϕ1) / (lc.n() * math.Pow(lc.t(lc.ϕ1), lc.n()))
	})
	return lc._F.value
}

func (lc *LambertConformal) ρ(l float64) float64 {
	return lc.a * lc.f() * math.Pow(lc.t(l), lc.n())
}

func (lc *LambertConformal) n() float64 {
	lc._n.once.Do(func() {
		lc._n.value = (math.Log(lc.m(lc.ϕ1)) - math.Log(lc.m(lc.ϕ2))) /
			(math.Log(lc.t(lc.ϕ1)) - math.Log(lc.t(lc.ϕ2)))
	})
	return lc._n.value
}

func (lc *LambertConformal) ρ0() float64 {
	lc._ρ0.once.Do(func() {
		lc._ρ0.value = lc.ρ(lc.ϕ0)
	})
	return lc._ρ0.value
}

func (lc LambertConformal) Forward(lon, lat float64) (float64, float64) {
	return lc.forward(math.Pi*lon/180.0, math.Pi*lat/180.0)
}

func (lc LambertConformal) forward(λ, ϕ float64) (float64, float64) {
	y := lc.n() * (λ - lc.λ0)

	X := lc.e0 + lc.ρ(ϕ)*math.Sin(y)
	Y := lc.n0 + lc.ρ0() - lc.ρ(ϕ)*math.Cos(y)

	return X, Y
}

func (lc LambertConformal) Inverse(x, y float64) (float64, float64) {
	λ, ϕ := lc.inverse(x, y)

	return 180.0 * λ / math.Pi, 180.0 * ϕ / math.Pi
}

func (lc LambertConformal) inverse(x, y float64) (float64, float64) {

	E := x - lc.e0
	N := y - lc.n0

	λ := lc.λ0 + math.Atan(E/(lc.ρ0()-N))/lc.n()

	ρ := math.Sqrt(math.Pow(E, 2.0) + math.Pow(lc.ρ0()-N, 2.0))
	if lc.n() < 0.0 {
		ρ = -ρ
	}

	t := math.Pow(ρ/(lc.a*lc.f()), 1.0/lc.n())

	ϕ, ϕ0 := math.Pi/2.0-2.0*math.Atan(t), math.Inf(1)
	for math.Abs(ϕ-ϕ0) > 1.0e-09 {
		ϕ, ϕ0 = math.Pi/2.0-2.0*math.Atan(t*math.Pow((1.0-lc.e()*math.Sin(ϕ))/(1.0+lc.e()*math.Sin(ϕ)), lc.e()/2.0)), ϕ
	}

	return λ, ϕ
}

func NewZealandLambertConformal() LambertConformal {
	return LambertConformal{
		a:  6378137,
		f0: 1.0 / 298.257222101,
		ϕ1: math.Pi * -37.5 / 180.0,
		ϕ2: math.Pi * -44.5 / 180.0,
		ϕ0: math.Pi * -41.0 / 180.0,
		λ0: math.Pi * 173.0 / 180.0,
		n0: 7000000,
		e0: 3000000,
	}
}

type NZCS2000 struct {
	LambertConformal
}

func NewNZCS2000() NZCS2000 {

	lc := LambertConformal{}

	return NZCS2000{
		LambertConformal: lc,
	}
}

/*
func (n NZCS2000) Forward(x, y float64) (float64, float64) {
	return (n.LambertConformal.Forward(x, y))
}

func (n NZCS2000) Inverse(x, y float64) (float64, float64) {
	return (n.LambertConformal.Inverse(x, y))
}
*/
