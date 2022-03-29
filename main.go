package main

import "fmt"

func main() {
	konton := map[string]int{}

	for {
		fmt.Println("1. Skapa konto")
		fmt.Println("2. Lista konton")
		fmt.Println("3. Avsluta")
		var sel string
		fmt.Scanln(&sel)
		if sel == "1" {

			fmt.Print("Ange kontonummer:")
			//beloppet kanske ni kan random
			//BORDE KOLLA IFALL KONTOT REDAN FINNS
			var kontonummer string
			fmt.Scanln(&kontonummer)
			konton[kontonummer] = 100
		}
		if sel == "2" {
			for key, value := range konton {
				fmt.Printf("%v : %v\n", key, value)
			}

		}
		if sel == "3" {
			break
		}
	}

	// players := map[int]string{}
	// players[13] = "Mats Sundin"
	// players[21] = "Peter Forsberg"

	// for key, value := range players {
	// 	fmt.Printf("%d : %s\n", key, value)
	// }

	//players := make(map[int]string, 20)

}
