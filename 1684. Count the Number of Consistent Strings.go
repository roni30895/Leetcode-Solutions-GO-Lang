func countConsistentStrings(allowed string, words []string) int {
    var count int 
    loop:
        for _,value := range words {
            for _,char := range value{
                if !strings.Contains(allowed,string(char)){
                    continue loop
                }
            }
            count++
        }
    return count    
}
