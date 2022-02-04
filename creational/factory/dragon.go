package main

type dragon struct {
	color  string
	breath string
}

func (d *dragon) setColor(color string) {
	d.color = color
}

func (d *dragon) getColor() string {
	return d.color
}

func (d *dragon) setBreath(breath string) {
	d.breath = breath
}

func (d *dragon) getBreath() string {
	return d.breath
}
