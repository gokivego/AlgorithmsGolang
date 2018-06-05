/*
Length of Longest Substring
*/

package main

import ("fmt"
)

func lengthOfLongestSubstring(s string) int {
    charHash := make(map[string]int)
    start := 0
    maxLen := 0
    for idx,rune := range s {
        if valIdx,ok := charHash[string(rune)] ; ok {
            for _,rune := range s[start:valIdx] {
                delete(charHash,string(rune))
            }
            start = valIdx + 1
            charHash[string(rune)] = idx
        } else {
            charHash[string(rune)] = idx
            if maxLen < (idx - start + 1) {
                maxLen = idx - start + 1
            }
        } 
    }
    return maxLen
}

func main() {
    //s1 := "abcabcbb"
    s2 := "abcdecdfghi"
    //s2 := "tmmzuxt"
    length := lengthOfLongestSubstring(s2)
    fmt.Println("Length of Longest Substring in the string ",s2,"is :",length)
}
