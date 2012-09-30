package procfile

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := os.Open("test/fixtures/Procfile")
	if err != nil {
		t.Error(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
	}

	procs := Parse(string(data))
	if procs["web"] != "bundle exec rackup" {
		t.Error("Unexpected web process:", procs)
	}
}
