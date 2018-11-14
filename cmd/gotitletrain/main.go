package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/co0p/gotitlegenerator/pkg/model"
)

func main() {

	// doing the flag dance
	var textFile string
	var outputFile string
	flag.StringVar(&textFile, "txt", "", "if specified loads lines from text file")
	flag.StringVar(&outputFile, "output", "output.mc", "output file to write model to")
	flag.Parse()

	if textFile == "" {
		fmt.Println("missing 'txt' argument")
		os.Exit(1)
	}

	// so we have a file? open it!
	reader, err := os.Open(textFile)
	if err != nil {
		fmt.Printf("failed to open file: %s\n", err.Error())
		os.Exit(1)
	}

	// go through the file line by line and add the line to the model
	model := model.NewModel()
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		model.Add(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("failed parsing file: %s\n", err.Error())
	}

	if err := model.ToFile(outputFile); err != nil {
		fmt.Printf("failed writing model: %s\n", err.Error())
	}
}
