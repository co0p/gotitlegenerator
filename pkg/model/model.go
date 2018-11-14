package model

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// CURRENT_VERSION indicates the current version of the model. This is necessary for loading models
// and ensuring consistency. See https://golang.org/src/encoding/gob/type.go?s=23228:23440#L783
const CURRENT_VERSION = 1

type Model struct {
	version int
	lines   map[string]Token
}

type Token []string

// NewModel returns as new Model
func NewModel() *Model {
	return &Model{
		version: CURRENT_VERSION,
		lines:   make(map[string]Token),
	}
}

// Add adds the line to the model. Overrides any existing entries with same value
func (m *Model) Add(line string) {
	s := trim(line)
	if s != "" {
		m.lines[line] = tokenize(line)
	}
}

func (m Model) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintln(&b, m.version, m.lines)
	return b.Bytes(), nil
}

func (m *Model) UnmarshalBinary(in io.Reader) error {
	n, err := fmt.Fscanln(in, &m.version, &m.lines)
	fmt.Printf("count: %d", n)
	return err
}

func (m *Model) ToFile(path string) error {
	var b bytes.Buffer
	encoder := gob.NewEncoder(&b)
	if err := encoder.Encode(m); err != nil {
		return err
	}
	return ioutil.WriteFile(path, b.Bytes(), 0644)
}

func FromFile(path string) (Model, error) {
	var m Model
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return Model{}, err
	}
	encoder := gob.NewDecoder(bytes.NewReader(b))
	err = encoder.Decode(&m)
	if m.version != CURRENT_VERSION {
		return Model{}, fmt.Errorf("could not load model: version mismatch. Wanted %d, got %d", CURRENT_VERSION, m.version)
	}

	return m, err
}

func tokenize(s string) Token {
	return Token(strings.Fields(s))
}

func trim(s string) string {
	line := strings.TrimSpace(s)
	if strings.HasPrefix(line, "#") {
		return ""
	}
	return line
}
