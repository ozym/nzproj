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

// NZCS2000 provides an implementation of the New Zealand Continental Shelf projection which satisfies the Projection interface.
func NZCS2000() lambertConformal {
	return NewLambertConformal(LambertConformalParams{
		SemiMajorAxisOfReferenceEllipsoid:    6378137,
		OriginFlatteningOfReferenceEllipsoid: 1.0 / 298.257222101,
		LatitudeOfFirstStandardParallel:      math.Pi * -37.5 / 180.0,
		LatitudeOfSecondStandardParallel:     math.Pi * -44.5 / 180.0,
		OriginLatitude:                       math.Pi * -41.0 / 180.0,
		OriginLongitude:                      math.Pi * 173.0 / 180.0,
		FalseNorthingOfProjection:            7000000,
		FalseEastingOfProjection:             3000000,
	})
}
