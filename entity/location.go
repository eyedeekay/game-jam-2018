package ent

import (
	"encoding/json"
	"io/ioutil"
)

type Size struct {
	Width  int
	Height int
}

type Location struct {
	X int
	Y int
	//	Z int
}

func LoadLocation(file string) (*Location, error) {
	var l Location
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}
