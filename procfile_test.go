package procfile

import (
	"testing"
)

func TestParseBasic(t *testing.T) {
	basicProcfile := "web: bundle exec rackup"
	procs := Parse(basicProcfile)

	if procs["web"].command != "bundle" {
		t.Error("Expected command to be bundle but got ", procs["web"].command)
	}

	if len(procs["web"].arguments) != 2 {
		t.Error("Expected arguments to be 2 in length, got ",
			len(procs["web"].arguments))
	}

}

func TestParseMultiProcess(t *testing.T) {
	multilineProcfile := `web: coffee server.coffee
worker: node worker.js`

	procs := Parse(multilineProcfile)

	if procs["web"].command != "coffee" {
		t.Error("Expected web command to be coffee")
	}

	if procs["worker"].command != "node" {
		t.Error("Expected worker command to be node")
	}
}
