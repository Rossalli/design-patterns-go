package main

type iDragon interface {
	setColor(color string)
	getColor() string
	setCategory(category string)
	getCategory() string
	setSize(size string)
	getSize() string
	setAge(age string)
	getAge() string
}

type dragon struct {
	color    string
	category string
	size     string
	age      string
}

func (d *dragon) setColor(color string) {
	d.color = color
}

func (d *dragon) getColor() string {
	return d.color
}

func (d *dragon) setCategory(category string) {
	d.category = category
}

func (d *dragon) getCategory() string {
	return d.category
}

func (d *dragon) setSize(size string) {
	d.size = size
}

func (d *dragon) getSize() string {
	return d.size
}

func (d *dragon) setAge(age string) {
	d.age = age
}

func (d *dragon) getAge() string {
	return d.age
}
