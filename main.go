package main

import (
	"armyOGophers/army"
	"os"
)

func main() {
	army.ParseArgs(os.Args[1:])
}
