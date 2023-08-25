package forms_test

import (
	"testing"

	"github.com/markgravity/golang-ic/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestForms(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Forms Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
