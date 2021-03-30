package main

import (
	"fmt"
	"runtime/debug"
)

var version = "dev"

func main() {
	info, _ := debug.ReadBuildInfo()
	fmt.Println("hello world, gitlab! (", info.Main, ")")
}
