package main

import "fmt"

func main() {
	var queue []string

	queue = append(queue, "Hello ")
	queue = append(queue, "world!")

	for len(queue) > 0 {
		fmt.Print(queue[0])
		queue = queue[1:]
	}
}
