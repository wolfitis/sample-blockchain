package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToHex - a utility function to convert int to hex
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// ReverseBytes - reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// https://stackoverflow.com/questions/47302402/how-to-convert-int-to-hex
// by Henry Henderson
// func IntToHex(n int64) []byte {
//     return []byte(strconv.FormatInt(n, 16))
// }
