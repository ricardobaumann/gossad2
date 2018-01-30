package piwik_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/WeltN24/gossad/piwik"
)

var _ = Describe("Extractor", func() {
	var (
		params Parameters
		vpl    VisitsPageLoaderMock
		result map[string][]string
	)
	Describe("getVisitsRecommendations", func() {

		BeforeEach(func() {
			params = Parameters{LimitPercontentID: 2, Limit: 20, MaxPages: 5, PiwikBaseURL: "http://baseurl"}
			vpl = VisitsPageLoaderMock{TriggerChannel: make(chan bool, params.MaxPages)}
			result = GetVisitsRecommendations(vpl, params)
			fmt.Printf("Results: %v\n", result)
		})

		It("Should respect the limit of results per content id", func() {
			for _, v := range result {
				Expect(v).To(HaveLen(2))
			}
		})
		It("Should not include the key content id inside the values", func() {
			for k, v := range result {
				Expect(v).NotTo(ContainElement(k))
			}
		})
	})
})
