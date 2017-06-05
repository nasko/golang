package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First string
	Last string
	Age int
	notExported int
}

func main() {
	nasko := Person{"Атанас", "Василев", 44, 007}
	json.NewEncoder(os.Stdout).Encode(nasko)
}
