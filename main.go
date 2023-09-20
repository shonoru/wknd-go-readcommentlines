package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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

func Walk(path string) {
	subDirToSkip := ".git"

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() && info.Name() == subDirToSkip {
			fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
			return filepath.SkipDir
		}

		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)

		if !info.IsDir() {
			ReadBetweenLines(path)
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	walkroot := os.Getenv("WALK_ROOT")
	fmt.Println("WALK_ROOT:", os.Getenv(walkroot))
	Walk(walkroot)
}
