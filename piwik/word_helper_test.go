package piwik_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/WeltN24/gossad/piwik"
)

var _ = Describe("WordHelper", func() {

	Describe("GetTopOcurrences", func() {
		It("Should count the word ocurrences on the list and return the top N ocurrences", func() {
			words := []string{"1", "2", "2", "3", "3", "3", "4", "4", "4", "4"}
			result1 := GetTopOcurrences(words, 2)
			Expect(result1).To(Equal([]string{"4", "3"}))

			result2 := GetTopOcurrences(words, 3)
			Expect(result2).To(Equal([]string{"4", "3", "2"}))
		})
	})

})
