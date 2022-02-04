package main

import "errors"

func getDragon(color string) (iDragon, error) {
	switch color {
	case "red":
		return newRedDragon(), nil
	case "blue":
		return newBlueDragon(), nil
	default:
		return nil, errors.New("unknown color")

	}
}
