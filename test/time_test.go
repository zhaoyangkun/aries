package test

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"testing"
	"time"
)

func TestGetNowTime(t *testing.T) {
	s := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println("s: ", s)
}
