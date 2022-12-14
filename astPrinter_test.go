package main

import (
	"bytes"
	"flag"
	"os"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")

func TestParse(t *testing.T) {
	golden := "testdata/a.golden"

	var result bytes.Buffer
	err := Parser("./testdata/a.go", &result)
	if err != nil {
		t.Errorf("want no err but got %v", err)
	}

	// update フラグが指定されていれば golden ファイルを更新する
	if *update {
		os.WriteFile(golden, result.Bytes(), 0644)
	}

	// golden ファイルを読み込む
	expected, err := os.ReadFile(golden)
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	// 結果が golden ファイルと変わらないことを確認
	if result.String() != string(expected) {
		t.Errorf("golden file not matched\n")
	}
}
