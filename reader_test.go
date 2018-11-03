package main

import (
	"testing"
)

func Test_TextReader_Load_should_load_lines_from_file(t *testing.T) {

	reader := TextFileReader{}
	count, err := reader.load("test_fixtures/9titles.txt")

	if err != nil {
		t.Errorf("did not expect err, got %s", err.Error())
	}

	expectedLineCount := 9
	if count != expectedLineCount {
		t.Errorf("expected %d lines, got %d", expectedLineCount, count)
	}
}

func Test_TextReader_Load_should_complain_load_lines_from_non_existing_file(t *testing.T) {
	reader := TextFileReader{}
	_, err := reader.load("does/not/exists.txt")

	if err == nil {
		t.Errorf("expect err, got nil")
	}
}

func Test_SandboxLoad_should_extract_titles_from_file(t *testing.T) {

	reader := SandboxJsonReader{}
	count, err := reader.load("test_fixtures/9titles.json")

	if err != nil {
		t.Errorf("did not expect err, got %s", err.Error())
	}

	expectedLineCount := 9
	if count != expectedLineCount {
		t.Errorf("expected %d lines, got %d", expectedLineCount, count)
	}
}

func Test_SandboxLoad_should_complain_when_non_existing_file_given(t *testing.T) {

	reader := SandboxJsonReader{}
	_, err := reader.load("does/not/exists.json")

	if err == nil {
		t.Errorf("expect err, got nil")
	}
}
