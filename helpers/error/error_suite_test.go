package errorhelpers_test

import (
	"testing"

	"github.com/markgravity/golang-ic/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestError(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Error Helper Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
