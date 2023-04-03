func key(str,char string)int{
	for key,val := range str{
		if string(val) == string(char){
			return key
		}
	}
	return -1
}
func lengthOfLongestSubstring(s string) int {
	if len(s)<=1{
		return len(s)
	}
	var count int
	var str string
 
	for _,val := range s {
		if key(str,string(val)) != -1 {
			str = str[key(str,string(val)) + 1:]
		}
		str = str + string(val)
		if len(str)>count{
			count = len(str)
		}
	}
	return count
}
  
