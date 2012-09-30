package procfile

import (
	"os"
	"testing"
)

func TestParseProcfile(t *testing.T) {
	file, err := os.Open("test/fixtures/Procfile")
	if err != nil {
		t.Error(err)
	}
	procs, err := ParseProcfile(file)
	if err != nil {
		t.Error(err)
	}
	if procs["web"] != "bundle exec rackup" {
		t.Error("Unexpected web process:", procs)
	}
}
