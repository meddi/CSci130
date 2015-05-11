package main

import (
	"html/template"
	"net/http"
    "github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("s3cret`c00kie_st0re+p@ssw0rd!"))
var mytemp *template.Template

func init() {

    mytemp = template.Must(template.ParseFiles("list.html","limited.html","alternative.html"))
    
	http.HandleFunc("/", handler)
	http.HandleFunc("/enter", enter)
	http.HandleFunc("/exit", exit)
    http.HandleFunc("/limited", limited)
    http.HandleFunc("/alternative", alternative)
   
}


func enter(w http.ResponseWriter, r *http.Request) {
    
    username   := r.FormValue("Name")
    password   := r.FormValue("Password")
    session, _ := store.Get(r, "UserSession")
    
    if(password == "ThePassword"){
        session.Values["name"] = username
        session.Values["error"] = ""
        session.Save(r, w)    
        http.Redirect(w,r,"/limited",302)
        return
    }else{
        session.Values["name"] = ""
        session.Values["error"] = "Invalid Password"
        session.Save(r, w)
        http.Redirect(w,r,"/",302)
        return
    }
    
}

func limited(w http.ResponseWriter, r *http.Request) {
    
    session, _ := store.Get(r, "UserSession")
    if(session.Values["name"] != nil && session.Values["name"] != ""){
        mytemp.ExecuteTemplate(w, "limited.html", "")
    }else{
        session.Values["error"] = "Access Denied"
        session.Save(r, w)
        http.Redirect(w,r,"/",302)
    }
    
}

func exit(w http.ResponseWriter, r *http.Request) {
    
    session, _ := store.Get(r, "UserSession")    
    session.Values["name"] = ""
    session.Values["error"] = "Logout Successful"
    session.Save(r, w)
    http.Redirect(w,r,"/",302)
    
}

func alternative(w http.ResponseWriter, r *http.Request) {
    
    session, _ := store.Get(r, "UserSession")
    if(session.Values["name"] != nil && session.Values["name"] != ""){
        mytemp.ExecuteTemplate(w, "alternative.html", "You are seeing this message cause you are logged in")
    }else{
        mytemp.ExecuteTemplate(w, "alternative.html","")
    }
}
func handler(w http.ResponseWriter, r *http.Request) {
    
    session, err := store.Get(r, "UserSession")
    errorText := ""
    if(err == nil){
        if(session.Values["name"] == nil || session.Values["name"] == ""){
            if(session.Values["error"] != nil){
                errorText = session.Values["error"].(string)
            }
            session.Values["error"] = ""
            session.Options = &sessions.Options{MaxAge: -1} 
            session.Save(r, w)
            mytemp.ExecuteTemplate(w, "list.html", errorText)
        }else{
            http.Redirect(w,r,"/limited",302)
        }
    }else{
        mytemp.ExecuteTemplate(w, "list.html", "")
    }
    
}


