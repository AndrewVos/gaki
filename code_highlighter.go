package gaki

import (
  "strings"
  "regexp"
  "fmt"
  "io"
  "errors"
  "path"
  "os/exec"
  "html"
)

func HighlightCode(body string) string {
  matcher := regexp.MustCompile("<pre>:::(\\w+)((?s).*?)</pre>")
  for _,matches:= range matcher.FindAllStringSubmatch(body, -1) {
    lang := matches[1]
    code := html.UnescapeString(matches[2])
    highlighted, err := highlightWithPygments(lang, code)
    if err == nil {
      body = strings.Replace(body, matches[0], highlighted, 1)
    } else {
      fmt.Print(err)
    }
  }
  return body
}

func pygmentizerPath() string {
  return path.Join("vendor", "pygments", "pygmentizer.py")
}

func highlightWithPygments(language string, code string) (string, error) {
  cmd:= exec.Command("/usr/bin/env", "python", pygmentizerPath(), "-l", language, "-f", "html")
  writer, _ := cmd.StdinPipe()

  io.WriteString(writer, code)
  writer.Close()
  output,err := cmd.CombinedOutput()

  if err != nil {
    fmt.Printf(string(output))
    return code, errors.New(string(output))
  }
  return string(output), nil
}
