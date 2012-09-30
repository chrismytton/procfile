package procfile

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := os.Open("test/fixtures/Procfile")
	if err != nil {
		t.Error(err)
	}
	procs, err := Parse(file)
	if err != nil {
		t.Error(err)
	}
	if procs["web"] != "bundle exec rackup" {
		t.Error("Unexpected web process:", procs)
	}
}
