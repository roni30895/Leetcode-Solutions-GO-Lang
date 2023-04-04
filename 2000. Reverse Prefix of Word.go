func reversePrefix(word string, ch byte) string {
    str:=""
    for key,value := range word {
        if string(value)==string(ch) {
            str= word[:key+1]
            break
        }
    }
    L:= []string{}
    for i:=len(str)-1;i>=0;i-- {
        L = append (L, string(str[i]))
    }
    return strings.Join(L,"")+word[len(str):]
}
