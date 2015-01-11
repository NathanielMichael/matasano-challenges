package main

import hex "encoding/hex"
import "fmt"

func main() {

    source1 := "1c0111001f010100061a024b53535009181c"
    source2 := "686974207468652062756c6c277320657965"

    source1Hex, _ := hex.DecodeString(source1)
    source2Hex, _ := hex.DecodeString(source2)

    finalHex := make([]byte, 0)

    for i, b := range source1Hex {
        finalHex = append(finalHex, b^source2Hex[i])
    }

    finalHexString := hex.EncodeToString(finalHex)

    fmt.Println(finalHexString)
}