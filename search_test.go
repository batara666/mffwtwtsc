package twitterscraper

import (
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGetSearchTweets(t *testing.T) {
	count := 0
	maxResult := 250
	a := New()
	p, err := a.SearchAccount("joko widodo", maxResult, "")
	log.Println(err)
	spew.Dump(p.GlobalObjects.Users)

	if count != maxResult {
		t.Errorf("Expected tweets count=%v, got: %v", maxResult, count)
	}
}
