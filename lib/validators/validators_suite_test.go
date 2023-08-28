package validators_test

import (
	"testing"

	"github.com/markgravity/golang-ic/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestValidators(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Validators Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
