package orbitelements

import "time"

const ()

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

//Planet ... todo
type Planet struct {
	Name             string
	OrbitElement     *OrbitElement
	OrbitDateElement *OrbitDateElement
	Solstice         *Solstice
	PlanetElement    *PlanetElement
}

//NewOrbitElement ... todo
func NewOrbitElement() *OrbitElement {
	return &OrbitElement{}
}
