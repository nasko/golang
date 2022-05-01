package main //expenses

import (
	"errors"
	"fmt"
) // Record represents an expense record.

type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	out := []Record{}
	for _, r := range in {
		if predicate(r) {
			out = append(out, r)
		}
	}
	return out
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	expenses := 0.0
	in = Filter(in, ByDaysPeriod(p))
	for _, r := range in {
		expenses += r.Amount
	}
	return expenses
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	expenses := 0.0
	categoryRecords := Filter(in, ByCategory(c))

	if len(categoryRecords) == 0 {
		return 0, errors.New("unknown category entertainment")
	}

	periodRecords := Filter(categoryRecords, ByDaysPeriod(p))

	for _, r := range periodRecords {
		expenses += r.Amount
	}
	return expenses, nil
}

func main() {
	p1 := DaysPeriod{From: 1, To: 30}
	p2 := DaysPeriod{From: 31, To: 60}

	records := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	fmt.Println(CategoryExpenses(records, p1, "entertainment"))
	fmt.Println(CategoryExpenses(records, p1, "rent"))
	fmt.Println(CategoryExpenses(records, p2, "rent"))
}
