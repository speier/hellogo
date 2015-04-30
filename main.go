package main

import (
	"fmt"
	"github.com/speier/hellogo/learn"
)

func main() {
	fmt.Println("hello go")
	// funcs
	fmt.Println("funcs: named return =", learn.NamedReturn())
	fmt.Printf("funcs: multi return = ")
	fmt.Println(learn.MultipleReturn(2, 3))
	// types
	fmt.Println("types: primitives   =", learn.PrimitiveTypes())
	fmt.Println("types: interfaces   =", learn.Interfaces())
	// chans
	fmt.Println("chans: send-recieve =", learn.SimpleChannels())
	// webs
	fmt.Println("wserv: srv response =", learn.WebServer())
}
