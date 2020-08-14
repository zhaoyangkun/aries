package test

import (
	"aries/utils"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestGetFileNameOnly(t *testing.T) {
	log.Println("fileName: ", utils.GetFileNameOnly("/a/b/c.md"))
}
