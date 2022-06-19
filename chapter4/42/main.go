package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"os"
)

type hashType string

const (
	SHA256 hashType = "SHA256"
	SHA512 hashType = "SHA512"
	SHA384 hashType = "SHA384"
)

func main() {
	hType := ""
	flag.StringVar(&hType, "type", string(SHA256), "parse type")
	flag.Parse()
	shType := hashType(hType)

	fmt.Println("enter the initial string:")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	if len(str) > 0 {
		final := createHash(str, shType)
		fmt.Println("the string: ", str, "of type: ", shType, "was converted into: ")
		fmt.Println(final)
	}

}

func createHash(s string, shType hashType) string {
	str := ""
	var h hash.Hash
	switch shType {
	case SHA512:
		h = sha512.New()
	case SHA384:
		h = sha512.New384()
	default:
		h = sha256.New()
	}
	h.Write([]byte(s))
	str = string(h.Sum(nil))
	return str
}
