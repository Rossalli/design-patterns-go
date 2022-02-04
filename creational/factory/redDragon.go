package main

type redDragon struct {
	dragon
}

func newRedDragon() iDragon {
	return &redDragon{
		dragon: dragon{
			color:  "Red",
			breath: "Fire",
		},
	}
}
