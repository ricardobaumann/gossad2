package piwik

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//VisitsPageLoader loads visits from a specific piwik visitors log url
type VisitsPageLoader interface {
	load(URL string, pageResultsChannel chan []Visit)
}

type visitsPageLoaderImpl struct {
	controlChannel chan bool
}

func (loader visitsPageLoaderImpl) load(URL string, pageResultsChannel chan []Visit) {
	<-loader.controlChannel
	timeout := time.Duration(600 * time.Second)
	httpClient := http.Client{
		Timeout: timeout,
	}
	//url := fmt.Sprintf("%s&filter_limit=%d&filter_offset=%d", baseUrl, limit, offset)
	println("Loading url")
	resp, err := httpClient.Get(URL)
	if err == nil {
		defer resp.Body.Close()
		bodyContent, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Printf("URL loaded: %v\n", resp.Status)
			var visits []Visit
			json.Unmarshal(bodyContent, &visits)
			pageResultsChannel <- visits
		} else {
			println(string(err.Error()))
			pageResultsChannel <- []Visit{}
		}
	} else {
		println(string(err.Error()))
		pageResultsChannel <- []Visit{}
	}

	loader.controlChannel <- true
}
