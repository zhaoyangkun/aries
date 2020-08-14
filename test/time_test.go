package test

import (
	"strconv"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func TestGetNowTime(t *testing.T) {
	s := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println("s: ", s)
}
