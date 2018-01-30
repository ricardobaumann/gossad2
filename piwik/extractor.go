package piwik

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

//GetVisitsRecommendations load recommendations from piwik visitors log
func GetVisitsRecommendations(vpl VisitsPageLoader, params Parameters) map[string][]string {
	println("Loading recommendations")
	piwikBaseURL := params.PiwikBaseURL

	limitPercontentID := params.LimitPercontentID
	limit := params.Limit
	maxPages := params.MaxPages

	fmt.Printf("Max pages %v\n", maxPages)
	fmt.Printf("Limit Per Content Id %v\n", limitPercontentID)

	start := time.Now()

	reg, _ := regexp.Compile("[^0-9]+")

	pageResultsChannel := make(chan []Visit)

	for page := 0; page < maxPages; page++ {
		//url := fmt.Sprintf("%s&filter_limit=%d&filter_offset=%d", baseUrl, limit, offset)
		go vpl.load(fmt.Sprintf("%s&filter_limit=%d&filter_offset=%d", piwikBaseURL, limit, (page*limit)), pageResultsChannel)
	}
	var relationsMap = make(map[string][]string)
	for page := 0; page < maxPages; page++ {
		visits := <-pageResultsChannel
		for _, visit := range visits {
			var contentIDList []string
			for _, detail := range visit.ActionDetails {
				parts := strings.Split(detail.URL, "/")
				if len(parts) > 2 {
					contentID := reg.ReplaceAllString(parts[len(parts)-2], "")
					if contentID != "" {
						contentIDList = append(contentIDList, contentID)
					}
				}
			}
			combineKeys(contentIDList, &relationsMap)
		}
	}

	resultsToBeSaved := make(map[string][]string)
	for k, v := range relationsMap {
		amount := len(v)
		if amount > 0 {
			ordered := GetTopOcurrences(v, limitPercontentID)
			if len(ordered) > 0 {
				resultsToBeSaved[k] = ordered
			}
		}

	}
	t := time.Now()
	elapsed := t.UnixNano() - start.UnixNano()
	fmt.Printf("Total time elapsed: %v\n minutes\n", (elapsed / 1000000000 / 60))
	return resultsToBeSaved
}

func combineKeys(contentIDList []string, relationsMap *map[string][]string) {
	for _, id := range contentIDList {
		if id != "" {
			_, ok := (*relationsMap)[id]
			if !ok {
				(*relationsMap)[id] = []string{}
			}
			newList := (*relationsMap)[id]
			for _, internalId := range contentIDList {
				if id != internalId && internalId != "" {
					newList = append(newList, internalId)
				}
			}
			(*relationsMap)[id] = newList
		}

	}
}
