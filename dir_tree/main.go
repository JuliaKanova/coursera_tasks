package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func dirTree(dirname string, withFile bool) {
	var visitedDir []string
	err := filepath.Walk(dirname,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !withFile {
				if info.IsDir() {
					fmt.Println(path, info.Size())
					visitedDir = append(visitedDir, path)

				}
			} else {
				fmt.Println(path, info.Size())
				visitedDir = append(visitedDir, path)
			}
			return nil
		})
	fmt.Println(visitedDir)
	if err != nil {
		log.Println(err)
	}
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
