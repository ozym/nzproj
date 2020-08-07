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

// NZTM2000 provides an implementation of the  New Zealand Transverse Mercator which satisfies the Projection interface.
func NZTM2000() transverseMercator {
	return NewTransverseMercator(TransverseMercatorParams{
		SemiMajorAxisOfReferenceEllipsoid: 6378137,
		FlatteningOfReferenceEllipsoid:    1.0 / 298.257222101,
		OriginLatitude:                    0.0,
		OriginLongitude:                   math.Pi * 173.0 / 180.0,
		FalseNorthingOfProjection:         10000000,
		FalseEastingOfProjection:          1600000,
		CentralMeridianScaleFactor:        0.9996,
	})
}
