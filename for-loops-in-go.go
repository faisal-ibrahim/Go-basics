package main

import "fmt"

func main() {
	fmt.Println("For loops in go")

	ages := []int{23, 34, 22, 44, 500}

	for i := 0; i < len(ages); i++ {
		fmt.Println(ages[i])
	}

	fmt.Println("Another way:")
	for i := 0; i < len(ages); {
		fmt.Println(ages[i])

		i++
	}

	fmt.Println("Using range in go")

	for index, value := range ages {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}
}
