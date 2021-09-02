package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github/fosmjo/algs4/ch3/symboltable"
	"os"
	"unicode"
)

var minKeyLen = flag.Int("l", 8, "min length of key")

func main() {
	flag.Parse()
	st := symboltable.NewBinarySearchST(12000)
	scanner := bufio.NewScanner(os.Stdin)
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			token = bytes.TrimFunc(token, func(r rune) bool { return unicode.IsPunct(r) })
		}
		return
	}
	scanner.Split(split)

	for scanner.Scan() {
		word := scanner.Text()
		if len(word) < *minKeyLen {
			continue
		}

		count, _ := st.Get(word)
		st.Put(word, count+1)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var (
		key          string
		maxFrequency int
	)

	for _, word := range st.Keys() {
		count, _ := st.Get(word)
		if count > maxFrequency {
			key = word
			maxFrequency = count
		}
	}

	fmt.Println(key, maxFrequency)
}
