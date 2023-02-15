package main

import "os"

type Dummy struct{}

func (d *Dummy) Exit() {
	return
}

func main() {
	os.Exit(1) // want "Usage of os.Exit in the main package is banned"

	dummy := Dummy{}
	dummy.Exit()
}

func notMain() {
	os.Exit(1)
}
