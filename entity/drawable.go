package ent

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/EngoEngine/engo/common"
)

type Drawable interface {
	Name() string
	Tags() []string
	Size() (*Size, error)
	Texture() (*common.Texture, error)
}

type Sprite struct {
	tags     []string
	size     *Size
	template string
	texture  *common.Texture
}

var _ Drawable = &Sprite{}

func (s *Sprite) Name() string {
	return strings.Join(s.tags, "")
}

func (s *Sprite) Tags() []string {
	return s.tags
}

func (s *Sprite) Size() (*Size, error) {
	return s.size, nil
}

func (s *Sprite) Texture() (*common.Texture, error) {
	return s.texture, nil
}

func Load(file string) (Drawable, error) {
	var s Sprite
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	if s.tags == nil || s.size == nil || s.template == "" {
		return nil, fmt.Errorf("invalid sprite, tags, size, and template must not be nil")
	}
	return &s, nil
}
