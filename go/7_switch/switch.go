package main

import "fmt"

func main() {

	// i := 4

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// case 4:
	// 	fmt.Println("Four")
	// 	i = 6

	// 	fmt.Println("i", i)
	// case 5:
	// 	fmt.Println("Five")
	// }

	// switch time.Now().Weekday() {
	// case time.Sunday, time.Saturday:
	// 	fmt.Println("Weekend, Let's Party")

	// default:
	// 	fmt.Println("Work hard it's weekday!")
	// }

	// switch time.Now() {
	// case time.Now():
	// 	fmt.Println(time.Now().UTC().Clock())

	// default:
	// 	fmt.Println("Work hard it's weekday!")
	// }

	getType := func(i interface{}) {
		switch t := i.(type) {
		case string:
			fmt.Println("It is a string", t)
		case int:
			fmt.Println("It is a Integer", t)
		case bool:
			fmt.Println("It is a Boolean", t)
		default:
			fmt.Println("others", t)
		}
	}

	getType("Hello")
}
