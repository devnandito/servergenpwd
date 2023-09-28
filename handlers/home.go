package handlers

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)


func HandleHome(w http.ResponseWriter, r *http.Request) {
	title := "Generate password"
	header := filepath.Join("views", "header.html")
	footer := filepath.Join("views", "footer.html")
	home := filepath.Join("views/home", "index.html")
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles(home, header, footer)
		res := tmpl.Execute(w, map[string]interface{} {
			"Title": title,
			"Text": "Hello áéíóú",
		})

		if res != nil {
			log.Println("Error executing template: ", res)
			return
		}
	// case "POST":
	// 	msg := &utils.ValidateHome {
	// 		Username: r.PostFormValue("username"),
	// 	}
	// 	if !msg.Validate() {
	// 		tmpl, _ := template.ParseFiles(home, header, footer)
	// 		res := tmpl.Execute(w, map[string]interface{} {
	// 			"Title": title,
	// 			"Msg": msg,
	// 		})

	// 		if res != nil {
	// 			log.Println("Error executing template: ", res)
	// 			return
	// 		}
	// 	} else {
	// 		tmpl, _ := template.ParseFiles(home, header, footer)
	// 		res := tmpl.Execute(w, map[string]interface{} {
	// 			"Title": title,
	// 			"Text": "Hello áéíóú",
	// 		})

	// 		if res != nil {
	// 			log.Println("Error executing template: ", res)
	// 			return
	// 		}
	// 	}
	}
}