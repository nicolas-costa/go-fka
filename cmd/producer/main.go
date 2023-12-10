package main

import "fmt"

func main() {
	fmt.Println("Producer")

	app := initialize()

	app.Start()
}
