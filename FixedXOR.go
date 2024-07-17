package main

import "encoding/hex"

func getXOR(first, second string) string {
	firstBytes, _ := hex.DecodeString(first)
	secondBytes, _ := hex.DecodeString(second)
	xorResult := make([]byte, len(firstBytes))
	for i := 0; i < len(firstBytes); i++ {
		xorResult[i] = firstBytes[i] ^ secondBytes[i]
	}
	xorHexString := hex.EncodeToString(xorResult)
	return xorHexString
}
