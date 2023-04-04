func partitionString(s string) int {
    if len(s)<=1{
		return len(s)
	}
	var str string
  result := []string{}
 	for _,val := range s {
		if key(str,string(val)) != -1 {
			result = append(result,str)
			str=""
		}
		str = str + string(val)
	}
	return len(result)+1
}
func key(str,char string)int{
	for key,val := range str{
		if string(val) == string(char){
			return key
		}
	}
	return -1
}
