func truncateSentence(s string, k int) string {
    str:= strings.Split(s," ")
    words:=[]string{}

    for i:=0;i<k;i++{
        words = append(words,string(str[i]))
    }

    return strings.Join(words," ")
}
