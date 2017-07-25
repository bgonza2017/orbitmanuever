package types

import (
	"fmt"
	"testing"
	"time"
)

// Create service
func TestPlanet(t *testing.T) {

	earth := &CelestialBody{
		Name:           "Earth",
		Category:       Planet,
		Radius:         float64(3959),     // mi
		Mass:           float64(5.972E24), // × 10^24 kg
		Gravity:        float64(9.807),    // m/s²
		Density:        float64(5.51),     // g/cm³
		EscapeVelocity: float64(11.19),
		StarSystem:     "Sol",
		OrbitsAround:   "Sol",
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
			PeriapsisDate: time.Date(2017, time.January, 1, 0, 0, 0, 0, &time.Location{}),
			ApoapsisDate:  time.Date(2017, time.July, 3, 0, 0, 0, 0, &time.Location{}),
		},
		Solstice: &Solstice{
			Top:    time.Date(2017, time.March, 21, 0, 0, 0, 0, &time.Location{}),
			Bottom: time.Date(2017, time.September, 23, 0, 0, 0, 0, &time.Location{}),
			Left:   time.Date(2017, time.June, 21, 0, 0, 0, 0, &time.Location{}),
			Right:  time.Date(2017, time.December, 21, 0, 0, 0, 0, &time.Location{}),
		},
	}
	fmt.Printf("earth:%v\n", earth)

	//
	// Mars Data
	mars := &CelestialBody{
		Name:     "Mars",
		Category: Planet,
		Radius:   float64(2106.00), // mi
		Mass:     float64(6.39),    // × 10^23 kg
		Gravity:  float64(3.711),   // m/s²
		//Density:      float64(0.00),  // g/cm³
		EscapeVelocity: float64(5.03),
		StarSystem:     "Sol",
		OrbitsAround:   "Sol",
		OrbitElement: &OrbitElement{
			SemiMajorAxis:            float64(227.92),
			Apoapsis:                 float64(249.23),
			Periapsis:                float64(206.62),
			Eccentricity:             float64(0.09341233),
			Inclination:              float64(1.85061),
			LongitudeOfAscendingNode: float64(49.57854),
			ArgumentOfPeriapsis:      float64(288.1),
			Period:                   float64(686.973),
			AverageSpeed:             float64(24.1),
			LongitudeOfPerihelion:    float64(336.04084),
		},
		Satellites: []*CelestialBody{
			&CelestialBody{
				Name:         "Phobos",
				Category:     Satellite,
				Radius:       float64(9.1),  //km
				Density:      float64(1900), // kg/cm³
				StarSystem:   "Sol",
				OrbitsAround: "Mars",
				OrbitElement: &OrbitElement{
					SemiMajorAxis: float64(9378), //km
					Eccentricity:  float64(0.0151),
					Inclination:   float64(1.08),
					Period:        float64(0.31891), //days
				},
			},
			&CelestialBody{
				Name:         "Deimos",
				Category:     Satellite,
				Radius:       float64(5.1),  //km
				Density:      float64(1750), // kg/cm³
				StarSystem:   "Sol",
				OrbitsAround: "Mars",
				OrbitElement: &OrbitElement{
					SemiMajorAxis: float64(23459), //km
					Eccentricity:  float64(0.0005),
					Inclination:   float64(1.79),
					Period:        float64(1.26244), //days
				},
			},
		},
		/*		OrbitDateElement: &OrbitDateElement{
					PeriapsisDate: time.Date(2017, time.January, 1, 0, 0, 0, 0, &time.Location{}),
					ApoapsisDate:  time.Date(2017, time.July, 3, 0, 0, 0, 0, &time.Location{}),
				},
				Solstice: &Solstice{
					Top:    time.Date(2017, time.March, 21, 0, 0, 0, 0, &time.Location{}),
					Bottom: time.Date(2017, time.September, 23, 0, 0, 0, 0, &time.Location{}),
					Left:   time.Date(2017, time.June, 21, 0, 0, 0, 0, &time.Location{}),
					Right:  time.Date(2017, time.December, 21, 0, 0, 0, 0, &time.Location{}),
				},

		*/
	}
	fmt.Printf("mars:%v\n", mars)

}
