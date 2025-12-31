package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4")
	exp := 4
	res, err := count(b, false, false)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := 3
	res, err := count(b, true, false)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4")
	exp := 23
	res, err := count(b, false, true)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}
