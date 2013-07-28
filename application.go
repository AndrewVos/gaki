package gaki

import (
  "strings"
  "os"
  "io/ioutil"
  "path/filepath"
  "github.com/hoisie/web"
  "github.com/hoisie/mustache"
)

type Application struct {
  Title string
  Author string
  Host string
  Port string
  Disqus string
  Layouts map[string] string
}

func (app *Application) Run() {
  app.Layouts = map[string]string { }
  layouts := []string {"layout", "index", "article", "index.xml"}
  for _, layout := range layouts {
    text,_ := ioutil.ReadFile("./views/" + layout + ".mustache")
    app.Layouts[layout] = string(text)
  }

  web.Get("/", app.home)
  web.Get("/(\\d{4}/\\d{2}/\\d{2}/.*)/", app.article)
  web.Get("/index.xml", app.rss)

  var scan = func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
    if fileInfo.IsDir() == false {
      web.Get("/(" + strings.Replace(filepath.Base(path), ".mustache", "", -1)+ ")", app.page)
    }
    return nil
  }
  filepath.Walk("./pages", scan)

  LoadAllArticles(app)
  web.Config.StaticDir = "./public"
  web.Run(":" + app.Port)
}

func (app *Application) article(path string) string {
  filePath := "articles/" + strings.Replace(path, "/", "-", -1) + ".txt"
  article:= FindArticleByFilePath(filePath)
  var context = map[string]interface{} {
    "title": app.Title + " - " + article.Title,
    "article": article,
    "disqus": app.Disqus,
  }
  return mustache.RenderInLayout(app.Layouts["article"], app.Layouts["layout"], context)
}

func (app *Application) home() string {
  var context = map[string]interface{} {
    "title": app.Title,
    "articlesByYear": ArticlesGroupedByYear(),
  }
  return mustache.RenderInLayout(app.Layouts["index"], app.Layouts["layout"], context)
}

func (app *Application) page(name string) string {
  page := ReadPage("./pages/" + name + ".mustache")
  var context = map[string]interface{} {
    "title": app.Title + " - " + page.Title,
  }
  return mustache.RenderInLayout(page.Body, app.Layouts["layout"], context)
}

func (app *Application) rss(ctx *web.Context) string {
  ctx.ContentType("application/xml")
  var context = map[string]interface{} {
    "url": app.Host,
    "title": app.Title,
    "author": app.Author,
    "articles": CachedArticles,
    "lastUpdated": CachedArticles[0].LastUpdated,
  }
  return mustache.Render(app.Layouts["index.xml"], context)
}
