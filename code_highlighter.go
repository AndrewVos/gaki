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
    highlighted, err := pygmentizer.Highlight(matches[1], matches[2])
    if err == nil {
      fmt.Print(err)
    } else {
      html = strings.Replace(html, matches[0], highlighted, 1)
    }
  }
  return html
}
