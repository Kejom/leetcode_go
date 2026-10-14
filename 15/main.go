package main

import "fmt"

func main() {
	var test1 = []string{"flower", "flow", "flight"}
	var test2 = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(test1))
	fmt.Println(longestCommonPrefix(test2))
}

func longestCommonPrefix(strs []string) string {
	var minLen int = 200
	for _, v := range strs {
		var strLen = len(v)
		if strLen < minLen {
			minLen = strLen
		}
	}

	result := ""

	for i := 0; i < minLen; i++ {
		curLetter := string(strs[0][i])

		for j := 1; j < len(strs); j++ {
			if string(strs[j][i]) != curLetter {
				return result
			}
		}

		result += curLetter
	}

	return result
}
