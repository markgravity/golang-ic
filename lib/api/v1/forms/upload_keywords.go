package forms

import (
	"encoding/csv"
	"errors"
	"io"
	"mime/multipart"

	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/lib/jobs"
	"github.com/markgravity/golang-ic/lib/models"

	"github.com/fatih/structs"
)

type UploadKeywordsForm struct {
	FileHeader *multipart.FileHeader `form:"file" binding:"required"`
	User       *models.User
}

func (f *UploadKeywordsForm) Save() error {
	db := database.GetDB()
	keywords, err := f.readCSVFile()
	if err != nil {
		return err
	}

	for _, k := range keywords {
		keyword := &models.Keyword{
			Text: k,
			User: f.User,
		}

		err := keyword.Save(db)
		if err != nil {
			return err
		}

		job := jobs.Crawl{}
		job.SetArgs(structs.Map(keyword))
		err = jobs.Dispatch(&job)
		if err != nil {
			return err
		}
	}

	return nil
}

func (f *UploadKeywordsForm) readCSVFile() ([]string, error) {
	if f.FileHeader.Header.Get("Content-Type") != "text/csv" {
		return nil, errors.New("file type is not supported")
	}

	file, err := f.FileHeader.Open()
	if err != nil {
		return nil, errors.New("file is not found")
	}

	reader := csv.NewReader(file)
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

	keywordLength := len(keywords)
	if keywordLength <= 0 || keywordLength > 1000 {
		return nil, errors.New("CSV file only accepts from 1 to 1000 keywords")
	}

	return keywords, nil
}
