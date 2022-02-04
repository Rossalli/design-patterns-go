package main

type blueDragon struct {
	dragon
}

func newBlueDragon() iDragon {
	return &redDragon{
		dragon: dragon{
			color:  "Blue",
			breath: "Ice",
		},
	}
}
