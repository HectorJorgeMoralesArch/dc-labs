// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 241.

// Crawl2 crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
//
// Crawl3 adds support for depth limiting.
//
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"gopl.io/ch5/links"
)
//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}
//!-sema
//!+
func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		log.Fatal("Not enough args")
	}
	if !strings.HasPrefix(args[0], "-depth") {
		log.Fatal("Not set depth")
	}
	data := strings.Split(args[0], "=")
	maxDepth, err := strconv.Atoi(data[1])
	if err != nil {
		log.Fatal(err.Error + "Error parsing")
	}
	if !strings.HasPrefix(args[0], "-results") {
		log.Fatal("Not set results File")
	}
	file := strings.Split(args[1], "=")
	filename := file[1]
	file,err:=os.Create(filename)
	if err != nil {
		log.Fatal(err.Error + "Error creating the file")
	}
	defer file.Close()
	worklist := make(chan []string)
	list := args[3]
	currentDepth := 0
	visited := make(map[string]bool)
	for true {
		if currentDepth > maxDepth {
			break
		}
		file.WriteString("Level:\t" + strconv.Itoa(currentDepth))
		Q := [][]string{}
		for _, link := range list {
			if !visited[link] {
				visited[link] = true
				file.WriteString("\t" + link)
				/*go func(link string) {
					links := crawl(link)
					worklist <- links
				}(link)
				*/
				go func(link string) {
					worklist <- crawl(link)
				}(link)
				Q = append(Q, <-worklist)
			}
		}
		list = []string{}
		for _, li := range Q {
			for _, link := range li {
				list = append(list, link)
			}
		}
		currentDepth++
	}
}

//!-
