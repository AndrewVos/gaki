package gaki

import (
  "time"
)

func ConvertTimeToISO8601(time time.Time) string {
  return time.Format("2006-01-02T15:04:05Z")
}
