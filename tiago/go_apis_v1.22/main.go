package main

import "fmt"

func main() {
	server := NewApiServer(":8099")
	err := server.Run()
	fmt.Println(err)
}
