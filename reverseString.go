package main

import (
        "fmt"
)

func reverseString(str string) string {
        runes := []rune(str)
        m := len(runes)
        for i := 0; i < m/2; i++ {
                temp := runes[i]
                runes[i],runes[m-1-i] = runes[m-1-i],temp
        }
        strFinal := string(runes)
        fmt.Println(strFinal)
        return strFinal
}


