package main

import "os"

func main() {
	os.Exit(1) // want "Usage of os.Exit in the main package is banned"
}

func notMain() {
	os.Exit(1)
}
