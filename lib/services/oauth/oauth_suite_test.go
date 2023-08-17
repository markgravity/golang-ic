package oauth_test

import (
	"testing"

	"github.com/markgravity/golang-ic/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOAuth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OAuth Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
