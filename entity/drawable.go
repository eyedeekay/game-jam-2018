package ent

type Drawable interface {
}

type Sprite struct {
}

var _ Drawable = &Sprite{}
