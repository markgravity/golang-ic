package modelsresponse

type KeywordDetailResponse struct {
	ID                  string `jsonapi:"primary,keyword_detail"`
	Text                string `jsonapi:"attr,text"`
	Status              string `jsonapi:"attr,status"`
	LinksCount          int    `jsonapi:"attr,links_count"`
	NonAdwordLinks      string `jsonapi:"attr,non_adword_links"`
	NonAdwordLinksCount int    `jsonapi:"attr,non_adword_links_count"`
	AdwordLinks         string `jsonapi:"attr,adword_links"`
	AdwordLinksCount    int    `jsonapi:"attr,adword_links_count"`
	HtmlCode            string `jsonapi:"attr,html_code"`
}
