package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadBetweenLines(filename string) {
	readfile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	filescanner := bufio.NewScanner(readfile)
	filescanner.Split(bufio.ScanLines)
	for filescanner.Scan() {
		text := filescanner.Text()
		if strings.Contains(text, "//") {
			fmt.Println(filescanner.Text())
		}
	}
	readfile.Close()
}

func main() {
	ReadBetweenLines("data.txt")
}
