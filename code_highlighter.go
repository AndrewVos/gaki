package gaki

import (
  "strings"
  "regexp"
  "fmt"
  "io"
  "errors"
  "path"
  "os/exec"
)

func HighlightCode(html string) string {
  matcher := regexp.MustCompile("<pre>:::(\\w+)((?s).*?)</pre>")
  for _,matches:= range matcher.FindAllStringSubmatch(html, -1) {
    lang := matches[1]
    code := matches[2]
    highlighted, err := highlightWithPygments(lang, code)
    if err == nil {
      html = strings.Replace(html, matches[0], highlighted, 1)
    } else {
      fmt.Print(err)
    }
  }
  return html
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
