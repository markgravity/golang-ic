package models

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KeywordStatus string

const (
	Failed     KeywordStatus = "failed"
	Pending    KeywordStatus = "pending"
	Completed  KeywordStatus = "completed"
	Processing KeywordStatus = "processing"

	InvalidKeywordStatusErr = "invalid keyword status"
)

type Keyword struct {
	Base                Base `gorm:"embedded;"`
	UserID              uuid.UUID
	Text                string `gorm:"unique;"`
	Status              KeywordStatus
	LinksCount          int
	NonAdwordLinks      sql.NullString `gorm:"type(json);null;"`
	NonAdwordLinksCount int
	AdwordLinks         sql.NullString `gorm:"type(json);null;"`
	AdwordLinksCount    int
	HtmlCode            string

	User *User `gorm:"foreignKey:UserID;"`
}

func (k *Keyword) Save(db *gorm.DB) error {
	return db.Save(k).Error
}
