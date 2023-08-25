package serializers_test

import (
	"testing"

	"github.com/markgravity/golang-ic/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSerializers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Serializers Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
