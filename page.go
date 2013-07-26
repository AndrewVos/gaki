package gaki

import (
  "strings"
  "io/ioutil"
  "regexp"
)

type Page struct {
  Title string
  Body string
}

func ReadPage(path string) Page {
  buffer, _ := ioutil.ReadFile(path)
  split := strings.SplitN(string(buffer), "\n\n", 2)

  titleMatcher := regexp.MustCompile("title:\\s*(.*)\\s*")
  return Page {
    Title: titleMatcher.FindStringSubmatch(split[0])[1],
    Body: split[1],
  }
}
