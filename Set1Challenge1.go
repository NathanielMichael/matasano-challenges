package main

import b64 "encoding/base64"
import hex "encoding/hex"
import "fmt"

func main() {
    data := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f" +
            "69736f6e6f7573206d757368726f6f6d"

    hexData, _ := hex.DecodeString(data)
    hexEnc := b64.StdEncoding.EncodeToString(hexData)

    fmt.Println(hexEnc)
}