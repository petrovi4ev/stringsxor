package main

import (
	"fmt"
)

func main() {
	key := "ЛШТШФУМ Ащьф"
	src := "plaintext (not secret)"
	cypher := XorStrings(key, src)
	fmt.Printf("Encrypt: %s ---> %s\n", src, cypher)
	fmt.Printf("Decrypt: %s ---> %s\n", cypher, XorStrings(key, cypher))
}

func XorStrings(key, src string) string {
	res := make([]byte, 0)
	byteArrSrc := []byte(src)
	byteArrKey := []byte(key)

	for i := 0; i < len(byteArrSrc); i++ {
		res = append(res, byteArrSrc[i]^byteArrKey[i%len(byteArrKey)])
	}

	return string(res)
}
