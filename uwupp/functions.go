package main

import (
	"fmt"
	"os"
)

// Shows the help menu, here to reduce clutter.
func show_help_and_exit() {
	fmt.Printf("\nuwupp <file.uwupp> - Interpret a file. Compilation coming soon\n\n")
	fmt.Println("uwupp version - Prints version information")
	fmt.Println("uwupp help - Prints the help menu")
	os.Exit(0)
}

// Prints the version of uwupp and exits.
func version_and_exit() {
	fmt.Println("v0.3.0")
	os.Exit(0)
}

// exception_and_exit() is deprecated.
