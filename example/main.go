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
