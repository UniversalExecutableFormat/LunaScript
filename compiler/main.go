// --------------------------------------------- MAIN ------------------------------------------------- \\
//                                   official part of LunaCompiler                                      \\
//                                        Authors: LunaTeam                                             \\
// ---------------------------------------------------------------------------------------------------- \\
package main

import (
	"fmt"
	"os"
	"strings"
	"golang.org/x/term"
)

var args []string = os.Args

func main() {
	parseFlags()

	switch len(args) {
	case 1:
		noargs()
	default:
		for _, arg := range args[1:] {
			switch arg {
			case "build", "compile":
				if DonePath == "///" || DonePath == "" {
					w, _, err := term.GetSize(int(os.Stdout.Fd()))
					if err != nil {
						fmt.Println("Error: ", err)
						return
					}
					w -= lenr(" Building with default output path... ")
					w /= 2
					fmt.Println(strings.Repeat("-", w), pog + "Building with default output path..." + reset, strings.Repeat("-", w))
				} else {
					fmt.Printf("Building to: %s\n", DonePath)
				}
				ifile(getInputFile())
				return
			case "help":
				help()
				return
			}
		}
		fmt.Println("Unknown command.")
	}
}

func parseFlags() {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-o=") {
			DonePath = strings.TrimPrefix(arg, "-o=")
			return
		} else if strings.HasPrefix(arg, "--output=") {
			DonePath = strings.TrimPrefix(arg, "--output=")
			return
		}
		DonePath = "///"
	}
}

func getInputFile() string {
	for _, arg := range args[1:] {
		if !strings.HasPrefix(arg, "-o=") && !strings.HasPrefix(arg, "--output=") && arg != "build" && arg != "compile" {
			return arg
		}
	}
	return "output.bin"
}