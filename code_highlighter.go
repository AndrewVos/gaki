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
    fmt.Printf("\ncode:" + pygmentizer.Highlight(matches[1], matches[2]) + "\n")
    html = strings.Replace(html, matches[0], pygmentizer.Highlight(matches[1], matches[2]), 1)
  }
  return html
}
