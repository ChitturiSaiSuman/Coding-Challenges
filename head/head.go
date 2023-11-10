package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 10, "print the first n lines instead of the first 10;")

	flag.Parse()

	files := flag.Args()

	if len(files) == 0 {
		head(os.Stdin, n)
	} else if len(files) == 1 {
		f, err := os.Open(files[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		head(f, n)
	} else {
		for i, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			defer f.Close()
			fmt.Printf("==> %v <==\n", file)
			head(f, n)
			if i != len(files)-1 {
				fmt.Println()
			}
		}
	}
}

func head(f *os.File, n int) {
	scanner := bufio.NewScanner(f)
	for n > 0 && scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		n--
	}
}
