package orbitelements

const ()

var ()

//PlanetElement ... todos
type PlanetElement struct {
	Mass         float64
	Radius       float64
	Gravity      float64
	Density      float64
	StarSystem   string
	OrbitsAround string
}

//NewPlanet ... todo
func NewPlanet(name string) *Planet {
	return &Planet{
		Name:             name,
		OrbitElement:     &OrbitElement{},
		OrbitDateElement: &OrbitDateElement{},
		Solstice:         &Solstice{},
	}
}
