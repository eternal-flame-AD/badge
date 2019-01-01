package shields

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type B struct {
	Subject        string
	Status         string
	Color          string
	Style          string
	Logo           *Logo
	Link           *Link
	ColorLOverride string
}
type Logo struct {
	URLOrName string
	Color     string
	Width     int
}
type Link struct {
	L string
	R string
}

func escape(s string) string {
	return url.PathEscape(strings.Replace(strings.Replace(s, "_", "__", -1), "-", "--", -1))
}

func URL(badge B) string {
	u := fmt.Sprintf("https://img.shields.io/badge/%s-%s-%s.svg", escape(badge.Subject), escape(badge.Status), escape(badge.Color))
	query := make(url.Values)
	if badge.Style != "" {
		query.Set("style", badge.Style)
	}
	if badge.ColorLOverride != "" {
		query.Set("colorA", badge.ColorLOverride)
	}
	if badge.Logo != nil {
		query.Set("logo", badge.Logo.URLOrName)
		if badge.Logo.Color != "" {
			query.Set("logoColor", badge.Logo.Color)
		}
		if badge.Logo.Width > 0 {
			query.Set("logoWidth", strconv.Itoa(badge.Logo.Width))
		}
	}
	if badge.Link != nil {
		query.Add("link", badge.Link.L)
		query.Add("link", badge.Link.R)
	}
	if qStr := query.Encode(); qStr != "" {
		u += "?" + qStr
	}
	return u
}
