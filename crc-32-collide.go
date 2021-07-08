package main

import (
    "hash/crc32"
    "fmt"
    "sync"
)

// https://stackoverflow.com/questions/19249588/go-programming-generating-combinations
// https://play.golang.org/p/0bWDCibSUJ
func AddLetter(c chan string, combo string, alphabet string, length int) {
    // Check if we reached the length limit
    // If so, we just return without adding anything
    if length <= 0 {
        return
    }

    var newCombo string
    for _, ch := range alphabet {
        newCombo = combo + string(ch)
        c <- newCombo
        AddLetter(c, newCombo, alphabet, length-1)
    }
}

func worker(wChan chan string, target uint32) {
    for tString := range wChan {
        if crc32.ChecksumIEEE([]byte(tString)) == target {
            fmt.Println("Collision found:", tString)
        }
    }
}


func main() {
    // target CRC-32
    // & 0xffffffff is to convert to uint
    // required since old python versions allowed negative values to be produced
    // hence its needed if you want to find a collision for a "negative" crc hash value
    const target = -432570933 & 0xffffffff
    // max string length
    maxLen := 5
    // max number of workers - More workers doesn't mean faster bruteforce
    numWorkers := 8
    // python printable alphabet excluding \t\n\r\x0b\x0c
    alphabet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}"

    var wg sync.WaitGroup
    wChan := make(chan string, 1000)
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            worker(wChan, target)
            wg.Done()
        }()
    }

    // generate all possible combinations
    AddLetter(wChan, "", alphabet, maxLen)

    close(wChan)
    wg.Wait()
}
