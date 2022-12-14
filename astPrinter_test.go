package main

import (
	"bytes"
	"flag"
	"os"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")

func TestParse(t *testing.T) {
	golden := "testdata/a.go.golden"
	var result bytes.Buffer
	err := Parser("./testdata/a.go", &result)
	if err != nil {
		t.Errorf("want no err but got %v", err)
	}

	if *update {
		os.WriteFile(golden, result.Bytes(), 0777)
	}

	expected, err := os.ReadFile(golden)
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	if bytes.Equal(result.Bytes(), expected) {
		t.Errorf("golden file not matched")
	}
}
