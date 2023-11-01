package serializers

import (
	"github.com/markgravity/golang-ic/lib/models"
	responsemodels "github.com/markgravity/golang-ic/lib/models/response"
)

type KeywordDetailSerializer struct {
	Keyword models.Keyword
}

func (s *KeywordDetailSerializer) Data() *responsemodels.KeywordDetailResponse {
	return &responsemodels.KeywordDetailResponse{
		ID:                  s.Keyword.Base.ID.String(),
		Text:                s.Keyword.Text,
		Status:              string(s.Keyword.Status),
		LinksCount:          s.Keyword.LinksCount,
		NonAdwordLinks:      s.Keyword.NonAdwordLinks.String,
		NonAdwordLinksCount: s.Keyword.NonAdwordLinksCount,
		AdwordLinks:         s.Keyword.AdwordLinks.String,
		AdwordLinksCount:    s.Keyword.AdwordLinksCount,
		HtmlCode:            s.Keyword.HtmlCode,
	}
}
