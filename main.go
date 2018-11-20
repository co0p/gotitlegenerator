package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mb-14/gomarkov"
)

func main() {

	var trainingFile string
	var ngram int
	var maxWords int
	flag.StringVar(&trainingFile, "input", "model.txt", "the path to load the titles from")
	flag.IntVar(&ngram, "ngram", 1, "ngram size to generate the chain from")
	flag.IntVar(&maxWords, "words", 6, "number of words to create the title from")

	flag.Parse()

	chain, err := trainModel(trainingFile, ngram)
	if err != nil {
		fmt.Println("failed training model:", err)
	}

	fmt.Println(generateTitle(chain, maxWords))
}

func trainModel(trainingFile string, ngram int) (*gomarkov.Chain, error) {

	chain := gomarkov.NewChain(ngram)
	reader, err := os.Open(trainingFile)
	if err != nil {
		fmt.Printf("failed to open file: %s\n", err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		chain.Add(strings.Split(scanner.Text(), " "))
	}
	return chain, nil
}

func generateTitle(chain *gomarkov.Chain, maxWords int) string {
	tokens := []string{gomarkov.StartToken}
	count := 0
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		count = count + 1
		next, _ := chain.Generate(tokens[(len(tokens) - 1):])
		tokens = append(tokens, next)
	}
	return (strings.Join(tokens[1:len(tokens)-1], " "))
}
