package main

import (
    "code.google.com/p/go-tour/wc"
     "strings"
)

func WordCount(s string) map[string]int {
    countmap:= make(map[string] int)
    sfields:= strings.Fields(s)
    for x:= range sfields{
       ( countmap[sfields[x]])++
    }
    
    return countmap
}

func main() {
    wc.Test(WordCount)
}
