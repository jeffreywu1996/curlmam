package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
)

type PageData struct {
	Result string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Render the form
			tmpl, err := template.ParseFiles("templates/index.html")
			if err != nil {
				http.Error(w, "Error rendering page", http.StatusInternalServerError)
				log.Println(err)
				return
			}
			tmpl.Execute(w, PageData{})
		} else if r.Method == http.MethodPost {
			// Handle form submission
			mamID := r.FormValue("mam_id")
			if mamID == "" {
				http.Error(w, "mam_id is required", http.StatusBadRequest)
				return
			}

			// Run the curl command
			cmd := exec.Command("curl", "-b", fmt.Sprintf("mam_id=%s", mamID), "https://t.myanonamouse.net/json/dynamicSeedbox.php")
			var out bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &out
			err := cmd.Run()

			result := out.String()
			if err != nil {
				result = fmt.Sprintf("Error: %v\n\nOutput:\n%s", err, result)
			}

			// Render the result back to the user
			tmpl, err := template.ParseFiles("templates/index.html")
			if err != nil {
				http.Error(w, "Error rendering page", http.StatusInternalServerError)
				log.Println(err)
				return
			}
			tmpl.Execute(w, PageData{Result: result})
		}
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
