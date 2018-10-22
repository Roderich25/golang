//go run hello2.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	const name, lastname string = "Rodrigo", "Avila"
	greeting := "HI"
	fmt.Println(strings.ToLower(greeting), name, lastname)
}