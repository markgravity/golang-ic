package forms

import (
	"encoding/csv"
	"io"
	"mime/multipart"

	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/lib/models"
	"github.com/markgravity/golang-ic/lib/services/crawler"

	"github.com/gocolly/colly/v2"
)

type UploadKeywordsForm struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	User       *models.User
}

func (f *UploadKeywordsForm) Save() error {
	db := database.GetDB()
	keywords, err := f.readCSVFile()
	if err != nil {
		return err
	}

	// TODO: Replace to use a worker in https://github.com/markgravity/golang-ic/issues/44
	for _, k := range keywords {
		keyword := &models.Keyword{
			Keyword: k,
			User:    f.User,
		}

		err := keyword.Save(db)
		if err != nil {
			return err
		}

		collector := colly.NewCollector()
		keywordCrawler := crawler.Crawler{
			DB:        db,
			Keyword:   keyword,
			Collector: collector,
		}
		err = keywordCrawler.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *UploadKeywordsForm) readCSVFile() ([]string, error) {
	reader := csv.NewReader(f.File)
	var keywords []string

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		keywords = append(keywords, row[0])
	}

	return keywords, nil
}
