package main

import "strings"

// Challenge: return the index of the first occurrence of needle in haystack,
// or -1 if needle is not part of haystack.
func main(){
	strStr("aaaaa", "bba")
}

func strStr(haystack string, needle string) int {
	// Case in which substring does not exist.
	//if strings.Contains(haystack, needle) == false {
	//	return -1
	//}else{
	//	strings.Index()
	//}
	result := strings.Index(haystack,needle)
	return result

}