// --------------------------------------------- ARGS ------------------------------------------------- \\
//                                   official part of LunaCompiler                                      \\
//                                        Authors: LunaTeam                                             \\
// ---------------------------------------------------------------------------------------------------- \\
package main

import (
	"fmt"
)

func ifile(filename string) {
    functions := mainPrase(filename)
	//// ---------------------------------- TESTS ------------------------------------------------ \\s
    fmt.Println("Tests")
	for name, fn := range functions {
		fmt.Printf("Function: %s\n", name)
		fmt.Printf("  Args: %v\n", fn.Args)
		fmt.Printf("  Returns: %v\n", fn.Returns)
		fmt.Printf("  Body:\n")
		for _, line := range fn.Lines {
			fmt.Printf("    %s\n", line)
		}
	}
	/// ------------------------------------------------------------------------------------------ \\\
	// for name, fn := range functions {

	// }
}

func noargs() {
	fmt.Println("Welcome to LunaScript compiler!")
	fmt.Println("Usage: ")
	fmt.Println(`
	lunac {filename}.luna    - 
	`)
}

func help() {
	// nil //
}