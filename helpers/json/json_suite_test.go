package jsonhelpers_test

import (
	"testing"

	"github.com/markgravity/golang-ic/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJSON(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "JSON Helper Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
