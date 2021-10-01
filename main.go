package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	// defer db.Close()
	// cli.Start()

	difficulty := 3
	target := strings.Repeat("0", difficulty)
	nonce := 1

	for {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("hello"+fmt.Sprint(nonce))))
		fmt.Printf("Hash:%s\n", hash)
		fmt.Printf("Target:%s\n", target)
		fmt.Printf("Nonce:%d\n\n", nonce)
		if strings.HasPrefix(hash, target) {
			return
		} else {
			nonce++
		}
	}
}
