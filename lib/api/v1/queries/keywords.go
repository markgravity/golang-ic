package queries

import (
	"strings"

	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/lib/models"
)

type KeywordsQueryParams struct {
	Offset int    `form:"offset" binding:"numeric"`
	Limit  int    `form:"limit" binding:"required,numeric"`
	Text   string `form:"text"`
}

type KeywordsQuery struct {
	User models.User
}

func (q *KeywordsQuery) Where(queryParams KeywordsQueryParams) ([]models.Keyword, error) {
	db := database.GetDB()

	var keywords []models.Keyword
	query := db.Where("user_id = ?", q.User.Base.ID.String()).
		Offset(queryParams.Offset).
		Limit(queryParams.Limit)

	if queryParams.Text != "" {
		query.Where("LOWER(text) LIKE ?", "%"+strings.ToLower(queryParams.Text)+"%")
	}

	err := query.Find(&keywords).Error
	if err != nil {
		return nil, err
	}

	return keywords, nil
}

func (q *KeywordsQuery) Find(id string) (*models.Keyword, error) {
	db := database.GetDB()

	var keyword models.Keyword
	err := db.First(&keyword, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &keyword, nil
}
