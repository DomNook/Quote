package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(PrintGlass())
	fmt.Println(PrintGo())
	fmt.Println(PrintHello())
	fmt.Println(PrintOpt())
}

func PrintGlass() string {
	return quote.Glass()
}

func PrintGo() string {
	return quote.Go()
}

func PrintHello() string {
	return quote.Hello()
}

func PrintOpt() string {
	return quote.Opt()
}
