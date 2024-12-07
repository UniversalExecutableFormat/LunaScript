//	--------------------------------------------- DATA ------------------------------------------------- \\
//	                                  official part of LunaCompiler                                      \\
//	                                       Authors: LunaTeam                                             \\
//
// ---------------------------------------------------------------------------------------------------- \\
package main

import (
	"regexp"
)

var nextbody bool
var dep int
var curFunc *lunafunc
var funcName string
var est string = "\033[1;31mError!\033[0;91m "
var reset string = "\033[0m"
var ekn string = "\033[0m\n"
var pog string = "\033[1m"
var kur string = "\033[3m"
var funcRegex *regexp.Regexp = regexp.MustCompile(`fn\s+(\w+)\s*\(([^)]*)\)\s*\(([^)]*)\)`) // SHIT PLEASE HELP ME
var comndts int = 0
var DonePath string = "output.bin"
var aTypes = map[string]bool{
	"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
	"float": true, "float32": true, "float64": true,
	"rune": true, "string": true, "type": true, "bool": true,
}

var (
	bcna = "The build/compile option requires you to pass the file to compile as an argument! use: lunac build/compile FileName.luna"
	iop  = "Invalid option! Use `help` for usage information"
	tma  = "Too many arguments! Use `help` for usage information"
	fif  = "Unexpected start of a new function declaration while parsing function "
)

type lunafunc struct {
	Args    map[string]string
	Returns []string
	Lines   []string
}

func typespp() {
	for typ := range aTypes {
		aTypes[typ+"[]"] = true
	}
}
