package serializers_test

import (
	"time"

	"github.com/markgravity/golang-ic/lib/api/v1/serializers"

	_ "github.com/go-oauth2/oauth2/v4"
	oauth2models "github.com/go-oauth2/oauth2/v4/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TokenSerializer", func() {
	Describe("#Data", func() {
		Context("given a VALID params", func() {
			It("returns serialized data", func() {
				token := oauth2models.NewToken()
				token.Access = "Access"
				token.Refresh = "Refresh"
				token.AccessExpiresIn = 7200 * time.Second

				serializer := serializers.TokenSerializer{
					Token:     token,
					TokenType: "bearer",
				}

				result := serializer.Data()

				Expect(result).NotTo(BeNil())
				Expect(result.ID).NotTo(BeEmpty())
				Expect(result.TokenType).To(Equal("bearer"))
				Expect(result.AccessToken).To(Equal("Access"))
				Expect(result.RefreshToken).To(Equal("Refresh"))
				Expect(result.ExpiresIn).To(Equal(int64(7200)))
			})
		})
	})
})
