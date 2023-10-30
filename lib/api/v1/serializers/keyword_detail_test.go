package serializers_test

import (
	"github.com/markgravity/golang-ic/lib/api/v1/serializers"
	"github.com/markgravity/golang-ic/test"
	"github.com/markgravity/golang-ic/test/fabricators"

	_ "github.com/go-oauth2/oauth2/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordDetailSerializer", func() {
	Describe("#Data", func() {
		AfterEach(func() {
			test.CleanUpDatabase()
		})

		It("returns serialized data", func() {
			user := fabricators.FabricateTester()
			keyword := fabricators.FabricateKeyword("test", user)

			serializer := serializers.KeywordDetailSerializer{
				Keyword: *keyword,
			}

			result := serializer.Data()

			Expect(result.ID).To(Equal(keyword.Base.ID.String()))
			Expect(result.Text).To(Equal(keyword.Text))
			Expect(result.Status).To(Equal(string(keyword.Status)))
			Expect(result.LinksCount).To(Equal(keyword.LinksCount))
			Expect(result.NonAdwordLinks).To(Equal(keyword.NonAdwordLinks.String))
			Expect(result.NonAdwordLinksCount).To(Equal(keyword.NonAdwordLinksCount))
			Expect(result.AdwordLinks).To(Equal(keyword.AdwordLinks.String))
			Expect(result.AdwordLinksCount).To(Equal(keyword.AdwordLinksCount))
			Expect(result.HtmlCode).To(Equal(keyword.HtmlCode))
		})
	})
})
