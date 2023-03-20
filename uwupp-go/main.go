package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// Prints help info and exits.
func show_help_and_exit() {
	fmt.Println("uwupp (go version) help")
	fmt.Println("point to an .uwu, .uwupp or .uwu++ file to run it.")
	fmt.Println("use --transpile or -t flag to convert it to a python file.")
	fmt.Println("use --run-after or -r flag to run the file immediately after creation.")
	fmt.Println("")
	fmt.Println("use the python version if you like slow runtimes and outdated code")
	os.Exit(0)
}

// Prints version information and exits.
func version_and_exit() {
	fmt.Println("")
    fmt.Println("v0.1.0 Pre-release")
    fmt.Println("Not all UwU++ features are supported in this binary, and it's HIGHLY EXPERIMENTAL.")
	fmt.Println("")
	os.Exit(0)
}

func exception_and_exit(e string, err error) {
	if err == nil {
		err = fmt.Errorf("<no error was provided>")
	}
	fmt.Println("error: something went wrong")
	fmt.Println(e)
	fmt.Println(err)
	os.Exit(-1)
}

func main() {
	transpile := false
	_ = transpile
	input_file := false
	run_after := false
	input_filepath := ""

	timer_start := time.Now().UnixMilli()

	if len(os.Args) > 1 {
		// command-line arguments were passed
		for i := 0; i < len(os.Args); i++ {
			// shorthand
			arg := os.Args[i]
			if strings.Contains(arg, "-v") ||
			   strings.Contains(arg, "--version") {
				version_and_exit()
			}
			if strings.Contains(arg, "-t") ||
			   strings.Contains(arg, "--transpile") {
				transpile = true
			}
			if strings.Contains(arg, "-r") ||
			   strings.Contains(arg, "--run-after") {
				run_after = true	
			}
			if strings.Contains(arg, ".uwu") ||
			   strings.Contains(arg, ".uwupp") ||
			   strings.Contains(arg, ".uwu++") {
				// there is an input file
				input_file = true
				input_filepath = arg
			}
		}
	} else {
		// command-line arguments were not passed
		fmt.Println("error: no input files were specified")
		fmt.Println("------------------------------------")
		show_help_and_exit()
	}

	// open the input file
	if !input_file {
		exception_and_exit("\"something weird happened somehow and no input file was assigned yet it's trying to do it anyways\"", nil)
	}
	fs, err := os.ReadFile(input_filepath)
	// SCREW YOU GO !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	fs2 := string(fs)
	if err != nil { exception_and_exit("error: opening the file failed! it probably doesn't exist, or something went wrong.", err) }

	// transpile UwU++ to python code
	newfile_name := strings.Replace(input_filepath, ".uwupp", ".py", 1)
	newfile_name = strings.Replace(newfile_name, ".uwu", ".py", 1)
	newfile_name = strings.Replace(newfile_name, ".uwu++", ".py", 1)
	fmt.Println("The new filename is " + newfile_name)
	fs2, err = Uwupp_to_python(fs2)
	if err != nil { exception_and_exit("error: transpiling the UwU++ code to python code failed.", err) }
	fmt.Println("Getting python code worked")
	// transpilation completed
	// write to the file now
	fmt.Println("Now writing to the file")

	err = os.WriteFile(newfile_name, []byte(fs2), 0644)
	
	if err != nil { exception_and_exit("error: writing to the new python file failed. you're likely out of storage space or your antivirus blocked the program...", err) }
	// get timer in seconds and milliseconds
	timer := float32 ( time.Now().UnixMilli() - timer_start ) / 1000
	fmt.Printf("done, operating completed in %v seconds\n", timer)
	if run_after {
		fmt.Println("Now running the file")
		if runtime.GOOS == "windows" {
			exec.Command("python main.py -t " + newfile_name)
		} else {
			exec.Command("python3 main.py -t " + newfile_name)
		}
		// exec.Command has no error variable so no need to check for errors here
	}
}