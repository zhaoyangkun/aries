package util

func GetEmailHTML(title string, receiver string, content string) string {
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
            font-size: 15px;
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
          >」` + title + `</span
        >
      </div>
      <div style="margin: 0px auto; width: 90%;">
        <p>
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
            margin: 15px 0px;
            padding-left: 20px;
            padding-right: 20px;
            border-top: #ddd 1px solid;
            border-right: #ddd 1px solid;
            padding-top: 20px;
          "
        >
          ` + content + `
        </p>
        <p style="color: #a8979a;">(邮件由系统自动发出，请勿回复，谢谢！)</p>
      </div>
    </div>`
}

func GetFeedBackEmailHTML() {

}
