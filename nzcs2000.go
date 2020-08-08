package nzproj

import (
	"math"
)

/*
	New Zealand Continental Shelf Lambert Conformal 2000

	The extents of NZCS2000 are 160° E to 170° W and 25° S to 60° S.

	A Lambert conformal conic projection was chosen rather than the Transverse Mercator projection
	because the latter becomes excessively distorted when extended over large longitudinal ranges.

	https://www.linz.govt.nz/data/geodetic-system/datums-projections-heights/projections/new-zealand-continental-shelf-lambert
*/

type NZCS2000 struct {
	LambertConformalConic
}

// NZCS2000 provides an implementation of the New Zealand Continental Shelf projection.
func NewNZCS2000() NZCS2000 {
	return NZCS2000{NewLambertConformalConic(LambertConformalConicParams{
		SemiMajorAxisOfReferenceEllipsoid:    6378137,
		OriginFlatteningOfReferenceEllipsoid: 1.0 / 298.257222101,
		LatitudeOfFirstStandardParallel:      math.Pi * -37.5 / 180.0,
		LatitudeOfSecondStandardParallel:     math.Pi * -44.5 / 180.0,
		OriginLatitude:                       math.Pi * -41.0 / 180.0,
		OriginLongitude:                      math.Pi * 173.0 / 180.0,
		FalseNorthingOfProjection:            7000000,
		FalseEastingOfProjection:             3000000,
	})}
}

func (cs NZCS2000) EPSG() int {
	return 3851
}

func (cs NZCS2000) Center() (float64, float64) {
	return 3142938.48, 7008029.86
}

func (cs NZCS2000) Bounds() (float64, float64, float64, float64) {
	return 2199991.09, 5263459.16, 4624385.49, 8545243.46
}

func (cs NZCS2000) WGS84() (float64, float64, float64, float64) {
	return 160.6, -55.95, -171.2, -25.88
}
