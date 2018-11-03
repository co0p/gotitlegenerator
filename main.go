package main

import (
	"flag"
	"log"
)

func main() {

	trainer := Trainer{}

	var jsonFile string
	var textFile string
	var chainFile string
	var markovOrder uint

	flag.StringVar(&jsonFile, "json", "", "if specified loads lines from sandbox-json")
	flag.StringVar(&textFile, "txt", "", "if specified loads lines from text file")
	flag.UintVar(&markovOrder, "order", 1, "specifies the order of the markov chain to train")
	flag.StringVar(&chainFile, "chain", "", "specifies the markov chain file to load")
	flag.Parse()

	var lines []string
	if jsonFile != "" {
		lines = readFromFile(jsonFile, &SandboxJsonReader{})
	}
	if textFile != "" {
		lines = readFromFile(textFile, &TextFileReader{})
	}

	if lines != nil {
		if err := trainer.train(lines, markovOrder); err != nil {
			log.Fatalf("failed to train the chain of order %d: %s", markovOrder, err.Error())
		}
		if err := trainer.save(); err != nil {
			log.Fatalf("failed to save the chain: %s", err.Error())
		}
	}

	if chainFile != "" {
		err := trainer.load(chainFile)
		if err != nil {
			log.Fatalf("failed to load chain %s: %s", chainFile, err.Error())
		}
	}

}

func readFromFile(path string, reader FileReader) []string {
	count, err := reader.load(path)
	if err != nil {
		log.Fatalf("failed loading from file: %s", err.Error())
	}

	log.Printf("loaded %d lines from file %s", count, path)
	return reader.getLines()
}
