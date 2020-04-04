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

// https://stackoverflow.com/questions/47302402/how-to-convert-int-to-hex
// by Henry Henderson
// func IntToHex(n int64) []byte {
//     return []byte(strconv.FormatInt(n, 16))
// }
