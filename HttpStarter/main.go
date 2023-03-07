package main

import (
	"fmt"
	"net/http"
	"text/template"
	// "encoding/json"
	// "io"
	// "log"
	// "os"
)

var tmplt *template.Template

// Create a Person struct to be able to work with JSON (`json:"name"` etc.)

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Send the user to the index page (ServeFile)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		// Send the user to the add page (ServeFile)
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.

		// Instantiate a Person and assign the values from the Form to the properties of the Person

		// Read the JSON file and retrieve the existing people, and put them in a slice (those flexible arrays)
		// Append the new Person you created to this existingPeople slice

		// Write to the JSON
		// Write into the JSON file the new slice of people you created

		// Redirect to /list (http.Redirect(response, request, "url", http.StatusSeeOther))
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func handleList(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		// Use the code you used above to Read the JSON file and get the existing people

		// Parse the list.html into a template (ParseFiles)
		// Execute the template with the proper data
	}
}

func main() {
	// Create HandleFuncs here (3 Handlers)

	// Set up the ListenAndServe
}
