package jobs

import (
	"encoding/json"

	"github.com/markgravity/golang-ic/database"
	"github.com/markgravity/golang-ic/lib/models"
	"github.com/markgravity/golang-ic/lib/services/crawler"

	"github.com/gocolly/colly/v2"
)

type Crawl struct {
	args map[string]interface{}
}

func (c *Crawl) SetArgs(args map[string]interface{}) {
	c.args = args
}

func (c *Crawl) GetArgs() map[string]interface{} {
	return c.args
}

func (c *Crawl) GetName() string {
	return "Crawl"
}

func (c *Crawl) Handle() error {
	jsonBody, err := json.Marshal(c.args)
	if err != nil {
		return err
	}

	keyword := models.Keyword{}
	if err := json.Unmarshal(jsonBody, &keyword); err != nil {
		return err
	}

	db := database.GetDB()
	collector := colly.NewCollector()
	keywordCrawler := crawler.Crawler{
		DB:        db,
		Keyword:   &keyword,
		Collector: collector,
	}

	return keywordCrawler.Run()
}
