package main

import (
	"fmt"

	quoteV3 "rsc.io/quote/v3"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func Proverb() string {
	return quoteV3.Concurrency()
}

func main() {
	fmt.Println(Hello("Andrew"))
}
