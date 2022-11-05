package main

import (
	"fmt"

	"github.com/desktop-go/src"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("Program starting")

	src.StartApp()
}
