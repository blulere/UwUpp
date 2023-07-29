package main

import (
	"fmt"
	"os"
)

func show_help_and_exit() {
	fmt.Println("uwupp <file.uwupp> - Interpret a file. Compilation coming soon\n")
	fmt.Println("uwupp version - Prints version information")
	fmt.Println("uwupp help - Prints the help menu")
	os.Exit(0)
}

func version_and_exit() {
	fmt.Println("v0.3.0")
	os.Exit(0)
}
