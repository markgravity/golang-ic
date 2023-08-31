package jobs_test

import (
	"testing"

	"github.com/markgravity/golang-ic/test"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJobs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Jobs Suite")
}

var _ = BeforeSuite(func() {
	test.SetupTestEnvironment()
})
