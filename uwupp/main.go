package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var input_filepath string
	input_file := false

	// start a timer for checking how long ""compilation"" takes
	// maybe this should be started later in the program as to not include parsing the request time
	// or just be removed completely
	timer_start := time.Now().UnixMilli()

	if len(os.Args) > 1 {
		// command-line arguments were passed
		for i := 0; i < len(os.Args); i++ {
			// shorthand
			arg := os.Args[i]
			if strings.Contains(arg, "help") {
				show_help_and_exit()
			} else if strings.Contains(arg, "version") {
				version_and_exit()
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
		fmt.Println("fatal: no arguments were specified.")
		fmt.Println("use argument 'help' for help")
		os.Exit(0)
	}

	// open the input file
	if !input_file {
		fmt.Println("fatal: no input file was specified, yet uwupp is attempting to read from one, which is unintended behaviour.")
		os.Exit(-1)
	}
	// read input with error handling
	fs, err := os.ReadFile(input_filepath)
	// SCREW YOU GO !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	// 22/03/2024 - just dropped in, thanks Go for making files byte[] and have to convert to strings, better than runes though LMFAOO
	fs2 := string(fs)
	_ = fs2
	if err != nil {
		fmt.Println("fatal: opening the file failed! it probably doesn't exist, or something went wrong.")
		fmt.Printf("error: %v\n", err)
		os.Exit(-2)
	}

	// fs2, err = uwupp_to_python(fs2)
	err = interpret(fs2)
	// err = nil
	if err != nil {
		fmt.Printf("fatal: code execution failed.\n%v\n", err)
	}
	// interpretation completed

	// this code doesn't work sometimes and I have no idea why.
	// it seems a lot more stable when running through ioutil which is weird because it's the same damn function.
	// err = os.WriteFile(newfile_name, []byte(fs2), 0644)
	// 22/03/2024 Well I have good news for you blulere, this no longer transpiles lol it's an interpreter now

	/* if err != nil {
		("error: writing to the new python file failed. you're likely out of storage space or your antivirus blocked the program...", err)
	} */
	// get timer in seconds and milliseconds
	timer := float32(time.Now().UnixMilli()-timer_start) / 1000
	fmt.Printf("done, operation completed in %v seconds\n", timer)

	// done!
}
