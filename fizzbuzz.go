package main
import (
	"fmt"

)

func main() {
	fmt.Println("Welcome to FizzBuzz")
		for i :=1; i<=100; i++ {
			if i %3 == 0 && i %5 == 0 {
				fmt.Println("FizzBuzz")
			} else if (i)%3 == 0 {
					fmt.Println("Fizz")
		
			} else if i %5 == 0 {
					fmt.Println("Buzz")
			} else {
				println(i)
			
	}
	}
}
