func restoreString(s string, indices []int) string {
    result:=make([]byte,len(s))
    key:=0
    if key<len(indices){
        for _,value := range indices{
            result[value]=s[key]
            key++
        }
    }
    return string(result)
}
