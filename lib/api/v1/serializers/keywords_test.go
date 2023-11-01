package serializers_test

import (
	"github.com/markgravity/golang-ic/lib/api/v1/serializers"
	"github.com/markgravity/golang-ic/lib/models"
	"github.com/markgravity/golang-ic/test/fabricators"

	_ "github.com/go-oauth2/oauth2/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordsSerializer", func() {
	Describe("#Data", func() {
		Context("given a VALID params", func() {
			It("returns serialized data", func() {
				user := fabricators.FabricateTester()
				keyword := fabricators.FabricateKeyword("test", user)

				serializer := serializers.KeywordsSerializer{
					Keywords: []models.Keyword{*keyword},
				}

				result := serializer.Data()

				Expect(result).NotTo(BeNil())
				Expect(result).To(HaveLen(1))
				Expect(result[0].ID).To(Equal(keyword.Base.ID.String()))
				Expect(result[0].Text).To(Equal(keyword.Text))
				Expect(result[0].Status).To(Equal(string(keyword.Status)))
			})
		})
	})
})
