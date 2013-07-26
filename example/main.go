package main

import (
  "os"
  "github.com/AndrewVos/gaki"
)

func main() {
  application := new(gaki.Application)
  application.Title = "example.com"
  application.Author = "Some Dude"

  application.Host = "http://localhost"
  if os.Getenv("URL") != "" {
    application.Host = os.Getenv("URL")
  }

  if os.Getenv("PORT") == "" {
    application.Port = "9292"
    application.Host += ":" + application.Port
  } else {
    application.Port = os.Getenv("PORT")
  }

  application.Disqus = "disqus_short_name"
  application.Run()
}
