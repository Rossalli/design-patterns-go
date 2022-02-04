package main

import "errors"

type iDragonBuilder interface {
	setColor()
	setBreath()
	getDragon() dragon
}

func getDragonsBuilder(color string) (iDragonBuilder, error) {
	switch color {
	case "red":
		return &redDragonBuilder{}, nil
	case "blue":
		return &blueDragonBuilder{}, nil
	default:
		return nil, errors.New("unknown color")
	}
}
