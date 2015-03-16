package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
Dear {{.Honorific}}. {{.LastName}},
{{if .Attended}}
It was a pleasure to see you at the fundraiser.{{else}}
It is a shame you couldn't make it to the fundraiser.{{end}}
{{with .Donation}}Thank you for the donation of {{.}} dollars.
{{end}}
We look forward to seeing our you at our upcoming events!
{{range .Upcoming}} {{.}} 
{{end}}

Best wishes,
The Duke
`
	// Prepare some data to insert into the template.
	type Recipient struct {
		Honorific, LastName string
		Attended            bool
		Donation            int
		Upcoming            []string
	}
	var recipients = []Recipient{
		{"Dr", "Jones", true, 10000, []string{"Golfing with the stars", "Laughing at poor people", "General douchebaggery"}},
		{"Mr", "Smith", false, 0, []string{"Golfing with the stars", "Laughing at poor people", "General douchebaggery"}},
		{"Admiral", "Brown", true, 50000, []string{"Golfing with the stars", "Laughing at poor people", "General douchebaggery"}},
	}
	// STEP 1 & STEP 2
	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))
	//STEP 3
	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}

// FROM : http://golang.org/pkg/text/template/#example_Template
