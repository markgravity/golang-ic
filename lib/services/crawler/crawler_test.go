package crawler_test

import (
	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/lib/models"
	"github.com/markgravity/golang-ic/lib/services/crawler"
	"github.com/markgravity/golang-ic/test"
	"github.com/markgravity/golang-ic/test/fabricators"

	"github.com/go-faker/faker/v4"
	"github.com/gocolly/colly/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crawler", func() {
	AfterEach(func() {
		test.CleanUpDatabase()
	})

	Describe("#Run", func() {
		Context("given VALID keyword", func() {
			It("returns the parsing result", func() {
				cassetteName := "services/crawler/success"
				_, recorder := test.GetRecorderClient(cassetteName)
				defer func() {
					if recorder.Stop() != nil {
						Fail("Fail to stop recorder")
					}
				}()

				db := database.GetDB()
				user := fabricators.FabricateUser(faker.Email(), faker.Password())
				keyword := fabricators.FabricateKeyword("iphone 12", user)

				collector := colly.NewCollector()
				collector.WithTransport(recorder)
				service := crawler.Crawler{Keyword: keyword, Collector: collector, DB: db}
				err := service.Run()
				if err != nil {
					Fail(err.Error())
				}

				var keywords []models.Keyword
				db.Find(&keywords)

				Expect(keywords).To(HaveLen(1))
				Expect(keywords[0].Keyword).To(Equal("iphone 12"))
				Expect(keywords[0].Status).To(Equal(models.Processed))
				Expect(keywords[0].LinksCount).To(Equal(10))
				Expect(keywords[0].NonAdwordLinksCount).To(Equal(10))
				Expect(keywords[0].NonAdwordLinks).NotTo(BeNil())
				Expect(keywords[0].AdwordLinksCount).To(Equal(0))
				Expect(keywords[0].AdwordLinks.String).To(BeEmpty())
				Expect(keywords[0].HtmlCode).ToNot(BeNil())
				Expect(keywords[0].UserID).To(Equal(user.Base.ID))
			})
		})

		Context("given INVALID keyword", func() {
			It("returns the parsing result", func() {
				service := crawler.Crawler{Keyword: nil}
				err := service.Run()

				Expect(err.Error()).To(Equal("keyword is required"))
			})
		})
	})
})
