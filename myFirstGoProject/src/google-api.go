package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/medical", medical)
	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rootForm)
}

const rootForm = `
  <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <title>Google Place Search API</title>
      </head>
      <body>
      <h1>Show Hospitals</h1>
        <form style="margin-left: 120px;" action="/medical" method="post" accept-charset="utf-8">
        <label for="str1"> Address</label><br/><br/>
          <input type="text" name="str1" placeholder="Adress...." id="str" /><br/><br/>
					<input type="submit" value="Submit" />
        </form>
      </body>
    </html>
`

var upperTemplate = template.Must(template.New("directions").Parse(upperTemplateHTML))

func medical(w http.ResponseWriter, r *http.Request) {
	// Sample address "1600 Amphitheatre Parkway, Mountain View, CA"
	addr1 := r.FormValue("str1")

	fullUrl := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/place/textsearch/xml?query=hospital+in+%s&key=AIzaSyD0nwU7DxevIexVJqoLCGy2rX2lred9568", addr1)

	// Build the request
	req, err := http.Get(fullUrl)
	if err != nil {
		panic(err)
	}

	jsonDataFromHttp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	tempErr := upperTemplate.Execute(w, string(jsonDataFromHttp))
	if tempErr != nil {
		http.Error(w, tempErr.Error(), http.StatusInternalServerError)
	}
}

const upperTemplateHTML = `
<!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <title>Display Image</title>
      <link rel="stylesheet" href="/stylesheets/goview.css">
    </head>
    <body>
      {{html .}}
    </body>
  </html>
	`
