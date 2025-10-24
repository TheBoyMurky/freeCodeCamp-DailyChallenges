package main

import (
	"fmt"
	"regexp"
)

func stripTags(html string) string {
	r, _ := regexp.Compile(`</?([^<]+)>`)
	strippedTag := r.ReplaceAllString(html, "")
	return strippedTag
}

func main() {
	fmt.Println(stripTags("<a href=\"#\">Click here</a>"))
	fmt.Println(stripTags("<p><b>Some Text</b></p>"))
	fmt.Println(stripTags("<p><cite>The Scream</cite> by Edward Munch. Painted in 1893.</p> "))
}
