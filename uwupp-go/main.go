package main

// This release's version will be v0.2.0-go.

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

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
			if strings.Contains(arg, "-h") ||
				strings.Contains(arg, "--help") {
				show_help_and_exit()
			}
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
		os.Exit(0)
	}

	// open the input file
	if !input_file {
		exception_and_exit("\"something weird happened somehow and no input file was assigned yet it's trying to do it anyways\"", nil)
	}
	fs, err := os.ReadFile(input_filepath)
	// SCREW YOU GO !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	fs2 := string(fs)
	if err != nil {
		exception_and_exit("error: opening the file failed! it probably doesn't exist, or something went wrong.", err)
	}

	// transpile UwU++ to python code
	newfile_name := strings.Replace(input_filepath, ".uwupp", ".py", 1)
	newfile_name = strings.Replace(newfile_name, ".uwu", ".py", 1)
	newfile_name = strings.Replace(newfile_name, ".uwu++", ".py", 1)
	fs2, err = uwupp_to_python(fs2)
	if err != nil {
		exception_and_exit("error: transpiling the UwU++ code to python code failed.", err)
	}
	// transpilation completed

	// this code doesn't work sometimes and I have no idea why.
	// it seems a lot more stable when running through ioutil which is weird because it's the same damn function.
	err = ioutil.WriteFile(newfile_name, []byte(fs2), 0644)

	if err != nil {
		exception_and_exit("error: writing to the new python file failed. you're likely out of storage space or your antivirus blocked the program...", err)
	}
	// get timer in seconds and milliseconds
	timer := float32(time.Now().UnixMilli()-timer_start) / 1000
	fmt.Printf("done, operation completed in %v seconds\n", timer)
	// FIXME: WHAT THE HELL IS WRONG WITH THIS CODE!?!?!?!?!?!?!?!?!?!? WHY DOESN'T IT WORK???????
	if run_after {
		if runtime.GOOS == "windows" {
			cmd := exec.Command("cmd", "python", "main.py", "-t"+newfile_name)
			out, err := cmd.CombinedOutput()
			fmt.Println(string(out))
			if err != nil {
				exception_and_exit("error: something went wrong whilst trying to run the python file. maybe you're out of storage?", err)
			}
		} else {
			cmd := exec.Command("bash", "python3", "main.py", "-t"+newfile_name)
			out, err := cmd.Output()
			fmt.Println(string(out))
			if err != nil {
				exception_and_exit("error: something went wrong whilst trying to run the python file. maybe you're out of storage?", err)
			}
		}
		// exec.Command has no error variable so no need to check for errors here
	}
}
