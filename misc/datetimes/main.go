package main

// reference datetime:
// Monday 01/02 03:04:05PM '06 -0700

import (
	"fmt"
	"time"
)

func main() {
	t := parse("Thu, 19-05-2011, 22:47 +01:00")
	fmt.Println(t.Format(time.RFC3339))
}

func parse(dt string) time.Time {
	layout := "Mon, 02-01-2006, 15:04 -07:00"

	t, _ := time.Parse(layout, dt)
	return t
}
