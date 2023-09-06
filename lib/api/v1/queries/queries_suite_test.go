package queries_test

import (
	"testing"

	"github.com/markgravity/golang-ic/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQueries(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Queries Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
