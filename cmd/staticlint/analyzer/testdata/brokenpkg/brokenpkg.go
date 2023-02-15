package main

import "os"

type dummy struct{}

func (d *dummy) Exit() {

}

func main() {
	os.Exit(1) // want "Usage of os.Exit in the main package is banned"

	d := dummy{}
	d.Exit()
}

func notMain() {
	os.Exit(1)
}
