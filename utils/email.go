package utils

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func GetEmailHTML(theme string, siteUrl string, receiver string, content string) string {
	return `<div
      style="
        box-shadow: 0 1px 4px rgba(0,0,0,.12);
        border-radius: 10px;
        color: #111;
        font-size: 12px;
        width: 95%;
        font-family: 微软雅黑, arial;
        margin: 10px auto;
        padding-bottom: 10px;
      "
    >
      <div class="adM"></div>
      <div
        style="
          width: 100%;
          background: #c2dbf6;
          min-height: 60px;
          color: white;
          border-radius: 6px 6px 0 0;
        "
      >
        <span
          style="
            line-height: 60px;
            min-height: 60px;
            margin-left: 30px;
            font-size: 16px;
          "
          >「<a
            style="color: #00a1ff; font-weight: 600; text-decoration: none;"
            href="` + siteUrl + `"
            target="_blank"
            rel="external nofollow"
            >Aries</a
          >」` + theme + `</span
        >
      </div>
      <div style="margin: 0px auto; width: 90%; font-size: 14px;">
        <p style="margin: 16px auto;">
          <span style="font-weight: bold;"
            >` + receiver + `</span
          >, 您好!
        </p>
        <p
          style="
            border-bottom: #ddd 1px solid;
            border-left: #ddd 1px solid;
            padding-bottom: 20px;
            background-color: #eff5fb;
            margin: 16px 0px;
            padding-left: 20px;
            padding-right: 20px;
            border-top: #ddd 1px solid;
            border-right: #ddd 1px solid;
            padding-top: 20px;
          "
        >
          ` + content + `
        </p>
        <p style="color: #a8979a; margin: 16px auto;">(邮件由系统自动发出，请勿回复，谢谢！)</p>
      </div>
    </div>`
}

func GetForgetPwdEmailHTML(userName string, code string) string {
	return `<div
      style="
        box-shadow: 0 1px 4px rgba(0,0,0,.12);
        border-radius: 10px;
        color: #111;
        font-size: 12px;
        width: 95%;
        font-family: 微软雅黑, arial;
        margin: 10px auto;
        padding-bottom: 10px;
      "
    >
      <div class="adM"></div>
      <div
        style="
          width: 100%;
          background: #c2dbf6;
          min-height: 60px;
          color: white;
          border-radius: 6px 6px 0 0;
        "
      >
        <span
          style="
            line-height: 60px;
            min-height: 60px;
            margin-left: 30px;
            font-size: 16px;
          "
          >「<a
            style="color: #00a1ff; font-weight: 600; text-decoration: none;"
            href=""
            target="_blank"
            rel="external nofollow"
            target="_blank"
            rel="external nofollow"
            target="_blank"
            >Aries</a
          >」忘记密码验证</span
        >
      </div>
      <div style="margin: 0px auto; width: 90%; font-size: 14px;">
        <p style="margin: 16px auto;">
          <span style="font-weight: bold;"
            >` + userName + `</span
          >, 您好!
        </p>
        <p style="margin: 16px auto;">
          您的验证码为：<a
            style="color: #00bbff; text-decoration: none;"
            href=""
            target="_blank"
            rel="external nofollow"
            target="_blank"
            >` + code + `</a
          >，15 分钟内有效
        </p>
        <p style="color: #a8979a; margin: 16px auto;">(邮件由系统自动发出，请勿回复，谢谢！)</p>
      </div>
    </div>`
}

func GetCommentEmailHTML(siteName, siteURL, theme, to, from, replierURL, title, titleURL, content string) string {
	return `<div style="
        box-shadow: 0 1px 4px rgba(0,0,0,.12);
        border-radius: 10px;
        color: #111;
        font-size: 12px;
        width: 95%;
        font-family: 微软雅黑, arial;
        margin: 10px auto;
        padding-bottom: 10px;
      ">
    <div class="adM"></div>
    <div style="
          width: 100%;
          background: #c2dbf6;
          min-height: 60px;
          color: white;
          border-radius: 6px 6px 0 0;
        ">
        <span style="
            line-height: 60px;
            min-height: 60px;
            margin-left: 30px;
            font-size: 16px;
          ">「<a style="color: #00a1ff; font-weight: 600; text-decoration: none;" href="` + siteURL + `" target="_blank"
                rel="external nofollow" target="_blank" rel="external nofollow" target="_blank">` + siteName + `</a>」` + theme + `</span>
    </div>
    <div style="margin: 0px auto; width: 90%; font-size: 14px;">
        <p style="margin: 16px auto;">
            <span style="font-weight: bold;">` + to + `</span>, 您好!
        </p>
        <p style="margin: 16px auto;">
            <a style="color: #00bbff; text-decoration: none;" href="` + replierURL + `" target="_blank" rel="external nofollow"
                target="_blank">` + from + `</a>在<a style="color: #00bbff; text-decoration: none;" href="` + titleURL + `" target="_blank"
                rel="external nofollow" target="_blank">` + title + `</a>下发表了评论：
        </p>
        ` + content + `
        <p style="color: #a8979a; margin: 16px auto;">(邮件由系统自动发出，请勿回复，谢谢！)</p>
    </div>
</div>`
}

func GetReplyEmailHTML(siteName, siteURL, theme, to, from, replierURL, title, titleURL, content string) string {
	return `<div style="
        box-shadow: 0 1px 4px rgba(0,0,0,.12);
        border-radius: 10px;
        color: #111;
        font-size: 12px;
        width: 95%;
        font-family: 微软雅黑, arial;
        margin: 10px auto;
        padding-bottom: 10px;
      ">
    <div class="adM"></div>
    <div style="
          width: 100%;
          background: #c2dbf6;
          min-height: 60px;
          color: white;
          border-radius: 6px 6px 0 0;
        ">
        <span style="
            line-height: 60px;
            min-height: 60px;
            margin-left: 30px;
            font-size: 16px;
          ">「<a style="color: #00a1ff; font-weight: 600; text-decoration: none;" href="` + siteURL + `" target="_blank"
                rel="external nofollow" target="_blank" rel="external nofollow" target="_blank">` + siteName + `</a>」` + theme + `</span>
    </div>
    <div style="margin: 0px auto; width: 90%; font-size: 14px;">
        <p style="margin: 16px auto;">
            <span style="font-weight: bold;">` + to + `</span>, 您好!
        </p>
        <p style="margin: 16px auto;">
            <a style="color: #00bbff; text-decoration: none;" href="` + replierURL + `" target="_blank" rel="external nofollow"
                target="_blank">` + from + `</a>在<a style="color: #00bbff; text-decoration: none;" href="` + titleURL + `" target="_blank"
                rel="external nofollow" target="_blank">` + title + `</a>回复了您：
        </p>
        ` + content + `
        <p style="color: #a8979a; margin: 16px auto;">(邮件由系统自动发出，请勿回复，谢谢！)</p>
    </div>
</div>`
}

func CreateRandomCode(num int) (string, error) {
	numBytes := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	length := bytes.NewReader(numBytes).Len()
	var sb strings.Builder

	for i := 0; i < num; i++ {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			return "", err
		}
		_, _ = fmt.Fprintf(&sb, "%d", numBytes[random.Int64()])
	}

	return sb.String(), nil
}
