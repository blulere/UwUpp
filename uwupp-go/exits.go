package main

import (
	"fmt"
	"os"
)

// Prints help info and exits.
func show_help_and_exit() {
	fmt.Println("uwupp-go help")
	fmt.Println("point to an .uwu, .uwupp or .uwu++ file to run it.")
	fmt.Println("use --transpile or -t flag to convert it to a python file.")
	fmt.Println("use --run-after or -r flag to run the file immediately after creation.")
	fmt.Println("")
	fmt.Println("use --versiion or -v to get version information")
	fmt.Println("use --help or -h to get to this screen")
	fmt.Println("")
	fmt.Println("use the python version if you like slow runtimes and outdated code")
	os.Exit(0)
}

// Prints version information and exits.
func version_and_exit() {
	fmt.Println("")
	fmt.Println("uwupp-go v0.2.0 Pre-release")
	fmt.Println("Not all UwU++ features are supported in this binary, and it's HIGHLY EXPERIMENTAL.")
	fmt.Println("")
	os.Exit(0)
}

// not in the python code becase python supports exceptions
func exception_and_exit(e string, err error) {
	if err == nil {
		err = fmt.Errorf("<no error was provided>")
	}
	fmt.Println("error: something went wrong")
	fmt.Println(e)
	fmt.Println(err)
	os.Exit(-1)
}
