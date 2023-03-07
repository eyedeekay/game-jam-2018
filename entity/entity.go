package ent

import (
	"net"
	"path/filepath"

	"log"

	"github.com/eyedeekay/i2pkeys"
)

type Interactable interface {
}

type Entity struct {
	Name        string
	Description string
	Body        Drawable
	Location    *Location
	Networks    []*Net
	Peers       []*Peer
	/*
		Inventory   *Inventory
		Stats       *Stats
		Actions     *Actions
		Effects     *Effects
		Equipment   *Equipment
		Skills      *Skills
		Items       *Items
		Quests      *Quests
		Spells      *Spells
		Class       *Class
	*/
}

func NewEntity(name, description, state string) (*Entity, error) {
	udp, err := NewUDP()
	if err != nil {
		return nil, err
	}
	i2p, err := NewI2P()
	if err != nil {
		return nil, err
	}
	spriteData := filepath.Join("data/", state, "/sprite.json")
	body, err := Load(spriteData)
	if err != nil {
		return nil, err
	}
	locationData := filepath.Join("data/", state, "/location.json")
	location, err := LoadLocation(locationData)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &Entity{
		Name:        name,
		Description: description,
		Networks:    []*Net{udp, i2p},
		Body:        body,
		Location:    location,
	}, nil
}

func (e *Entity) Send(message string) (int, error) {
	var sent int
	for _, peer := range e.Peers {
		go func(peer *Peer) {
			for _, peeraddr := range peer.Addresses {
				switch peeraddr.(type) {
				case *net.UDPAddr:
					s, err := e.Networks[0].Send(message, peeraddr)
					if err != nil {
						log.Println(err)
					}
					sent += s
				case *i2pkeys.I2PAddr:
					s, err := e.Networks[1].Send(message, peeraddr)
					if err != nil {
						log.Println(err)
					}
					sent += s
				}
			}
		}(peer)
	}
	return sent, nil
}

func (e *Entity) Receive() ([]Message, error) {
	var messages []Message
	for _, net := range e.Networks {
		message, err := net.Receive()
		messages = append(messages, message)
		if err != nil {
			return messages, err
		}
	}
	return messages, nil
}
