package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type structMap struct {
	Value string
	Key   int
}

func Top10(s string) []string {
	uniqueMap := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		uniqueMap[word]++
	}
	sortedStruct := make([]structMap, 0, len(uniqueMap))
	for value, key := range uniqueMap {
		sortedStruct = append(sortedStruct, structMap{value, key})
	}

	sort.Slice(sortedStruct, func(i, j int) bool {
		freqI := sortedStruct[i].Key
		freqJ := sortedStruct[j].Key
		if freqI == freqJ {
			return sortedStruct[i].Value < sortedStruct[j].Value
		}
		return freqI > freqJ
	})

	res := make([]string, 0, 10)
	for j, uniqueValue := range sortedStruct {
		if j >= 10 {
			break
		}
		res = append(res, uniqueValue.Value)
	}
	return res
}
