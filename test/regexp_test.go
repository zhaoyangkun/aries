package test

import (
	"aries/utils"
	"log"
	"testing"
)

func TestRegHTML(t *testing.T) {
	html := `<div class='box'><a>132123213      </a><img/><b>
hello<b></div>`
	content := utils.GetHtmlContent(html)
	log.Println("content: ", content)
}
