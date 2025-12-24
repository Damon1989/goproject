package main

import "fmt"

func main() {
	/*	var cheeses [2]string
		cheeses[0] = "Mu"
		cheeses[1] = "Cheddar"
		fmt.Println(cheeses[0], cheeses[1])
		fmt.Println(len(cheeses))
		fmt.Println(cheeses)*/

	/*var cheeses = make([]string, 2)
	cheeses[0] = "Mu1"
	cheeses[1] = "Cheddar1"
	fmt.Println(cheeses[0], cheeses[1])
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
	cheeses = append(cheeses, "Cheddar2", "Cheddar3", "Cheddar4")
	fmt.Println(cheeses)
	fmt.Println(len(cheeses))
	cheeses = append(cheeses[:2], cheeses[2+1:]...)
	fmt.Println(cheeses)
	fmt.Println(len(cheeses))*/

	/*	var cheeses = make([]string, 3)
		cheeses[0] = "Mu1"
		cheeses[1] = "Cheddar1"
		cheeses[2] = "Gouda1"
		var smellCheeses = make([]string, 2)
		copy(smellCheeses, cheeses)
		fmt.Println(smellCheeses)
		fmt.Println(len(smellCheeses))
		fmt.Println(cheeses)
		fmt.Println(len(cheeses))*/

	var players = make(map[string]int)
	players["Alice"] = 1
	players["Bob"] = 2
	players["Charlie"] = 3
	fmt.Println(players)
	fmt.Println(len(players))
	delete(players, "Bob")
	fmt.Println(players)
	fmt.Println(len(players))
}
