package procfile

import (
	"regexp"
	"strings"
)

var procfileRegex = regexp.MustCompile("^([A-Za-z0-9_]+):\\s*(.+)$")

// Parse will read the contents of a procfile and returns a map of process names
// to commands.
func Parse(procfile string) (ret map[string]string) {
	ret = make(map[string]string)

	lines := strings.Split(procfile, "\n")
	for _, line := range lines {
		matches := procfileRegex.FindStringSubmatch(line)
		if matches != nil {
			ret[matches[1]] = matches[2]
		}
	}

	return
}
