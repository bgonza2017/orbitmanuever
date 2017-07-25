package types

import "time"

//CelestialBodyType ...
type CelestialBodyType int

const (
	//Planet ...
	Planet CelestialBodyType = 1 + iota
	//Sun ...
	Sun
	//Moon ...
	Moon
	//Asteroid ...
	Asteroid
	//Satellite ...
	Satellite

	//AUmeter Astronomical Unit Meters
	AUmeter float64 = 149597870700

	//AUmile Astronomical Unit Miles
	AUmile float64 = 92955807
)

var ()

//OrbitDateElement ... todo
type OrbitDateElement struct {
	PeriapsisDate time.Time
	ApoapsisDate  time.Time
}

//Solstice ... todo
type Solstice struct {
	Top    time.Time
	Bottom time.Time
	Left   time.Time
	Right  time.Time
}

//OrbitElement ... todo
type OrbitElement struct {
	SemiMajorAxis            float64
	Eccentricity             float64
	Inclination              float64
	ArgumentOfPeriapsis      float64
	TimeOfPeriapsisPassage   float64
	Periapsis                float64
	Apoapsis                 float64
	LongitudeOfAscendingNode float64
	LongitudeOfPerihelion    float64
	AscendingNode            float64
	DescendingNode           float64
	Period                   float64
	AverageSpeed             float64
}

//CelestialBody ... todo
type CelestialBody struct {
	Name             string
	Category         CelestialBodyType
	EscapeVelocity   float64
	Mass             float64
	Radius           float64
	Gravity          float64
	Density          float64
	StarSystem       string
	OrbitsAround     string
	OrbitElement     *OrbitElement
	OrbitDateElement *OrbitDateElement
	Solstice         *Solstice
	Satellites       []*CelestialBody
}

//StarSystem ... todo
type StarSystem struct {
	Name   string
	Bodies []*CelestialBody
}
