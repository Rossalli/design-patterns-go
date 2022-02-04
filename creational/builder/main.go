package main

import "fmt"

func main() {
	blueDragonBuilder, _ := getDragonsBuilder("blue")
	redDragonBuilder, _ := getDragonsBuilder("red")

	fmt.Println("-----------------------Blue Dragon Lair-----------------------")
	lair := newLair(blueDragonBuilder)
	blueDragon := lair.buildDragon()
	printDragonDetails(blueDragon)

	fmt.Println("-----------------------Red Dragon Lair-----------------------")
	lair.setDragonBuilder(redDragonBuilder)
	redDragon := lair.buildDragon()
	printDragonDetails(redDragon)

}

func printDragonDetails(d dragon) {
	fmt.Println("Color:", d.color)
	fmt.Println("Breath:", d.breath)
}
