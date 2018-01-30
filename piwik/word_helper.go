package piwik

import "sort"

func wordCount(words []string) map[string]int {
	wordFreq := make(map[string]int)
	for _, word := range words {
		_, ok := wordFreq[word]
		if ok == true {
			wordFreq[word]++
		} else {
			wordFreq[word] = 1
		}
	}
	return wordFreq
}

//GetTopOcurrences Return the top 'limit' ocurrences words in the input list
func GetTopOcurrences(v []string, limit int) []string {
	countMap := wordCount(v)
	//fmt.Printf("Count map %v\n", countMap)
	countIndexMap := make(map[int][]string)
	var resultList []string
	var countList []int
	for k, v := range countMap {

		_, ok := countIndexMap[v]
		if !ok {
			countIndexMap[v] = []string{}
		}
		countIndexMap[v] = append(countIndexMap[v], k)
		countList = append(countList, v)

	}
	sort.Sort(sort.Reverse(sort.IntSlice(countList)))
	for _, countValue := range countList {

		for _, internalId := range countIndexMap[countValue] {
			resultList = append(resultList, internalId)
		}
	}
	var length = len(resultList)
	var topResults = resultList
	if length > limit {
		topResults = resultList[0:limit]
	}

	return topResults
}
