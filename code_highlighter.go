package gaki

import (
  "strings"
  "regexp"
  "github.com/AndrewVos/pygmentizer"
  "fmt"
)

func HighlightCode(html string) string {
  matcher := regexp.MustCompile("<pre>:::(\\w+)((?s).*?)</pre>")
  for _,matches:= range matcher.FindAllStringSubmatch(html, -1) {
    lang := matches[1]
    code := matches[2]
    highlighted, err := pygmentizer.Highlight(lang, code)
    if err == nil {
      html = strings.Replace(html, matches[0], highlighted, 1)
    } else {
      fmt.Print(err)
    }
  }
  return html
}
