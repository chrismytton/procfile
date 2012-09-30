package procfile

import (
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

var procfileRegex = regexp.MustCompile("^([A-Za-z0-9_]+):\\s*(.+)$")

func Parse(procfile io.Reader) (ret map[string]string, err error) {
	ret = make(map[string]string)
	data, err := ioutil.ReadAll(procfile)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		matches := procfileRegex.FindStringSubmatch(line)
		if matches != nil {
			ret[matches[1]] = matches[2]
		}
	}
	return
}
