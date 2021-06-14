package main

import (
	"fmt"

	quoteV3 "rsc.io/quote/v3"
)

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}
	return fmt.Sprint(englishHelloPrefix, name)
}

func Proverb() string {
	return quoteV3.Concurrency()
}

func main() {
	fmt.Println(Hello("Andrew"))
}
