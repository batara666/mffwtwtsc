package twitterscraper

import (
	"net/url"
	"strconv"
)

// SearchAccount gets tweets for a given search query, via the Twitter frontend API
func (s *Scraper) SearchAccount(query string, maxResult int, cursor string) (TwitterGlobal, error) {
	query = url.PathEscape(query)
	if maxResult > 40 {
		maxResult = 40
	}

	req, err := s.newRequest("GET", "https://twitter.com/i/api/2/search/adaptive.json")
	if err != nil {
		return TwitterGlobal{}, err
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

	var timeline TwitterGlobal
	err = s.RequestAPI(req, &timeline)
	if err != nil {
		return TwitterGlobal{}, err
	}

	return timeline, nil
}
