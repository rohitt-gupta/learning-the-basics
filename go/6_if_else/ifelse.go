package main

import "fmt"

func main() {
	// age := 10

	// if age >= 18 {
	// 	fmt.Println("Person in an adult")
	// } else if age >= 12 {
	// 	fmt.Println("Person in a teenager")

	// } else {
	// 	fmt.Println("Person in a Kid")
	// }
	var role = "admin"
	var hasPermission = true

	if role == "admin" && hasPermission {
		fmt.Println("Has access and is admin")
	}

}
