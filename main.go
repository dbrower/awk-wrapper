package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	blankLine := regexp.MustCompile("\n[ \t]*\n")
	if len(os.Args) < 3 {
		fmt.Println("usage: awk-wrapper <command file> <target files>...")
		os.Exit(1)
	}
	commandfile := os.Args[1]
	fmt.Println("Using awk file", commandfile)
	for _, inputfile := range os.Args[2:] {
		fmt.Println("Processing", inputfile)
		// prepare file input
		input, err := ioutil.ReadFile(inputfile)
		if err != nil {
			fmt.Println("Error opening", inputfile, ":", err)
			continue
		}
		// convert from mac line endings to unix line endings
		input = bytes.Replace(input, []byte("\r"), []byte("\n"), -1)
		// remove empty lines
		input = blankLine.ReplaceAllLiteral(input, []byte{})

		cmd := exec.Command("awk", "-f", commandfile)
		cmd.Stdin = bytes.NewReader(input)
		var out bytes.Buffer
		var errout bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &errout

		err = cmd.Run()

		if err != nil {
			fmt.Println("Error running awk:", err)
			fmt.Println(errout.String())
			break
		}

		// convert output into dos line endings
		output := out.Bytes()
		output = bytes.Replace(output, []byte("\n"), []byte("\r\n"), -1)
		ioutil.WriteFile("sierra-"+inputfile, output, 0666)
	}
}
