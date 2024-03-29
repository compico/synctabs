package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var (
	roothtmldir string = "public/html/"
)

func main() {
	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/getSiteData", getSiteDataHandle)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(
		roothtmldir+"index.html",
		roothtmldir+"navbar.html",
		roothtmldir+"header.html",
		roothtmldir+"footer.html",
		roothtmldir+"popup.html",
	)

	if err != nil {
		fmt.Fprintf(w, "Error Parse: %v", err)
	}

	hdata := "Index - SyncTabs"

	err = t.ExecuteTemplate(w, "header", hdata)
	if err != nil {
		fmt.Fprintf(w, "Error Header: %v", err)
	}
	err = t.ExecuteTemplate(w, "popup", nil)
	if err != nil {
		fmt.Fprintf(w, "Error Popup: %v", err)
	}
	err = t.ExecuteTemplate(w, "navbar", "")
	if err != nil {
		fmt.Fprintf(w, "Error Navbar: %v", err)
	}
	err = t.ExecuteTemplate(w, "index", nil)
	if err != nil {
		fmt.Fprintf(w, "Error Index: %v", err)
	}
	err = t.ExecuteTemplate(w, "footer", nil)
	if err != nil {
		fmt.Fprintf(w, "Error Footer: %v", err)
	}
}

func getSiteDataHandle(w http.ResponseWriter, r *http.Request) {

	var pf = struct {
		Site string `json:"site"`
	}{}

	w.Header().Add("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&pf)
	if err != nil {
		fmt.Printf("ERROR №1: %v", err)
		r.Response.Status = http.StatusText(500)
		fmt.Fprintf(w, "{ \"error\":\""+err.Error()+"\"}")
	}

	url := pf.Site

	res, err := getSiteData(url)
	if err != nil {
		fmt.Printf("ERROR №2: %v", err)
	}

	jsondata, err := res.toJSON()
	if err != nil {
		fmt.Fprintf(w, "ERROR №3: %v", err)
	}

	fmt.Fprint(w, jsondata)
}
