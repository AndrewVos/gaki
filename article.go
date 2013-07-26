package gaki

import (
  "os"
  "time"
  "io/ioutil"
  "strings"
  "regexp"
  "fmt"
  "path/filepath"
  "sort"
  "github.com/hoisie/mustache"
)

type Article struct {
  FilePath string
  Text string
  Title string
  Date time.Time
  Year int
  WebPath string
  LastUpdated string
  Rendered string
}

type Articles []*Article

func (s Articles) Len() int           { return len(s) }
func (s Articles) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Articles) Less(i, j int) bool { return s[i].Date.Unix() > s[j].Date.Unix() }

type ArticlesForYear struct {
  Year int
  Articles []*Article
}

var CachedArticles Articles

func ConvertTimeToISO8601(time time.Time) string {
  return time.Format("2006-01-02T15:04:05Z")
}

func FindArticleByFilePath(filePath string) *Article {
  for _, article := range CachedArticles {
    if article.FilePath == filePath {
      return article
    }
  }
  return nil
}

func LoadAllArticles(app *Application) {
  var scan = func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
    if fileInfo.IsDir() == false {
      buffer, _ := ioutil.ReadFile(path)
      split := strings.SplitN(string(buffer), "\n\n", 2)

      titleMatcher := regexp.MustCompile("title:\\s*(.*)\\s*")
      dateMatcher := regexp.MustCompile("date:\\s([\\d\\/]*)\\s*")

      article:= new(Article)
      article.FilePath = path
      article.Text = split[1]
      article.Title = titleMatcher.FindStringSubmatch(split[0])[1]
      article.Date,_ = time.Parse("2006/01/02", dateMatcher.FindStringSubmatch(split[0])[1])
      article.LastUpdated = ConvertTimeToISO8601(article.Date)
      article.WebPath = path
      article.WebPath = strings.Replace(article.WebPath, "articles/", "", -1)
      article.WebPath = strings.Replace(article.WebPath, "-", "/", 3)
      article.WebPath = strings.Replace(article.WebPath, ".txt", "", 1)
      article.WebPath += "/"
      CachedArticles = append(CachedArticles, article)
      fmt.Printf("[loaded article] " + article.Title +"\n")

      var context = app.DefaultContext(article.Title)
      context["date"] = article.Date.Format("2006-01-02")
      context["path"] = path
      article.Rendered = mustache.Render(article.Text, context)
    }
    return nil
  }

  err := filepath.Walk("./articles", scan)
  if err != nil { panic(err) }

  sort.Sort(CachedArticles)
}

func ArticlesGroupedByYear() []*ArticlesForYear {
  var years = make(map[int]bool)
  var uniqueYears = []int{}

  for _, article:= range CachedArticles {
    years[article.Date.Year()] = true
  }

  for year,_:= range years {
    uniqueYears = append(uniqueYears, year)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(uniqueYears)))

  articlesByYear := []*ArticlesForYear{}
  for _, year:= range uniqueYears {
    currentYear := new(ArticlesForYear)
    currentYear.Year = year
    for _, article:= range CachedArticles {
      if article.Date.Year() == year {
        currentYear.Articles = append(currentYear.Articles, article)
      }
    }
    articlesByYear = append(articlesByYear, currentYear)
  }

  return articlesByYear
}
