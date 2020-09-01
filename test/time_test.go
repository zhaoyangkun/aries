package test

import (
	"aries/config/setting"
	"log"
	"strconv"
	"testing"
	"time"
)

func TestGetNowTime(t *testing.T) {
	s := strconv.FormatInt(time.Now().Unix(), 10)
	log.Println("s: ", s)
}

func TestTimeDuration(t *testing.T) {
	maxAge := time.Duration(setting.Config.Logger.MaxAge) * 24 * time.Hour
	log.Println("duration: ", time.Duration(setting.Config.Logger.MaxAge))
	log.Println("maxAge: ", maxAge)
}
