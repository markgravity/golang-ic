package modelsresponse

type KeywordResponse struct {
	ID     string `jsonapi:"primary,keyword"`
	Text   string `jsonapi:"attr,text"`
	Status string `jsonapi:"attr,status"`
}
