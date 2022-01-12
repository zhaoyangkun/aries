package test

import (
	"aries/config/setting"
	"fmt"
	"testing"
)

func TestMarkdown(t *testing.T) {
	//setting.LuteEngine.SetCodeSyntaxHighlight(true)
	//setting.LuteEngine.SetCodeSyntaxHighlightStyleName("github")
	html := setting.LuteEngine.MarkdownStr("demo", "```go\npackage main\n```\n")
	fmt.Println(html)
}
