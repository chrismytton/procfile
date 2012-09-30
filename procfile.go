package procfile

import (
	"regexp"
	"strings"
)

var procfileRegex = regexp.MustCompile("^([A-Za-z0-9_]+):\\s*(.+)$")

type processType struct {
	command string
	arguments []string
}

// Parse will read the contents of a procfile and returns a map of process names
// to commands.
func Parse(procfile string) (procs map[string]processType) {
	procs = make(map[string]processType)

	for _, line := range strings.Split(procfile, "\n") {
		if matches := procfileRegex.FindStringSubmatch(line); matches != nil {
			name, command := matches[1], matches[2]
			args := strings.Split(command, " ")
			procs[name] = processType{args[0], args[1:]}
		}
	}

	return
}
