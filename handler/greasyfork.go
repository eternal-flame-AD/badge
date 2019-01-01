package handler

import "github.com/eternal-flame-AD/badge/internal/shields"
import "net/url"

func GreasyForkTotalInstall(scriptname string) (*shields.B, error) {
	b, err := CSSSelector("dd.script-show-total-installs > span", "https://greasyfork.org/en/scripts/"+url.PathEscape(scriptname))
	if err != nil {
		return nil, err
	}
	b.Subject = "GreasyFork"
	b.Status = b.Status + " downloads"
	return b, nil
}

func GreasyForkVersion(scriptname string) (*shields.B, error) {
	b, err := CSSSelector("dd.script-show-version > span", "https://greasyfork.org/en/scripts/"+url.PathEscape(scriptname))
	if err != nil {
		return nil, err
	}
	b.Subject = "GreasyFork"
	return b, nil
}
