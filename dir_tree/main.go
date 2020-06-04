package main

import (
	"flag"
	"fmt"
	"os"
)

func dirTree(dirname string, withFile bool) {
	fmt.Println(dirname, withFile)
}

func main() {
	withFile := flag.Bool("f", false, "show files")
	flag.Parse()
	dirname := flag.Arg(0)
	finfo, err := os.Stat(dirname)
	if err != nil {
		panic(err)
	}
	if finfo.IsDir() {
		dirTree(dirname, *withFile)
	} else {
		fmt.Println("it's not a directory")
	}
}
