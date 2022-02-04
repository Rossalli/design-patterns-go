package main

import "fmt"

func main() {
	blueDragon, _ := getDragon("blue")
	redDragon, _ := getDragon("red")

	fmt.Println("-----------------------Blue Dragon-----------------------")
	printDragonDetails(blueDragon)

	fmt.Println("-----------------------Red Dragon-----------------------")
	printDragonDetails(redDragon)

}

func printDragonDetails(d iDragon) {
	fmt.Println("Color:", d.getColor())
	fmt.Println("Breath:", d.getBreath())
}
