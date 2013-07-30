package gaki

import (
  "io/ioutil"
  "time"
  "strings"
  "regexp"
  "html"
  "github.com/hoisie/mustache"
)

type Article struct {
  Application *Application
  FilePath string
  Text string
  Title string
  Date time.Time
  Year int
  LastUpdated string
  rendered string
}

func CreateArticle(application *Application, path string) *Article {
  buffer, _ := ioutil.ReadFile(path)
  split := strings.SplitN(string(buffer), "\n\n", 2)

  titleMatcher := regexp.MustCompile("title:\\s*(.*)\\s*")
  dateMatcher := regexp.MustCompile("date:\\s([\\d\\/]*)\\s*")

  article := new(Article)
  article.Application = application
  article.FilePath = path
  article.Text = split[1]
  article.Title = titleMatcher.FindStringSubmatch(split[0])[1]
  article.Date,_ = time.Parse("2006/01/02", dateMatcher.FindStringSubmatch(split[0])[1])
  article.LastUpdated = ConvertTimeToISO8601(article.Date)
  return article
}

func (article *Article) WebPath() string {
  webPath := article.FilePath
  webPath = strings.Replace(webPath, "articles/", "", -1)
  webPath = strings.Replace(webPath, "-", "/", 3)
  webPath = strings.Replace(webPath, ".txt", "", 1)
  webPath += "/"
  return webPath
}

func (article *Article) ShortDate() string {
  return article.Date.Format("2006-01-02")
}

func (article *Article) Render() string {
  if article.rendered == "" {
    var context = map[string]interface{} {
      "url": article.Application.Host,
      "title": article.Application.Title,
      "author": article.Application.Author,
      "lastUpdated": CachedArticles[0].LastUpdated,
    }
    article.rendered = mustache.Render(article.Text, context)
    article.rendered = html.UnescapeString(article.rendered)
    article.rendered = HighlightCode(article.rendered)
  }
  return article.rendered
}
