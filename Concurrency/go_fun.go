package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, playground")
	urls := []string{"as", "bb", "CC"}

	for _, url := range urls {
		go func() {
			fmt.Println(url)
		}()
	}

	time.Sleep((1 * time.Second))

	numbers := []int{1, 2, 3, 4}
	for _, n := range numbers {
		go func() {
			fmt.Println(n)
		}()
		// time.Sleep((1 * time.Second)) // Enable this will correct the result :D
	}

	// Result:
	// 4
	// 4
	// 4
	// 4

	time.Sleep((4 * time.Second))
}
