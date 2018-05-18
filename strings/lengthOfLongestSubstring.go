/*
Length of Longest Substring
*/

package main

import ("fmt"
)

func lengthOfLongestSubstring(s string) int {
    charHash := make(map[string]int)
    start := -1
    maxLen := 0
    for idx,rune := range s {
        if valIdx,ok := charHash[string(rune)] ; ok {
            start = valIdx
            charHash[string(rune)] = idx
        } else {
            charHash[string(rune)] = idx
            if maxLen < (idx - start) {
                maxLen = idx - start
            }
        } 
    }
    return maxLen
}

func main() {
    //s1 := "abcabcbb"
    s2 := "abcdecdfghi"
    length := lengthOfLongestSubstring(s2)
    fmt.Println("Length of Longest Substring in the string ",s2,"is :",length)
}
