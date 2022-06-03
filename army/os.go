package army

import (
	"armyOGophers/cfg"
	"fmt"
)

func OsDisplay() {
	fmt.Println("Your operating system is: " + cfg.UserOs)
}

func ArchDisplay() {
	fmt.Println("CPU Architecture: " + cfg.UserOsArch)
}
