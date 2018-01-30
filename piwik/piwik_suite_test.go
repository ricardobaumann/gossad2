package piwik_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPiwik(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Piwik Suite")
}
