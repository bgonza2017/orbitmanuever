package orbitelements

const ()

var ()

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

//NewOrbitElement ... todo
func NewOrbitElement() *OrbitElement {
	return &OrbitElement{}
}
