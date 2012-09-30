# procfile

A go package for parsing `Procfile` entries.

## Install

```
go get github.com/hecticjeff/procfile
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/hecticjeff/procfile"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("Procfile")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	proclist := procfile.Parse(string(data))
	fmt.Println(proclist)
}
```
