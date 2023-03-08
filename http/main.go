package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

var tmplt *template.Template

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func handleAdd(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "add.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		age := r.FormValue("age")

		p := Person{
			Name: name,
			Age:  age,
		}

		// Read the JSON
		readFile, _ := os.Open("people.json")
		defer readFile.Close()
		var existingPeople []Person
		bytes, _ := io.ReadAll(readFile)
		json.Unmarshal(bytes, &existingPeople)
		people := append(existingPeople, p)

		// Write to the JSON
		writeFile, _ := os.Create("people.json")
		defer writeFile.Close()
		jsonData, _ := json.MarshalIndent(people, "", "  ")
		writeFile.Write(jsonData)
		http.Redirect(w, r, "/list", http.StatusSeeOther)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func handleList(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmplt, _ = template.ParseFiles("list.html")
		readFile, _ := os.Open("people.json")
		defer readFile.Close()
		var existingPeople []Person
		bytes, _ := io.ReadAll(readFile)
		json.Unmarshal(bytes, &existingPeople)

		data := map[string]interface{}{
			"people": existingPeople,
		}
		err := tmplt.Execute(w, data)

		if err != nil {
			return
		}
	}
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/list", handleList)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe("localhost:3000", nil); err != nil {
		log.Fatal(err)
	}
}
