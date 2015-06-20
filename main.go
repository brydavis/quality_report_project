package main

// _ "github.com/mattn/go-sqlite3"

func main() {
	report := NewReport(
		"Data Quality Report",
		"Shelter and Housing Programs",
		"data/data.json",
		6,
		2015,
	)
	report.ListenAndServe(8080)
}
