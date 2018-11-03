package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type FileReader interface {
	load(string) (int, error)
	getLines() []string
}

type SandboxJsonReader struct {
	lines []string
}

type SandboxResponse struct {
	Documents []DocumentResponse
}

type DocumentResponse struct {
	Title string
}

func (reader *SandboxJsonReader) load(path string) (int, error) {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("failed opening file: %s", err.Error())
	}

	var res SandboxResponse
	json.Unmarshal(file, &res)

	for _, v := range res.Documents {
		reader.lines = append(reader.lines, strings.TrimSpace(v.Title))
	}

	return len(reader.lines), nil

}

func (reader *SandboxJsonReader) getLines() []string {
	return reader.lines
}

type TextFileReader struct {
	lines []string
}

func (reader *TextFileReader) load(path string) (int, error) {
	fp, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("failed reading file: %s", err.Error())
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		reader.lines = append(reader.lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed parsing file: %s", err.Error())
	}

	return len(reader.lines), nil
}

func (reader *TextFileReader) getLines() []string {
	return reader.lines
}
