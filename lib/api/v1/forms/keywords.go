package forms

import (
	"encoding/csv"
	"errors"
	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/helpers/log"
	"io"
	"mime/multipart"

	"github.com/markgravity/golang-ic/lib/jobs"
	"github.com/markgravity/golang-ic/lib/models"

	"github.com/fatih/structs"
)

type KeywordsForm struct {
	FileHeader *multipart.FileHeader `form:"file" binding:"required"`
	User       *models.User
}

func (f *KeywordsForm) Save() error {
	keywords, err := f.readKeywordsFromCSVFile()
	if err != nil {
		return err
	}

	db := database.GetDB()
	err = db.Create(&keywords).Error
	if err != nil {
		return err
	}

	for _, k := range keywords {
		job := jobs.Crawl{}
		job.SetArgs(structs.Map(k))
		err = jobs.Dispatch(&job)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *KeywordsForm) readKeywordsFromCSVFile() ([]models.Keyword, error) {
	if f.FileHeader.Header.Get("Content-Type") != "text/csv" {
		return nil, errors.New("file type is not supported")
	}

	file, err := f.FileHeader.Open()
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Errorf("unable to close file %v", err)
		}
	}(file)

	if err != nil {
		return nil, errors.New("file is not found")
	}

	reader := csv.NewReader(file)
	var keywords []models.Keyword

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		keyword := models.Keyword{
			Text: row[0],
			User: f.User,
		}
		keywords = append(keywords, keyword)
	}

	keywordLength := len(keywords)
	if keywordLength <= 0 || keywordLength > 1000 {
		return nil, errors.New("CSV file only accepts from 1 to 1000 keywords")
	}

	if err != nil {
		return nil, err
	}

	return keywords, nil
}
