package main

import (
	"encoding/base32"
	"fmt"
	"hash"
	"hash/fnv"
	"os"
	"strings"
)

func myHash() hash.Hash {
	return fnv.New64()
}

var mHash = func() hash.Hash {
	return fnv.New64()
}

// Kubernetes names must match the regexp '[a-z0-9]([-a-z0-9]*[a-z0-9])?'.
var encoding = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")

func main() {
	//fmt.Printf("id2name\n")
	if len(os.Args) <= 1 || os.Args[1] == "" {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: dexi2n <clientID>\n")
		os.Exit(2)
	}
	id := os.Args[1]
	fmt.Printf("%s\n", strings.TrimRight(encoding.EncodeToString(mHash().Sum([]byte(id))), "="))
}
