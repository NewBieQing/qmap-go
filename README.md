# qmap-Go

An extension for go map, which  make you easier to fetch different type of value from map.


## Installation

With [Go module][] support (Go 1.11+), simply add the following import

```go
import "github.com/NewBieQing/qmap-go"
```
Otherwise, to install the `qmap-go` package, run the following command:

```console
$ go get -u github.com/NewBieQing/qmap-go
```
## Example
```
package main

import (
	"qmap-go/qmap"
)

func main() {
	jsonStr := `{"name":"lee", "age":12, "deposit":123456.78}`
	jsonQM, err := qmap.NewWithString(jsonStr)
	if err != nil {
		return
	}
	println("name:", jsonQM.String("name"))
	println("age:", jsonQM.Int("age"))
	println("deposit:", jsonQM.Float64Round("deposit", 2))
}
```