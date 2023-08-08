package factorymethod

import "fmt"

func TestFactoryMethod() {
	ak47Factory := newAK47Factory()
	ak47, _ := ak47Factory.getGun()

	musketFactory := newMusketFactory()
	musket, _ := musketFactory.getGun()

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
