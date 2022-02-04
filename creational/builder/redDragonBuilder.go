package main

type redDragonBuilder struct {
	color  string
	breath string
}

func newRedDragonBuilder() *redDragonBuilder {
	return &redDragonBuilder{}
}

func (rdb *redDragonBuilder) setColor() {
	rdb.color = "Red"
}

func (rdb *redDragonBuilder) setBreath() {
	rdb.breath = "Fire"
}

func (rdb *redDragonBuilder) getDragon() dragon {
	return dragon{
		color:  rdb.color,
		breath: rdb.breath,
	}
}
