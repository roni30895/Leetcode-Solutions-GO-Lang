func findTheLongestBalancedSubstring(s string) int {
    substring := []string{s[:],}

    for i:=0;i<len(s);i++{
        for j:=i+1;j<=len(s);j++{
            substring=append(substring,s[i:j])
        }
    }
    max_length:=0
       
    for _,value := range substring{
        mid:=len(value)/2
        count:= strings.Count(value,"0") + strings.Count(value,"1")
        if count % 2==0 && count >= max_length && value[:mid]==strings.Repeat("0",mid) && value[mid:]==strings.Repeat("1",mid) {
            max_length = count
            
        }
    }
    return max_length    
}
