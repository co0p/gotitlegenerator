package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Trainer struct {
	lines []string
	words map[string]int
	chain Chain
}

func (trainer *Trainer) train(lines []string, order uint) error {
	log.Printf("training chain of order %d", order)
	trainer.lines = lines
	trainer.words = make(map[string]int)

	for _, line := range trainer.lines {
		tokens := strings.Fields(line)
		for _, token := range tokens {
			trainer.words[token] = trainer.words[token] + 1
		}
	}
	log.Printf("found %d individual words", len(trainer.words))

	return nil
}

func (trainer *Trainer) save() error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	trainer.chain.name = "markov chain"
	enc.Encode(trainer.chain)

	return ioutil.WriteFile("chain.mc", buf.Bytes(), os.ModePerm)
}

func (trainer *Trainer) load(path string) error {

	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(fp)
	var chain Chain
	dec.Decode(&chain)
	log.Printf("loaded chain %v from file %s", chain, path)

	return nil
}
