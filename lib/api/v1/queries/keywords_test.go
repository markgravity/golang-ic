package queries_test

import (
	"github.com/markgravity/golang-ic/lib/api/v1/queries"
	"github.com/markgravity/golang-ic/test"
	"github.com/markgravity/golang-ic/test/fabricators"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeywordsQuery", func() {
	Describe("Where", func() {
		AfterEach(func() {
			test.CleanUpDatabase()
		})

		Context("Given VALID params", func() {
			It("returns without error", func() {
				user := fabricators.FabricateUser("test@gmail.com", "123456")
				fabricators.FabricateKeyword("k1", user)
				fabricators.FabricateKeyword("k2", user)

				params := queries.KeywordsQueryParams{
					Offset: 0,
					Limit:  1,
				}
				query := queries.KeywordsQuery{
					User: *user,
				}
				_, err := query.Where(params)

				Expect(err).To(BeNil())
			})

			It("returns correct keywords", func() {
				user := fabricators.FabricateUser("test@gmail.com", "123456")
				fabricators.FabricateKeyword("k1", user)
				fabricators.FabricateKeyword("k2", user)

				params := queries.KeywordsQueryParams{
					Offset: 0,
					Limit:  1,
				}
				query := queries.KeywordsQuery{
					User: *user,
				}
				keywords, _ := query.Where(params)

				Expect(keywords).To(HaveLen(1))
				Expect(keywords[0].Text).To(Equal("k1"))
			})
		})

		Context("Given Offset", func() {
			It("returns correct keywords", func() {
				user := fabricators.FabricateUser("test@gmail.com", "123456")
				fabricators.FabricateKeyword("k1", user)
				fabricators.FabricateKeyword("k2", user)

				params := queries.KeywordsQueryParams{
					Offset: 1,
					Limit:  1,
				}
				query := queries.KeywordsQuery{
					User: *user,
				}
				keywords, _ := query.Where(params)

				Expect(keywords).To(HaveLen(1))
				Expect(keywords[0].Text).To(Equal("k2"))
			})
		})

		Context("Given Limit", func() {
			It("returns correct keywords", func() {
				user := fabricators.FabricateUser("test@gmail.com", "123456")
				fabricators.FabricateKeyword("k1", user)
				fabricators.FabricateKeyword("k2", user)

				params := queries.KeywordsQueryParams{
					Offset: 0,
					Limit:  2,
				}
				query := queries.KeywordsQuery{
					User: *user,
				}
				keywords, _ := query.Where(params)

				Expect(keywords).To(HaveLen(2))
			})
		})

		Context("Given Text", func() {
			It("returns correct keywords", func() {
				user := fabricators.FabricateUser("test@gmail.com", "123456")
				fabricators.FabricateKeyword("k1", user)
				fabricators.FabricateKeyword("k2", user)

				params := queries.KeywordsQueryParams{
					Offset: 0,
					Limit:  2,
					Text:   "1",
				}
				query := queries.KeywordsQuery{
					User: *user,
				}
				keywords, _ := query.Where(params)

				Expect(keywords).To(HaveLen(1))
				Expect(keywords[0].Text).To(Equal("k1"))
			})
		})

		Context("Given INCORRECT user", func() {
			It("returns empty keywords", func() {
				user1 := fabricators.FabricateUser("test@gmail.com", "123456")
				user2 := fabricators.FabricateUser("test2@gmail.com", "123456")
				fabricators.FabricateKeyword("k1", user1)
				fabricators.FabricateKeyword("k2", user1)

				params := queries.KeywordsQueryParams{
					Offset: 0,
					Limit:  1,
				}
				query := queries.KeywordsQuery{
					User: *user2,
				}
				keywords, _ := query.Where(params)

				Expect(keywords).To(BeEmpty())
			})
		})
	})

	Describe("Find", func() {
		AfterEach(func() {
			test.CleanUpDatabase()
		})

		Context("Given EXISTS keyword ID", func() {
			It("returns without error", func() {
				user := fabricators.FabricateTester()
				keyword := fabricators.FabricateKeyword("k1", user)

				query := queries.KeywordsQuery{
					User: *user,
				}
				_, err := query.Find(keyword.Base.ID.String())

				Expect(err).To(BeNil())
			})

			It("returns correct keyword", func() {
				user := fabricators.FabricateTester()
				keyword := fabricators.FabricateKeyword("k1", user)

				query := queries.KeywordsQuery{
					User: *user,
				}
				result, _ := query.Find(keyword.Base.ID.String())

				Expect(result.Base.ID).To(Equal(keyword.Base.ID))
			})
		})

		Context("Given NON-EXISTS keyword ID", func() {
			It("returns the database error", func() {
				user := fabricators.FabricateTester()

				query := queries.KeywordsQuery{
					User: *user,
				}
				_, err := query.Find("invalid")

				Expect(err).NotTo(BeNil())
			})
		})
	})
})
