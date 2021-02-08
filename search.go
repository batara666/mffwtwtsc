package twitterscraper

import (
	"net/url"
	"strconv"
)

// SearchAccount gets tweets for a given search query, via the Twitter frontend API
func (s *Scraper) SearchAccount(query string, maxResult int, cursor string) ([]TwAccount, error) {
	query = url.PathEscape(query)
	if maxResult > 40 {
		maxResult = 40
	}
	acc := make([]TwAccount, 0)

	req, err := s.newRequest("GET", "https://twitter.com/i/api/2/search/adaptive.json")
	if err != nil {
		return acc, err
	}

	q := req.URL.Query()
	q.Add("q", query)
	q.Add("count", strconv.Itoa(maxResult))
	q.Add("query_source", "typed_query")
	q.Add("pc", "1")
	q.Add("spelling_corrections", "1")
	if cursor != "" {
		q.Add("cursor", cursor)
	}
	q.Add("result_filter", "user")

	req.URL.RawQuery = q.Encode()

	glob := TwitterGlobal{}
	err = s.RequestAPI(req, &glob)
	if err != nil {
		return acc, err
	}

	for _, val := range glob.GlobalObjects.Users {
		acc = append(acc, val)
	}

	return acc, nil
}
