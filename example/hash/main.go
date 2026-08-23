// Using the Sodium library to calculate a hash of a string.
//
// Usage:
//
//	./build/hash 'Hello, World!'
package main

import (
	"solod.dev/so/c"
	"solod.dev/so/encoding/hex"
	"solod.dev/so/errors"
	"solod.dev/so/fmt"
	"solod.dev/so/os"
	"solod.dev/sodium/libsodium"
)

var ErrHash = errors.New("hash failed")

func main() {
	if libsodium.Init() < 0 {
		fmt.Fprintf(os.Stderr, "error: libsodium could not be initialized\n")
		os.Exit(1)
	}

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <string>\n", os.Args[0])
		os.Exit(1)
	}

	msg := os.Args[1]
	hash := make([]byte, libsodium.Generichash_BYTES)
	if err := calcHash(hash, msg); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}

	hexd := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(hexd, hash)
	fmt.Printf("hash(%s) = %s\n", msg, string(hexd))
}

// calcHash computes the hash of the input string s and
// writes the result into the provided hash slice.
func calcHash(hash []byte, s string) error {
	rc := libsodium.Generichash(
		c.SliceData[c.UChar](hash), c.Size(len(hash)),
		c.StringData[c.UChar](s), c.ULongLong(len(s)),
		nil, 0)
	if rc != 0 {
		return ErrHash
	}
	return nil
}
