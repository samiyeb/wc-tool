package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Parameters are needed")
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
			length += len(reader.Bytes()) + 1
		}
		fmt.Printf("%d %s\n", length - 1, *fileName)
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
		fmt.Printf("%d %s\n", length - 1, *fileName)
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
	} else if os.Args[1] == "-m" {
		fileName = flag.String("m", "default.txt", "File name parameter")
		flag.Parse()
		file, err := os.Open(*fileName)
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewScanner(file)
		for reader.Scan() {
			length += utf8.RuneCount(reader.Bytes()) + 1
		}
		fmt.Printf("%d %s\n", length - 1, *fileName)
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
			m = 0
		)
		for reader.Scan() {
			c += len(reader.Bytes()) + 1
			l += 1
			w += len(strings.Fields(reader.Text()))
			m += utf8.RuneCount(reader.Bytes()) + 1
		}
		fmt.Printf("%d %d %d %d %s\n", c-1, l-1, w, m-1, os.Args[1])
	}
}

// Sam Bakri