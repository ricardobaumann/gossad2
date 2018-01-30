package piwik

//VisitsPageLoaderMock is a mock for the visitsPageLoader
type VisitsPageLoaderMock struct {
	TriggerChannel chan bool
}

func (loader VisitsPageLoaderMock) load(URL string, pageResultsChannel chan []Visit) {
	//fixtures for extractor_test spec. cab be moved to a file or param if needed
	visits := []Visit{Visit{ActionDetails: []ActionDetail{
		ActionDetail{URL: "ttps://www.welt.de/wirtschaft/article172596460/Emirates-sichert-Zukunft-des-Airbus-A380.html"},
		ActionDetail{URL: "ttps://www.welt.de/wirtschaft/article172596460/Emirates-sichert-Zukunft-des-Airbus-A380.html"},
		ActionDetail{URL: "ttps://www.welt.de/wirtschaft/article172596461/Emirates-sichert-Zukunft-des-Airbus-A380.html"},
		ActionDetail{URL: "ttps://www.welt.de/wirtschaft/article172596462/Emirates-sichert-Zukunft-des-Airbus-A380.html"},
		ActionDetail{URL: "ttps://www.welt.de/wirtschaft/article172596462/Emirates-sichert-Zukunft-des-Airbus-A380.html"},
		ActionDetail{URL: "ttps://www.welt.de/wirtschaft/article172596460/Emirates-sichert-Zukunft-des-Airbus-A380.html"},
	}}}
	pageResultsChannel <- visits
	loader.TriggerChannel <- true
}
