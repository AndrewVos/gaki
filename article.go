package gaki

import (
  "time"
  "github.com/hoisie/mustache"
)

type Article struct {
  Application *Application
  FilePath string
  Text string
  Title string
  Date time.Time
  Year int
  WebPath string
  LastUpdated string
}

func (article *Article) Render() string {
  var context = article.Application.DefaultContext(article.Title)
  context["date"] = article.Date.Format("2006-01-02")
  context["path"] = article.WebPath
  rendered := mustache.Render(article.Text, context)
  rendered = HighlightCode(rendered)
  return rendered
}
