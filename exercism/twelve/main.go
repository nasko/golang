package twelve

import (
	"fmt"
	"strings"
)

// Gift exported
type Gift struct {
	day       int
	dayCount  string
	giftCount string
	gift      string
}

var gifts = []Gift{
	{1, "first", "a", "Partridge"},
	{2, "second", "two", "Turtle Doves"},
	{3, "third", "three", "French Hens"},
	{4, "fourth", "four", "Calling Birds"},
	{5, "fifth", "five", "Gold Rings"},
	{6, "sixth", "six", "Geese-a-Laying"},
	{7, "seventh", "seven", "Swans-a-Swimming"},
	{8, "eighth", "eight", "Maids-a-Milking"},
	{9, "ninth", "nine", "Ladies Dancing"},
	{10, "tenth", "ten", "Lords-a-Leaping"},
	{11, "eleventh", "eleven", "Pipers Piping"},
	{12, "twelfth", "twelve", "Drummers Drumming"},
}

// Verse exported
func Verse(input int) string {
	verseTemplate := `On the %v day of Christmas my true love gave to me, %v in a Pear Tree.`
	return fmt.Sprintf(verseTemplate, gifts[input-1].dayCount, compileGifts(input))
}

func compileGifts(input int) string {
	var giftParts []string
	var prefixPartrige string

	for i := input; i > 0; i-- {
		if input != 1 && i == 1 {
			prefixPartrige = "and "
		}
		giftParts = append(giftParts, fmt.Sprintf("%v%v %v", prefixPartrige, gifts[i-1].giftCount, gifts[i-1].gift))
	}
	return strings.Join(giftParts, ", ")
}

// Song exported
func Song() string {
	var verses []string
	for _, d := range gifts {
		verses = append(verses, Verse(d.day))
	}
	return fmt.Sprintf("%v\n", strings.Join(verses, "\n"))
}
