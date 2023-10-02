package serializers

import (
	"github.com/markgravity/golang-ic/lib/models"
	responsemodels "github.com/markgravity/golang-ic/lib/models/response"
)

type KeywordsSerializer struct {
	Keywords []models.Keyword
}

func (s *KeywordsSerializer) Data() (responses []*responsemodels.KeywordResponse) {
	for _, keyword := range s.Keywords {
		responses = append(responses, &responsemodels.KeywordResponse{
			ID:     keyword.Base.ID.String(),
			Text:   keyword.Text,
			Status: string(keyword.Status),
		})
	}
	return responses
}
