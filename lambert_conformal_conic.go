package nzproj

import (
	"math"
)

/*
	The Lambert Conformal Conic projection is a projection in which geographic meridians are
	represented by straight lines which meet at the projection of the pole and geographic parallels
	are represented by a series of arcs of circles with this point as their centre.
*/

type LambertConformalConicParams struct {
	SemiMajorAxisOfReferenceEllipsoid    float64
	OriginFlatteningOfReferenceEllipsoid float64
	LatitudeOfFirstStandardParallel      float64
	LatitudeOfSecondStandardParallel     float64
	OriginLatitude                       float64
	OriginLongitude                      float64
	FalseNorthingOfProjection            float64
	FalseEastingOfProjection             float64
}

type lambertConformalConic struct {
	// projection parameters
	a  float64 // semi-major axis of reference ellipsoid
	f0 float64 // origin flattening of reference ellipsoid
	ϕ1 float64 // latitude of the first standard parallel
	ϕ2 float64 // latitude of the second standard parallel
	ϕ0 float64 // origin latitude
	λ0 float64 // origin longitude
	n0 float64 // false Northing of projection
	e0 float64 // false Easting of projection

	// derived parameters
	e  float64
	n  float64
	f  float64
	ρ0 float64
}

// NewLambertConformalConic provides an implementation of the Lambert Conformal Conic projection with the given parameters which satisfies the Projection interface.
func NewLambertConformalConic(params LambertConformalConicParams) lambertConformalConic {
	lc := lambertConformalConic{
		a:  params.SemiMajorAxisOfReferenceEllipsoid,
		f0: params.OriginFlatteningOfReferenceEllipsoid,
		ϕ1: params.LatitudeOfFirstStandardParallel,
		ϕ2: params.LatitudeOfSecondStandardParallel,
		ϕ0: params.OriginLatitude,
		λ0: params.OriginLongitude,
		n0: params.FalseNorthingOfProjection,
		e0: params.FalseEastingOfProjection,
	}
	lc.e = math.Sqrt(2.0*lc.f0 - math.Pow(lc.f0, 2.0))
	lc.n = (math.Log(lc.m(lc.ϕ1)) - math.Log(lc.m(lc.ϕ2))) /
		(math.Log(lc.t(lc.ϕ1)) - math.Log(lc.t(lc.ϕ2)))
	lc.f = lc.m(lc.ϕ1) / (lc.n * math.Pow(lc.t(lc.ϕ1), lc.n))
	lc.ρ0 = lc.ρ(lc.ϕ0)

	return lc
}

func (lc *lambertConformalConic) m(l float64) float64 {
	return math.Cos(l) / math.Sqrt(1.0-math.Pow(lc.e, 2.0)*math.Pow(math.Sin(l), 2.0))
}
func (lc *lambertConformalConic) t(l float64) float64 {
	return math.Tan((math.Pi/4.0)-(l/2.0)) / math.Pow((1.0-lc.e*math.Sin(l))/(1.0+lc.e*math.Sin(l)), lc.e/2.0)
}

func (lc *lambertConformalConic) ρ(l float64) float64 {
	return lc.a * lc.f * math.Pow(lc.t(l), lc.n)
}

func (lc lambertConformalConic) Forward(lon, lat float64) (float64, float64) {
	return lc.forward(math.Pi*lon/180.0, math.Pi*lat/180.0)
}

func (lc lambertConformalConic) forward(λ, ϕ float64) (float64, float64) {
	y := lc.n * (λ - lc.λ0)

	X := lc.e0 + lc.ρ(ϕ)*math.Sin(y)
	Y := lc.n0 + lc.ρ0 - lc.ρ(ϕ)*math.Cos(y)

	return X, Y
}

func (lc lambertConformalConic) Inverse(x, y float64) (float64, float64) {
	λ, ϕ := lc.inverse(x, y)

	return 180.0 * λ / math.Pi, 180.0 * ϕ / math.Pi
}

func (lc lambertConformalConic) inverse(x, y float64) (float64, float64) {

	E := x - lc.e0
	N := y - lc.n0

	λ := lc.λ0 + math.Atan(E/(lc.ρ0-N))/lc.n

	ρ := math.Sqrt(math.Pow(E, 2.0) + math.Pow(lc.ρ0-N, 2.0))
	if lc.n < 0.0 {
		ρ = -ρ
	}

	t := math.Pow(ρ/(lc.a*lc.f), 1.0/lc.n)

	ϕ, ϕ0 := math.Pi/2.0-2.0*math.Atan(t), math.Inf(1)
	for math.Abs(ϕ-ϕ0) > 1.0e-09 {
		ϕ, ϕ0 = math.Pi/2.0-2.0*math.Atan(t*math.Pow((1.0-lc.e*math.Sin(ϕ))/(1.0+lc.e*math.Sin(ϕ)), lc.e/2.0)), ϕ
	}

	return λ, ϕ
}
