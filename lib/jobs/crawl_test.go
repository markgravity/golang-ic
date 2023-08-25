package jobs_test

import (
	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/lib/jobs"
	"github.com/markgravity/golang-ic/lib/models"
	"github.com/markgravity/golang-ic/test"
	"github.com/markgravity/golang-ic/test/fabricators"

	"github.com/fatih/structs"
	"github.com/go-faker/faker/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crawl", func() {
	AfterEach(func() {
		test.CleanUpDatabase()
	})

	Describe("#Handle", func() {
		Context("given VALID args", func() {
			It("returns without error", func() {
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword := fabricators.FabricateKeyword("iphone 12", user)

				job := jobs.Crawl{}
				job.SetArgs(structs.Map(keyword))
				err := job.Handle()

				Expect(err).To(BeNil())
			})

			It("crawls keyword data", func() {
				db := database.GetDB()
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword := fabricators.FabricateKeyword("iphone 12", user)

				job := jobs.Crawl{}
				job.SetArgs(structs.Map(keyword))
				_ = job.Handle()

				db.First(&keyword)
				Expect(keyword.Status).To(Equal(models.Processed))
			})
		})

		Context("given INVALID args", func() {
			It("returns the error", func() {
				job := jobs.Crawl{}
				job.SetArgs(map[string]interface{}{})
				err := job.Handle()

				Expect(err.Error()).To(Equal("keyword is required"))
			})
		})
	})
})
