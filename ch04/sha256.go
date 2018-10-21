package main

import (
	"fmt"
	"crypto/sha256"
)

func main() {
	chksum1 := sha256.Sum256([]byte("x"))
	chksum2 := sha256.Sum256([]byte("X"))
	fmt.Printf("chksum1=%x\nchksum2=%x\n%t\n%T\n", chksum1, chksum2, chksum1==chksum2, chksum1)
}


