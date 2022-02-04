package main

import (
	"errors"
)

type iDragonsFactory interface {
	makeYoung() iDragon
	makeAncient() iDragon
}

func getDragonsFactory(color string) (iDragonsFactory, error) {
	switch color {
	case "red":
		return &redDragon{}, nil
	case "blue":
		return &blueDragon{}, nil
	default:
		return nil, errors.New("unknown color")
	}
}
