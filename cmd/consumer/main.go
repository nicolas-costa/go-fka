package main

import "fmt"

func main() {
	fmt.Println("Consumer")

	app := initialize()

	app.Start()
}
