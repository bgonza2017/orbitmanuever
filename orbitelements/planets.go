package orbitelements

import "time"

const ()

var ()

//NewEarth ... todo
//https://en.wikipedia.org/wiki/Earth%27s_orbit
func NewEarth() *Planet {
	return &Planet{
		Name: "Earth",
		OrbitElement: &OrbitElement{
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
		},
		OrbitDateElement: &OrbitDateElement{
			PeriapsisDate: time.Date(2017, time.January, 1, 0, 0, 0, 0, nil),
			ApoapsisDate:  time.Date(2017, time.July, 3, 0, 0, 0, 0, nil),
		},
		Solstice: &Solstice{
			Top:    time.Date(2017, time.March, 21, 0, 0, 0, 0, nil),
			Bottom: time.Date(2017, time.September, 23, 0, 0, 0, 0, nil),
			Left:   time.Date(2017, time.June, 21, 0, 0, 0, 0, nil),
			Right:  time.Date(2017, time.December, 21, 0, 0, 0, 0, nil),
		},
	}

}
