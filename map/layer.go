package cartogram

import ent "github.com/eyedeekay/game/entity"

type Layer struct {
	Name string
	Tags []string
	Data []ent.Entity
}

func NewLayer(name string, tags ...string) (*Layer, error) {
	return &Layer{
		Name: name,
		Tags: tags,
		Data: make([]ent.Entity, 0),
	}, nil
}
