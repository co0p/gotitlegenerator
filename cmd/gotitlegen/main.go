package main

import (
	"flag"
	"fmt"

	"github.com/co0p/gotitlegenerator/pkg/model"
)

func main() {

	var inputFile string

	flag.StringVar(&inputFile, "input", "model.mc", "the path to load the model file from")
	flag.Parse()

	m, err := model.FromFile(inputFile)
	if err != nil {
		fmt.Printf("failed loading model: %s\n", err.Error())
	}

	fmt.Printf("using model: %v", m)

}
