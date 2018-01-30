package piwik

import (
	"os"
	"strconv"
)

//Parameters for the recommendations fetching
type Parameters struct {
	LimitPercontentID int
	Limit             int
	MaxPages          int
	PiwikBaseURL      string
}

//GetRecommendations load results from piwik and calculate the recommendations
func GetRecommendations() map[string][]string {
	piwikBaseURL := "https://piwik.up.welt.de/index.php?module=API&method=Live.getLastVisitsDetails&format=JSON&showColumns=actionDetails,deviceType" +
		"&idSite=1&period=day&date=today&expanded=1&filter_sort_column=lastActionTimestamp&filter_sort_order=desc" +
		"&token_auth=" + os.Getenv("piwikToken")

	limit, _ := strconv.Atoi(os.Getenv("limitPerPage"))
	maxPages, _ := strconv.Atoi(os.Getenv("maxPages"))

	params := Parameters{10, limit, maxPages, piwikBaseURL}

	throughput, _ := strconv.Atoi(os.Getenv("throughput"))

	controlChannel := make(chan bool, throughput)
	for i := 0; i < throughput; i++ {
		go putOnChannel(true, controlChannel)
	}
	return GetVisitsRecommendations(visitsPageLoaderImpl{controlChannel}, params)
}

func putOnChannel(value bool, channel chan bool) {
	channel <- value
}
