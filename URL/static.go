package static

import (
	"html/template"
    "net/http"
)
var globalTemp *template.Template // trying out what's done in the example with template pointers

func init() {
    http.HandleFunc("/", handler)
	http.HandleFunc("/message", message)
	http.HandleFunc("/readmsg", readmsg)
	globalTemp = template.Must(template.ParseFiles("index.html","indexform.html","readmsg.html"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := globalTemp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func message(w http.ResponseWriter, r *http.Request) {
	err := globalTemp.ExecuteTemplate(w, "indexform.html", struct{Name string; Message string;}{r.FormValue("name"),r.FormValue("msg")})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func readmsg(w http.ResponseWriter, r *http.Request) {
	err := globalTemp.ExecuteTemplate(w, "readmsg.html", struct{Name string; Message string;}{r.FormValue("name"),r.FormValue("msg")})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}