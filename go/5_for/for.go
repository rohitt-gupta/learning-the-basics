package main

import "fmt"

func main() {
	// while loop
	// i := 1
	// for i <= 4 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// infinite loop
	// for {
	// 	fmt.Println(i)
	// }

	// for loop
	for i := 1; i <= 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
}
