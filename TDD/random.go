package main

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

type T struct {
	F int `yaml:"a,omitempty"`
	B int
}

var t T

func main() {
	yaml.Unmarshal([]byte("a: 1"), &t)
	fmt.Printf("%#v", t)
}
