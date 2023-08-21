package helpers

import "net/url"

func GenerateURLParams(params map[string]interface{}) url.Values {
	values := url.Values{}

	for key, value := range params {
		values.Add(key, value.(string))
	}

	return values
}
