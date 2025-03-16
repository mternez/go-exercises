package main

import (
	"fmt"
	"html-link-parser/parser"
)

func main() {
	//fmt.Println(parser.ParseLinks("ex1.html"))
	//fmt.Println(parser.ParseLinks("ex2.html"))
	//fmt.Println(parser.ParseLinks("ex3.html"))
	//fmt.Println(parser.ParseLinks("ex4.html"))
	fmt.Println(parser.ParseWithXHTML("ex1.html"))
	fmt.Println(parser.ParseWithXHTML("ex2.html"))
	fmt.Println(parser.ParseWithXHTML("ex3.html"))
	fmt.Println(parser.ParseWithXHTML("ex4.html"))
}
