package cartogram

import ent "github.com/eyedeekay/game/entity"

type Tower struct {
	Name  string
	Tags  []string
	layer []*Layer
	Net   []*ent.Net
}

func NewTower(name string, layers int, tags ...string) (*Tower, error) {
	udp, err := ent.NewUDP()
	if err != nil {
		return nil, err
	}
	i2p, err := ent.NewI2P()
	if err != nil {
		return nil, err
	}
	newlayers := make([]*Layer, layers)
	for i := 0; i < layers; i++ {
		newlayers[i], err = NewLayer(name, tags...)
	}
	return &Tower{
		Name:  name,
		Tags:  tags,
		layer: newlayers,
		Net:   []*ent.Net{udp, i2p},
	}, nil
}
