

# nzproj
`import "github.com/ozym/nzproj"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [func NewLambertConformalConic(params LambertConformalConicParams) lambertConformalConic](#NewLambertConformalConic)
* [func NewTransverseMercator(params TransverseMercatorParams) transverseMercator](#NewTransverseMercator)
* [type LambertConformalConicParams](#LambertConformalConicParams)
* [type NZCS2000](#NZCS2000)
  * [func NewNZCS2000() NZCS2000](#NewNZCS2000)
  * [func (cs NZCS2000) Bounds() (float64, float64, float64, float64)](#NZCS2000.Bounds)
  * [func (cs NZCS2000) Center() (float64, float64)](#NZCS2000.Center)
  * [func (cs NZCS2000) EPSG() int](#NZCS2000.EPSG)
  * [func (lc NZCS2000) Forward(lon, lat float64) (float64, float64)](#NZCS2000.Forward)
  * [func (lc NZCS2000) Inverse(x, y float64) (float64, float64)](#NZCS2000.Inverse)
  * [func (cs NZCS2000) WGS84() (float64, float64, float64, float64)](#NZCS2000.WGS84)
* [type NZTM2000](#NZTM2000)
  * [func NewNZTM2000() NZTM2000](#NewNZTM2000)
  * [func (tm NZTM2000) Bounds() (float64, float64, float64, float64)](#NZTM2000.Bounds)
  * [func (tm NZTM2000) Center() (float64, float64)](#NZTM2000.Center)
  * [func (tm NZTM2000) EPSG() int](#NZTM2000.EPSG)
  * [func (tm NZTM2000) Forward(lon, lat float64) (float64, float64)](#NZTM2000.Forward)
  * [func (tm NZTM2000) Inverse(x, y float64) (float64, float64)](#NZTM2000.Inverse)
  * [func (tm NZTM2000) WGS84() (float64, float64, float64, float64)](#NZTM2000.WGS84)
* [type TransverseMercatorParams](#TransverseMercatorParams)


#### <a name="pkg-files">Package files</a>
[lambert_conformal_conic.go](/src/target/lambert_conformal_conic.go) [nzcs2000.go](/src/target/nzcs2000.go) [nztm2000.go](/src/target/nztm2000.go) [transverse_mercator.go](/src/target/transverse_mercator.go) 





## <a name="NewLambertConformalConic">func</a> [NewLambertConformalConic](/src/target/lambert_conformal_conic.go?s=1466:1553#L44)
``` go
func NewLambertConformalConic(params LambertConformalConicParams) lambertConformalConic
```
NewLambertConformalConic provides an implementation of the Lambert Conformal Conic projection with the given parameters.



## <a name="NewTransverseMercator">func</a> [NewTransverseMercator](/src/target/transverse_mercator.go?s=1367:1445#L48)
``` go
func NewTransverseMercator(params TransverseMercatorParams) transverseMercator
```
NewTransverseMercator provides an implementation of the Transverse Mercator with the given parameters.




## <a name="LambertConformalConicParams">type</a> [LambertConformalConicParams](/src/target/lambert_conformal_conic.go?s=416:827#L14)
``` go
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

```
LambertConformalConicParams holds the projection parameters for a Lambert Conformal Conic projection.










## <a name="NZCS2000">type</a> [NZCS2000](/src/target/nzcs2000.go?s=486:533#L18)
``` go
type NZCS2000 struct {
    // contains filtered or unexported fields
}

```






### <a name="NewNZCS2000">func</a> [NewNZCS2000](/src/target/nzcs2000.go?s=623:650#L23)
``` go
func NewNZCS2000() NZCS2000
```
NZCS2000 provides an implementation of the New Zealand Continental Shelf projection.





### <a name="NZCS2000.Bounds">func</a> (NZCS2000) [Bounds](/src/target/nzcs2000.go?s=1331:1395#L44)
``` go
func (cs NZCS2000) Bounds() (float64, float64, float64, float64)
```



### <a name="NZCS2000.Center">func</a> (NZCS2000) [Center](/src/target/nzcs2000.go?s=1248:1294#L40)
``` go
func (cs NZCS2000) Center() (float64, float64)
```



### <a name="NZCS2000.EPSG">func</a> (NZCS2000) [EPSG](/src/target/nzcs2000.go?s=1200:1229#L36)
``` go
func (cs NZCS2000) EPSG() int
```



### <a name="NZCS2000.Forward">func</a> (NZCS2000) [Forward](/src/target/lambert_conformal_conic.go?s=2602:2678#L75)
``` go
func (lc NZCS2000) Forward(lon, lat float64) (float64, float64)
```



### <a name="NZCS2000.Inverse">func</a> (NZCS2000) [Inverse](/src/target/lambert_conformal_conic.go?s=2944:3016#L88)
``` go
func (lc NZCS2000) Inverse(x, y float64) (float64, float64)
```



### <a name="NZCS2000.WGS84">func</a> (NZCS2000) [WGS84](/src/target/nzcs2000.go?s=1456:1519#L48)
``` go
func (cs NZCS2000) WGS84() (float64, float64, float64, float64)
```



## <a name="NZTM2000">type</a> [NZTM2000](/src/target/nztm2000.go?s=504:548#L16)
``` go
type NZTM2000 struct {
    // contains filtered or unexported fields
}

```
NZTM2000 is an implementation of the New Zealand Transverse Mercator projection.







### <a name="NewNZTM2000">func</a> [NewNZTM2000](/src/target/nztm2000.go?s=631:658#L21)
``` go
func NewNZTM2000() NZTM2000
```
NewNZTM2000 returns an implementation of the New Zealand Transverse Mercator.





### <a name="NZTM2000.Bounds">func</a> (NZTM2000) [Bounds](/src/target/nztm2000.go?s=1211:1275#L41)
``` go
func (tm NZTM2000) Bounds() (float64, float64, float64, float64)
```



### <a name="NZTM2000.Center">func</a> (NZTM2000) [Center](/src/target/nztm2000.go?s=1128:1174#L37)
``` go
func (tm NZTM2000) Center() (float64, float64)
```



### <a name="NZTM2000.EPSG">func</a> (NZTM2000) [EPSG](/src/target/nztm2000.go?s=1080:1109#L33)
``` go
func (tm NZTM2000) EPSG() int
```



### <a name="NZTM2000.Forward">func</a> (NZTM2000) [Forward](/src/target/transverse_mercator.go?s=2881:2954#L105)
``` go
func (tm NZTM2000) Forward(lon, lat float64) (float64, float64)
```



### <a name="NZTM2000.Inverse">func</a> (NZTM2000) [Inverse](/src/target/transverse_mercator.go?s=4297:4366#L155)
``` go
func (tm NZTM2000) Inverse(x, y float64) (float64, float64)
```



### <a name="NZTM2000.WGS84">func</a> (NZTM2000) [WGS84](/src/target/nztm2000.go?s=1335:1398#L45)
``` go
func (tm NZTM2000) WGS84() (float64, float64, float64, float64)
```



## <a name="TransverseMercatorParams">type</a> [TransverseMercatorParams](/src/target/transverse_mercator.go?s=315:656#L13)
``` go
type TransverseMercatorParams struct {
    SemiMajorAxisOfReferenceEllipsoid float64
    FlatteningOfReferenceEllipsoid    float64
    OriginLatitude                    float64
    OriginLongitude                   float64
    FalseNorthingOfProjection         float64
    FalseEastingOfProjection          float64
    CentralMeridianScaleFactor        float64
}

```
TransverseMercatorParams describes a Transverse Mercator projection.














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
