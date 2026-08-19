package main

import (
	"fmt"
	"os"
)

// mock_server - Dynamic mock server
func mock_server(path string) {
	fmt.Println("========================================")
	fmt.Println("  Mock-Server")
	fmt.Println("  Dynamic mock server")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	mock_server(path)
}
