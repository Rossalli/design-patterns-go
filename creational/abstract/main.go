package main

import "fmt"

func main() {
	blueDragonFactory, _ := getDragonsFactory("blue")

	redDragonFactory, _ := getDragonsFactory("red")

	blueAncient := blueDragonFactory.makeAncient()
	blueYoung := blueDragonFactory.makeYoung()

	fmt.Println("-----------------------Blue Dragon Factory-----------------------")
	fmt.Println("<><><><><>Ancient Dragon:<><><><><>")
	printDragonDetails(blueAncient)
	fmt.Println("<><><><><>Young Dragon:<><><><><>")
	printDragonDetails(blueYoung)
	fmt.Println("-----------------------------------------------------------------")

	redAncient := redDragonFactory.makeAncient()
	redYoung := redDragonFactory.makeYoung()

	fmt.Println("-----------------------Red Dragon Factory-----------------------")
	fmt.Println("<><><><><>Ancient Dragon:<><><><><>")
	printDragonDetails(redAncient)
	fmt.Println("<><><><><>Young Dragon:<><><><><>")
	printDragonDetails(redYoung)
	fmt.Println("-----------------------------------------------------------------")

}

func printDragonDetails(d iDragon) {
	fmt.Println("Color:", d.getColor())
	fmt.Println("Category:", d.getCategory())
	fmt.Println("Age: ", d.getAge())
}
