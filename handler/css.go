package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/eternal-flame-AD/badge/internal/shields"
)

func CSSSelector(selector string, url string) (*shields.B, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}
	el := doc.Find(selector)
	if el == nil {
		return nil, errors.New("element not found")
	}
	return &shields.B{
		Subject: "CSS Selector",
		Status:  el.Text(),
		Color:   "green",
	}, nil
}
