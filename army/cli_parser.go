package army

import "fmt"

func showHelp() {

	fmt.Println("Hello and welcome to your personal Gopher army !\nWe provide you with all the soldiers you need for your journey.")
	fmt.Println("\rTo start using the app, please pass in the arguments or call it with ? to get this menu.")
	// TODO: add list of arguments here...
}

func ParseArgs(args []string) {

	// TODO: make sure you use the `flags` package next time...

	if (len(args) == 0) || ((len(args) == 1) && (args[0] == "?" || args[0] == "help")) {
		showHelp()
	} else {

		if args[0] == "os" {
			OsDisplay()
		} else if args[0] == "arch" {
			ArchDisplay()
		} else {

			for _, arg := range args {

				fmt.Println(arg)
			}
		}
	}
}
