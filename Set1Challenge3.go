package main

import hex "encoding/hex"
import "fmt"
import "sort"

func main() {
    sourceString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373" +
                    "e783a393b3736"
    sourceHex, _ := hex.DecodeString(sourceString)

    // This is an array of every 'character' that may be the key used to
    // decode the source string
    decoderArray := make([]byte, 94)

    // This is the array where we'll store a 'score' for each key used
    // in the previous array
    scoreArray := make([]int, 94)
    scoreMap := make(map[int]string)

    for i, _ := range decoderArray {
        finalHex := make([]byte, 0)

        for _, v := range sourceHex {
            finalHex = append(finalHex, byte(i + 32) ^ v)
        }

        finalString := string(finalHex)
        scoreArray[i] = scoreString(finalString)
        scoreMap[i] = finalString
    }

    sortedScoreArray := make([]int, 94)
    copy(sortedScoreArray, scoreArray)

    sort.Ints(sortedScoreArray)
    sort.Sort(sort.Reverse(sort.IntSlice(sortedScoreArray)))

    highestScore := sortedScoreArray[0]
    var decoderIndex int

    for i, v := range scoreArray {
        if v == highestScore {
            decoderIndex = i
        }
    }

    fmt.Println(scoreMap[decoderIndex])
}

func scoreString(target string) int {
    score := 0

    for _, v := range target {

        switch {
        case byte(v) < 0x65:
            score = score - 1
        }

        switch string(v) {
        case "e":
            score = score + 5
        case "t":
            score = score + 5
        case "a":
            score = score + 4
        case "o":
            score = score + 4
        case "i":
            score = score + 4
        }
    }

    return score
}