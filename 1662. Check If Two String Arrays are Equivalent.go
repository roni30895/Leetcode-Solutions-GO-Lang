func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    var str1 string
    var str2 string
    for _,val1 := range word1 {
        str1 = str1 + val1
    }
    for _,val2 := range word2 {
        str2 = str2 + val2
    }
    if str1==str2 {
        return true
    }
    return false
}
