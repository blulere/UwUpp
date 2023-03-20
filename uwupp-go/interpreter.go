package main

import (
	"fmt"
	"strings"
)

// easier interop with python and go codebases
// TODO: move all language definitions into a common language
type Line struct {
	s string
}

func (l Line) replace(repl string, with string) Line {
	// fmt.Println("warning: don't use this anymore.")
    l.s = strings.Replace(l.s, repl, with, -1)
    return l
}

func (l Line) Contains(_ Line, s string) bool {
	if strings.Contains(l.s, s) { 
		return true
	} else { 
		return false
	}
}

func Uwupp_to_python(fs string) (string, error) {
	var pythoncode string
	var err error = nil

	// split fs into lines

	lineslist := strings.Split(fs, "\n")

	current_line := 0

	for i := 0; i < len(lineslist); i++ {
		// shorthand
		// screw Go and not having classes :3
		line := Line {lineslist[i],}
		
		// each line is automatically this until it isn't.
		is_comment := false
		is_string := false

		if !is_comment {
			// this line was never a coment...
			if line.Contains(line, "UwU") {
				// but now the rest of it is
				is_comment = true
				line = line.replace("UwU", "# UwU")
			}
		}

		// if !is_string {
			// this line was never a string...
			// if line.Contains(line, "\"\"")
				// TODO: implement string ignoring
			// }
		// }

		// errors
		if line.Contains(line, "notices i iws 0") {
			return "", fmt.Errorf("error: for loops are not supported at this time due to the transpiler's limitations at the moment. please convert this to a while loop")
		}
		// warnings
		if line.Contains(line, "gweatew twan") {
			fmt.Printf("warning: this uses issue syntax from the original version of UwU++ (gweatew twan). make sure to fix this to avoid deprecation (line %v)\n", current_line)
			line = line.replace("gweatew twan", ">")
		}
		if line.Contains(line, "wess twan") {
			fmt.Printf("warning: this uses issue syntax from the original version of UwU++ (wess twan). make sure to fix this to avoid deprecation (line %v)\n", current_line)
			line = line.replace("wess twan", "<")
		}
		if line.Contains(line, "eqwall twoo") {
			fmt.Printf("warning: this uses issue syntax from the original version of UwU++ (eqwall twoo). make sure to fix this to avoid deprecation (line %v)\n", current_line)
			line = line.replace("eqwall twoo", "==")
		}
		if line.Contains(line, "nuzzels") {
			fmt.Printf("warning: this uses issue syntax from the original version of UwU++ (nuzzels). make sure to fix this to avoid deprecation (line %v)\n", current_line)
			line = line.replace("nuzzels(", "print(")
		}

		if !is_string {
			line = line.replace("stwing", "str")
			line = line.replace("iws", "=")
			line = line.replace("OwO *notices", "while")
			line = line.replace("ewse *notices", "elif")
			line = line.replace("ewse", "else:")
			line = line.replace("*notices", "if")
			line = line.replace("wess than", "<")
			line = line.replace("gweatew than", ">")
			line = line.replace("not eqwaws", "!=")
			line = line.replace("*", ":")
			line = line.replace("twimes", "*")
			line = line.replace("moduwo", "%")
			line = line.replace("eqwaw", "==")
			line = line.replace("diwide", "/")
			line = line.replace("nuzzles(", "print(")
			line = line.replace("stawp", "")
			line = line.replace("pwus", "+")
			
			/*
			Python doesn't support fixed array allocation.
			This is a workaround.
			So far it works only with Strings, Ints
			and Nones.
			*/

			if line.Contains(line, "awway<") {
				if line.Contains(line, "stwing") {
					line = line.replace("awway<", "[\"\"]*")
				} else if line.Contains(line, "int") {
					line = line.replace("awway<", "[0]*")
				} else {
					line = line.replace("awway<", "[None]*")
				}
				line = line.replace("stwing", "str")
				line = line.replace("stawp", "")
           	 line = line.replace("pwus", "+")
			}
		}

		pythoncode += line.s
		pythoncode += "\n"
		current_line ++
	}

	return pythoncode, err
}

func Python_to_uwupp(fs string) string {
	// coming soon
	return ""
}

func Interpreter(fs string) string {
	// coming soon
	return ""
}