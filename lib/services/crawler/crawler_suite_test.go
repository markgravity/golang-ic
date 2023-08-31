package crawler_test

import (
	"testing"

	"github.com/markgravity/golang-ic/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCrawler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services/Crawler Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
