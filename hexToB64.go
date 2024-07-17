package main

import (
	"encoding/base64"
	"encoding/hex"
)

func convertHexToB64(hexString string) string {
	hexByte, _ := hex.DecodeString(hexString)
	base64String := base64.StdEncoding.EncodeToString(hexByte)
	return base64String
}
