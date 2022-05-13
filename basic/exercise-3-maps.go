package tour

import "strings"

func WordCount(s string) map[string]int {
	arr := strings.Fields(s)
	wordCountMap := make(map[string]int)
	for i := range arr {
		v, ok := wordCountMap[arr[i]]
		if ok {
			wordCountMap[arr[i]] = v + 1
		} else {
			wordCountMap[arr[i]] = 1
		}
	}
	return wordCountMap
}
