package main

type blueDragonBuilder struct {
	color  string
	breath string
}

func newBlueDragonBuilder() *blueDragonBuilder {
	return &blueDragonBuilder{}
}

func (rdb *blueDragonBuilder) setColor() {
	rdb.color = "Blue"
}

func (rdb *blueDragonBuilder) setBreath() {
	rdb.breath = "Ice"
}

func (rdb *blueDragonBuilder) getDragon() dragon {
	return dragon{
		color:  rdb.color,
		breath: rdb.breath,
	}
}
