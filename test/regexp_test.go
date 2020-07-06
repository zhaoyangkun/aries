package test

import (
	"aries/util"
	"log"
	"testing"
)

func TestRegHTML(t *testing.T) {
	html := "<div class='box'><a>132123213</a><img/><b>hello<b></div>"
	content := util.GetHtmlContent(html)
	log.Println("content: ", content)
}
