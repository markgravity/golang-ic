package crawler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/url"
	"strings"

	"github.com/markgravity/golang-ic/helpers/log"
	"github.com/markgravity/golang-ic/lib/models"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)

type Crawler struct {
	DB        *gorm.DB
	Keyword   *models.Keyword
	Collector *colly.Collector

	parsingResult *ParsingResult
}

type ParsingResult struct {
	HTMLCode        string
	NonAdwordLinks  []string
	AdwordLinks     []string
	ShopAdwordLinks []string
}

var selectors = map[string][]string{
	"nonAds":      {"#search .yuRUbf > a", ".ZINbbc.xpd .kCrYT > a"},
	"mobileLinks": {".ezO2md a.fuLhoc.ZWRArf"},
	"mobileAds":   {"span.dloBPe"},
}

const urlPattern = "https://www.google.com/search?q=%s"

func (c *Crawler) Run() error {
	if c.Keyword == nil || c.Keyword.Base.ID.ID() == 0 {
		return errors.New("keyword is required")
	}

	c.Keyword.Status = models.Processing
	err := c.Keyword.Save(c.DB)
	if err != nil {
		c.Keyword.Status = models.Failed
		_ = c.Keyword.Save(c.DB)
		return err
	}

	parsingResult := ParsingResult{}
	if c.Collector == nil {
		c.Collector = colly.NewCollector()
	}

	extensions.RandomUserAgent(c.Collector)

	c.Collector.OnRequest(func(r *colly.Request) {
		log.Infof("Visiting: %v user-agent: %v", r.URL, r.Headers.Get("User-Agent"))
	})

	c.Collector.OnResponse(func(r *colly.Response) {
		parsingResult.HTMLCode = string(r.Body[:])
	})

	c.Collector.OnError(func(r *colly.Response, err error) {
		log.Errorf("Request URL: %v failed with error: %v", r.Request.URL, err)
		c.Keyword.Status = models.Failed
		_ = c.Keyword.Save(c.DB)
	})

	for _, pattern := range selectors["nonAds"] {
		c.Collector.OnHTML(pattern, func(e *colly.HTMLElement) {
			link := formatLink(e.Attr("href"))
			parsingResult.NonAdwordLinks = append(parsingResult.NonAdwordLinks, link)
		})
	}

	c.Collector.OnHTML(selectors["mobileLinks"][0], func(e *colly.HTMLElement) {
		link := formatLink(e.Attr("href"))
		if len(e.DOM.Find(selectors["mobileAds"][0]).Nodes) > 0 {
			parsingResult.AdwordLinks = append(parsingResult.AdwordLinks, link)
		} else {
			parsingResult.NonAdwordLinks = append(parsingResult.NonAdwordLinks, link)
		}
	})

	c.Collector.OnScraped(func(r *colly.Response) {
		c.parsingResult = &parsingResult
		err := c.save()
		if err != nil {
			log.Error("Error when saving to DB:", err)

			c.Keyword.Status = models.Failed
			err = c.Keyword.Save(c.DB)
			if err != nil {
				log.Error("Error, cannot update status to Failed:", err)
			}
		}
	})

	requestUrl := fmt.Sprintf(urlPattern, url.QueryEscape(c.Keyword.Text))
	return c.Collector.Visit(requestUrl)
}

func (c *Crawler) save() error {
	nonAdwordLinks, err := json.Marshal(c.parsingResult.NonAdwordLinks)
	if err != nil {
		return err
	}

	adwordLinks, err := json.Marshal(c.parsingResult.AdwordLinks)
	if err != nil {
		return err
	}

	nonAdwordLinksCount := len(c.parsingResult.NonAdwordLinks)
	adwordLinksCount := len(c.parsingResult.AdwordLinks)
	shopAdwordLinksCount := len(c.parsingResult.ShopAdwordLinks)
	totalCount := nonAdwordLinksCount + adwordLinksCount + shopAdwordLinksCount

	keyword := c.Keyword
	keyword.NonAdwordLinksCount = nonAdwordLinksCount
	keyword.NonAdwordLinks = sql.NullString{String: string(nonAdwordLinks)}
	keyword.AdwordLinksCount = adwordLinksCount
	keyword.AdwordLinks = sql.NullString{String: string(adwordLinks)}
	keyword.LinksCount = totalCount
	keyword.HtmlCode = c.parsingResult.HTMLCode

	keyword.Status = models.Completed

	return keyword.Save(c.DB)
}

func formatLink(link string) string {
	if strings.HasPrefix(link, "http") {
		return link
	}

	return "https://www.google.com" + link
}
