// Package is for analyzer test only
package main

import "os"

// since analyzer shall only ban exits from the os struct, we declare a dummy struct to show
// it's Exit is fine for the analyzer
type dummy struct{}

// Exit func of dummy struct won't be banned by the analyze
func (d *dummy) Exit() {

}

func main() {
	os.Exit(1) // want "Usage of os.Exit in the main package is banned"

	d := dummy{}
	d.Exit()
}

// Analyzer will only ban os.Exits in the main func of the package
func notMain() {
	os.Exit(1)
}
