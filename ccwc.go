package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Silly goose, I need parameters!")
	}

	length := 0
	fileName := new(string)
	*fileName = "default.txt"

	if os.Args[1] == "-c" {
		fileName = flag.String("c", "default.txt", "File name parameter")
		flag.Parse()
	
		file, err := os.Open(*fileName)
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewScanner(file)

		for reader.Scan() {
			length += len(reader.Text())
		}
		fmt.Printf("%d %s\n", length, *fileName)
	} else if os.Args[1] == "-l" {
		fileName = flag.String("l", "default.txt", "File name parameter")
		flag.Parse()

		file, err := os.Open(*fileName)
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewScanner(file)

		for reader.Scan() {
			length += 1
		}
		fmt.Printf("%d %s\n", length, *fileName)
	} else if os.Args[1] == "-w" {
		fileName = flag.String("w", "default.txt", "File name parameter")
		flag.Parse()

		file, err := os.Open(*fileName)
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewScanner(file)

		for reader.Scan() {
			length += len(strings.Fields(reader.Text()))
		}

		fmt.Printf("%d %s\n", length, *fileName)
	} else {
		file, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewScanner(file)
		var (
			c = 0
			l = 0
			w = 0
		)
		for reader.Scan() {
			c += len(reader.Text())
			l += 1
			w += len(strings.Fields(reader.Text()))
		}

		fmt.Printf("%d %d %d %s\n", c, l, w, os.Args[1])

	}
}