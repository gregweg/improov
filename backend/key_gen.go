package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func run() {
	key := make([]byte, 32) // 256-bit key
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(key))
}
