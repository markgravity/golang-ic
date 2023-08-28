package forms_test

import (
	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/lib/api/v1/forms"
	"github.com/markgravity/golang-ic/lib/models"
	"github.com/markgravity/golang-ic/test"
	"github.com/markgravity/golang-ic/test/fabricators"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UploadKeywords", func() {
	Describe("Save", func() {
		AfterEach(func() {
			test.CleanUpDatabase()
		})

		Context("Given VALID form", func() {
			It("returns without error", func() {
				file, fileHeader, _ := test.GetMultipartAttributesFromFile(
					"keywords.csv",
					"text/csv",
				)
				user := fabricators.FabricateUser("test@gmail.com", "12345678")
				form := forms.UploadKeywordsForm{
					File:       file,
					FileHeader: fileHeader,
					User:       user,
				}

				err := form.Save()

				Expect(err).To(BeNil())
			})

			It("stores correct keywords", func() {
				db := database.GetDB()
				file, fileHeader, _ := test.GetMultipartAttributesFromFile(
					"keywords.csv",
					"text/csv",
				)
				user := fabricators.FabricateUser("test@gmail.com", "12345678")
				form := forms.UploadKeywordsForm{
					File:       file,
					FileHeader: fileHeader,
					User:       user,
				}

				_ = form.Save()

				var keywords []models.Keyword
				db.Find(&keywords)

				Expect(keywords).To(HaveLen(2))
				Expect(keywords[0].Keyword).To(Equal("iphone 12"))
				Expect(keywords[1].Keyword).To(Equal("macbook pro"))
			})
		})
	})
})
