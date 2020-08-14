package test

import (
	"aries/utils"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestRegHTML(t *testing.T) {
	html := `<div class='box'><a>132123213      </a><img/><b>
hello<b></div>`
	content := utils.GetHtmlContent(html)
	log.Println("content: ", content)
}
