package main

import "fmt"

func main() {
	//Convert Hex TO Base 64
	baseStr := convertHexToB64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fmt.Println(baseStr == "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")
	//Get XOR of two strings
	xorStr := getXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	fmt.Println(xorStr == "746865206b696420646f6e277420706c6179")
	//
}
