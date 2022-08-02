package main

import "fmt"

func main() {
	// fmt.Println(canConstruct("a", "b"))
	// fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
	fmt.Println(canConstruct("fihjjjjei", "hjibagacbhadfaefdjaeaebgi"))
	fmt.Println(canConstruct("aabcd", "zxradbca"))
}

func canConstruct(ransomNote string, magazine string) bool {
	magRunes := []rune(magazine)
	noteRunes := []rune(ransomNote)
	for c, note := range ransomNote {
		for i, mag := range magRunes {
			if note == mag {
				magRunes[i] = '0'
				noteRunes[c] = '1'
				break
			}
		}
		if noteRunes[c] != '1' {
			return false
		}
	}
	return true
}
