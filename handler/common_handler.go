package handler

import "strings"

func parseURL(url string) string {
	p := strings.Split(url, "/")
	return p[len(p)-1]
}
