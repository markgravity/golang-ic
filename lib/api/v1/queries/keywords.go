package queries

import (
	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/lib/models"
)

type KeywordsQueryParams struct {
	Offset int
	Limit  int
}

type KeywordsQuery struct {
	User models.User
}

func (q *KeywordsQuery) Where(queryParams KeywordsQueryParams) ([]models.Keyword, error) {
	db := database.GetDB()

	var keywords []models.Keyword
	err := db.Where("user_id = ?", q.User.Base.ID.String()).
		Offset(queryParams.Offset).
		Limit(queryParams.Limit).
		Find(&keywords).Error
	if err != nil {
		return nil, err
	}

	return keywords, nil
}
