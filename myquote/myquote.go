package myquote

//import needed packages.
//fmt for printing output and of course
//the required quote package
import (
	"fmt"

	"rsc.io/quote"
)

//main function that prints the text from underlying functions
func PrintAllQuotes() {
	fmt.Println(PrintGlass())
	fmt.Println(PrintGo())
	fmt.Println(PrintHello())
	fmt.Println(PrintOpt())
}

//functions that return the stings of all 4 quote commands to the main function
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
