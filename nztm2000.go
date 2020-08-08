package nzproj

import (
	"math"
)

/*
	New Zealand Transverse Mercator 2000 (NZTM2000) is the projection used for New Zealand's
	Topo50 1:50,000 and other small scale mapping. Spatial data users are encouraged to use
	NZTM2000 where a projection is required within mainland New Zealand.

	https://www.linz.govt.nz/data/geodetic-system/datums-projections-and-heights/projections/new-zealand-transverse-mercator-2000
*/

// NZTM2000 is an implementation of the New Zealand Transverse Mercator projection.
type NZTM2000 struct {
	TransverseMercator
}

// NewNZTM2000 returns an implementation of the New Zealand Transverse Mercator.
func NewNZTM2000() NZTM2000 {
	return NZTM2000{NewTransverseMercator(TransverseMercatorParams{
		SemiMajorAxisOfReferenceEllipsoid: 6378137,
		FlatteningOfReferenceEllipsoid:    1.0 / 298.257222101,
		OriginLatitude:                    0.0,
		OriginLongitude:                   math.Pi * 173.0 / 180.0,
		FalseNorthingOfProjection:         10000000,
		FalseEastingOfProjection:          1600000,
		CentralMeridianScaleFactor:        0.9996,
	})}
}

func (tm NZTM2000) EPSG() int {
	return 2193
}

func (tm NZTM2000) Center() (float64, float64) {
	return 1743159.15, 5469287.49
}

func (tm NZTM2000) Bounds() (float64, float64, float64, float64) {
	return 827933.23, 3729820.29, 3195373.59, 7039943.58
}

func (tm NZTM2000) WGS84() (float64, float64, float64, float64) {
	return 160.6, -55.95, -171.2, -25.88
}
