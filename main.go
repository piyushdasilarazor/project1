package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	fmt.Println("This is the end")

	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}
}
