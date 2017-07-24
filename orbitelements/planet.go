package orbitelements

const ()

var ()

//NewPlanet ... todo
func NewPlanet(name string) *Planet {
	return &Planet{
		Name:             name,
		OrbitElement:     &OrbitElement{},
		OrbitDateElement: &OrbitDateElement{},
		Solstice:         &Solstice{},
	}
}
