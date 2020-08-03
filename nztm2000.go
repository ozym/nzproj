package nzproj

import (
	"math"
	"sync"
)

/*
	New Zealand Transverse Mercator 2000 (NZTM2000) is the projection used for New Zealand's
	Topo50 1:50,000 and other small scale mapping. Spatial data users are encouraged to use
	NZTM2000 where a projection is required within mainland New Zealand.

	https://www.linz.govt.nz/data/geodetic-system/datums-projections-and-heights/projections/new-zealand-transverse-mercator-2000

*/

type TransverseMercator struct {
	a  float64
	f  float64
	ϕ0 float64
	λ0 float64
	N0 float64
	E0 float64
	k0 float64

	_b, _e, _e2, _e4, _e6, _A0, _A2, _A4, _A6, _m0 struct {
		once  sync.Once
		value float64
	}
}

func (tm *TransverseMercator) b() float64 {
	tm._b.once.Do(func() {
		tm._b.value = tm.a * (1.0 - tm.f)
	})
	return tm._b.value
}

func (tm *TransverseMercator) e() float64 {
	tm._e.once.Do(func() {
		tm._e.value = math.Sqrt(tm.e2())
	})
	return tm._e.value
}

func (tm *TransverseMercator) e2() float64 {
	tm._e2.once.Do(func() {
		tm._e2.value = 2.0*tm.f - math.Pow(tm.f, 2.0)
	})
	return tm._e2.value
}

func (tm *TransverseMercator) e4() float64 {
	tm._e4.once.Do(func() {
		tm._e4.value = math.Pow(tm.e(), 4.0)
	})
	return tm._e4.value
}

func (tm *TransverseMercator) e6() float64 {
	tm._e6.once.Do(func() {
		tm._e6.value = math.Pow(tm.e(), 6.0)
	})
	return tm._e6.value
}

func (tm *TransverseMercator) a0() float64 {
	tm._A0.once.Do(func() {
		tm._A0.value = 1.0 - (tm.e2() / 4.0) - (3.0 * tm.e4() / 64.0) - (5.0 * tm.e6() / 256.0)
	})
	return tm._A0.value
}

func (tm *TransverseMercator) a2() float64 {
	tm._A2.once.Do(func() {
		tm._A2.value = 3.0 * (tm.e2() + tm.e4()/4.0 + 15.0*tm.e6()/128.0) / 8.0
	})
	return tm._A2.value
}

func (tm *TransverseMercator) a4() float64 {
	tm._A4.once.Do(func() {
		tm._A4.value = 15.0 * (tm.e4() + 3.0*tm.e6()/4.0) / 256.0
	})
	return tm._A4.value
}

func (tm *TransverseMercator) a6() float64 {
	tm._A6.once.Do(func() {
		tm._A6.value = 35.0 * tm.e6() / 3072.0
	})
	return tm._A6.value
}

func (tm *TransverseMercator) m(ϕ float64) float64 {

	sin2ϕ := math.Sin(2.0 * ϕ)
	sin4ϕ := math.Sin(4.0 * ϕ)
	sin6ϕ := math.Sin(6.0 * ϕ)

	return tm.a * (tm.a0()*ϕ - tm.a2()*sin2ϕ + tm.a4()*sin4ϕ - tm.a6()*sin6ϕ)
}

func (tm *TransverseMercator) m0() float64 {
	tm._m0.once.Do(func() {
		tm._m0.value = tm.m(tm.ϕ0)
	})
	return tm._m0.value
}

func (tm *TransverseMercator) ρ(ϕ float64) float64 {

	sinϕ2 := math.Pow(math.Sin(ϕ), 2.0)

	return tm.a * (1.0 - tm.e2()) / math.Pow(1.0-tm.e2()*sinϕ2, 3/2)
}

func (tm *TransverseMercator) ν(ϕ float64) float64 {
	sinϕ2 := math.Pow(math.Sin(ϕ), 2.0)

	return tm.a / math.Sqrt(1.0-tm.e2()*sinϕ2)
}

func (tm *TransverseMercator) ψ(ϕ float64) float64 {
	return tm.ν(ϕ) / tm.ρ(ϕ)
}

func (tm TransverseMercator) Forward(lon, lat float64) (float64, float64) {
	return tm.forward(math.Pi*lon/180.0, math.Pi*lat/180.0)
}

func (tm TransverseMercator) forward(λ, ϕ float64) (float64, float64) {

	t := math.Tan(ϕ)
	t2 := math.Pow(t, 2.0)
	t4 := math.Pow(t, 4.0)
	t6 := math.Pow(t, 6.0)

	φ := λ - tm.λ0
	φ2 := math.Pow(φ, 2.0)
	φ4 := math.Pow(φ, 4.0)
	φ6 := math.Pow(φ, 6.0)
	φ8 := math.Pow(φ, 8.0)

	ψ := tm.ψ(ϕ)
	ψ2 := math.Pow(ψ, 2.0)
	ψ3 := math.Pow(ψ, 3.0)
	ψ4 := math.Pow(ψ, 4.0)

	ν := tm.ν(ϕ)

	sinϕ := math.Sin(ϕ)
	cosϕ := math.Cos(ϕ)

	cosϕ2 := math.Pow(cosϕ, 2.0)
	cosϕ3 := math.Pow(cosϕ, 3.0)
	cosϕ4 := math.Pow(cosϕ, 4.0)
	cosϕ5 := math.Pow(cosϕ, 5.0)
	cosϕ6 := math.Pow(cosϕ, 6.0)
	cosϕ7 := math.Pow(cosϕ, 7.0)

	nT1 := φ2 * ν * sinϕ * cosϕ / 2.0
	nT2 := φ4 * ν * sinϕ * cosϕ3 * (4.0*ψ2 + ψ - t2) / 24.0
	nT3 := φ6 * ν * sinϕ * cosϕ5 * (8.0*ψ4*(11.0-24.0*t2) - 28.0*ψ3*(1.0-6.0*t2) + ψ2*(1.0-32.0*t2) - 2.0*ψ*t2 + t4) / 720.0
	nT4 := φ8 * ν * sinϕ * cosϕ7 * (1385.0 - 3111.0*t2 + 543.0*t4 - t6) / 40320.0

	N := tm.N0 + tm.k0*(tm.m(ϕ)-tm.m0()+nT1+nT2+nT3+nT4)

	eT1 := φ2 * cosϕ2 * (ψ - t2) / 6.0
	eT2 := φ4 * cosϕ4 * (4.0*ψ3*(1.0-6.0*t2) + ψ2*(1.0+8.0*t2) - 2.0*ψ*t2 + t4) / 120.0
	eT3 := φ6 * cosϕ6 * (61.0 - 479.0*t2 + 179.0*t4 - t6) / 5040.0

	E := tm.E0 + tm.k0*ν*φ*cosϕ*(1.0+eT1+eT2+eT3)

	return E, N
}

func (tm TransverseMercator) Inverse(x, y float64) (float64, float64) {
	λ, ϕ := tm.inverse(x, y)

	return 180.0 * λ / math.Pi, 180.0 * ϕ / math.Pi
}

func (tm TransverseMercator) inverse(x, y float64) (float64, float64) {

	N := y - tm.N0

	m := tm.m0() + N/tm.k0

	n := (tm.a - tm.b()) / (tm.a + tm.b())
	n2 := math.Pow(n, 2.0)
	n3 := math.Pow(n, 3.0)
	n4 := math.Pow(n, 4.0)

	G := tm.a * (1.0 - n) * (1.0 - n2) * (1.0 + 9.0*n2/4.0 + 225.0*n4/64.0) * math.Pi / 180.0

	σ := m * math.Pi / (180.0 * G)
	sin2σ := math.Sin(2.0 * σ)
	sin4σ := math.Sin(4.0 * σ)
	sin6σ := math.Sin(6.0 * σ)
	sin8σ := math.Sin(8.0 * σ)

	ϕd := σ + (3.0*n/2.0-27.0*n3/32.0)*sin2σ + (21.0*n2/16.0-55.0*n4/32.0)*sin4σ + (151.0*n3/96.0)*sin6σ + (1097.0*n4/512.0)*sin8σ
	sinϕd := math.Sin(ϕd)
	sinϕd2 := math.Pow(sinϕd, 2.0)

	ρd := tm.a * (1.0 - tm.e2()) / math.Pow(1.0-tm.e2()*sinϕd2, 3.0/2.0)
	νd := tm.a / math.Sqrt(1.0-tm.e2()*sinϕd2)

	t := math.Tan(ϕd)
	t2 := math.Pow(t, 2.0)
	t4 := math.Pow(t, 4.0)
	t6 := math.Pow(t, 6.0)

	ψ := νd / ρd
	ψ2 := math.Pow(ψ, 2.0)
	ψ3 := math.Pow(ψ, 3.0)
	ψ4 := math.Pow(ψ, 4.0)

	E := x - tm.E0

	x1 := E / (tm.k0 * νd)
	x3 := math.Pow(x1, 3.0)
	x5 := math.Pow(x1, 5.0)
	x7 := math.Pow(x1, 7.0)

	secϕd := 1.0 / math.Cos(ϕd)

	ϕT1 := t * E * x1 / (2.0 * tm.k0 * ρd)
	ϕT2 := t * E * x3 * (-4.0*ψ2 + 9.0*ψ*(1-t2) + 12.0*t2) / (24.0 * tm.k0 * ρd)
	ϕT3 := t * E * x5 * (8.0*ψ4*(11.0-24.0*t2) - 12.0*ψ3*(21.0-71.0*t2) + 15.0*ψ2*(15.0-98.0*t2+15*t4) + 180.0*ψ*(5.0*t2-3.0*t4) + 360.0*t4) / (720.0 * tm.k0 * ρd)
	ϕT4 := t * E * x7 * (1385.0 + 3633.0*t2 + 4095.0*t4 + 1575.0*t6) / (40320.0 * tm.k0 * ρd)

	ϕ := ϕd - ϕT1 + ϕT2 - ϕT3 + ϕT4

	λT1 := x1 * secϕd
	λT2 := x3 * secϕd * (ψ + 2.0*t2) / 6.0
	λT3 := x5 * secϕd * (-4.0*ψ3*(1.0-6.0*t2) + ψ2*(9.0-68.0*t2) + 72.0*ψ*t2 + 24*t4) / 120.0
	λT4 := x7 * secϕd * (61.0 + 662.0*t2 + 1320.0*t4 + 720*t6) / 5040.0

	λ := tm.λ0 + λT1 - λT2 + λT3 - λT4

	return λ, ϕ
}

func NewZealandTransverseMercator() TransverseMercator {
	return TransverseMercator{
		a:  6378137,
		f:  1.0 / 298.257222101,
		ϕ0: 0.0,
		λ0: math.Pi * 173.0 / 180.0,
		N0: 10000000,
		E0: 1600000,
		k0: 0.9996,
	}
}

type NZTM2000 struct {
	TransverseMercator
}

func NewNZTM2000() NZTM2000 {

	tm := NewZealandTransverseMercator()

	return NZTM2000{
		TransverseMercator: tm,
	}
}
