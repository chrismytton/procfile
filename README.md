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
	"os"
)

func main() {
	file, err := os.Open("Procfile")
	if err != nil {
		panic(err)
	}
	proclist, err := procfile.Parse(file)
	fmt.Println(proclist)
}
```
