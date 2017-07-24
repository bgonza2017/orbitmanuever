package orbitelements

const ()

var ()

//NewEarth ... todo
func NewEarth() *OrbitElement {
	return &OrbitElement{
		SemiMajorAxis:            float64(92.96),
		Apoapsis:                 float64(94.51),
		Periapsis:                float64(91.40),
		Eccentricity:             float64(0.0167086),
		Inclination:              float64(7.155),
		LongitudeOfAscendingNode: float64(174.9),
		ArgumentOfPeriapsis:      float64(288.1),
		Period:                   float64(365.256),
		AverageSpeed:             float64(66600.00),
		LongitudeOfPerihelion:    float64(102.9),
	}

}
