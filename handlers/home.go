package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/devnandito/servergenpwd/utils"
)


func HandleHome(w http.ResponseWriter, r *http.Request) {
	title := "Home"
	subtitle := "Generate password"
	m := utils.GetMenu()
	header := filepath.Join("views", "header.html")
	footer := filepath.Join("views", "footer.html")
	menu := filepath.Join("views", "menu.html")
	js := filepath.Join("views", "js.html")
	home := filepath.Join("views/home", "index.html")
	gen := filepath.Join("views/home", "generate.html")
	switch r.Method {
	case "GET":
		tmpl, _ := template.ParseFiles(home, header, menu, footer, js)
		res := tmpl.Execute(w, map[string]interface{}{
			"Title": title,
			"Subtitle": subtitle,
			"Menu": m,
		})

		if res != nil {
			log.Println("Error executing template: ", res)
			return
		}
	case "POST":
		// len, err  := strconv.Atoi(r.PostFormValue("length"))
		// utils.Ckeck(err)
		msg := &utils.Home{
			Username: r.PostFormValue("username"),
			System: r.PostFormValue("system"),
			Length: r.PostFormValue("length"),
			Greetings: r.PostFormValue("greetings"),
			Filename: r.PostFormValue("filename"),
		}

		if !msg.Validate() {
			tmpl, _ := template.ParseFiles(home, header, menu, footer, js)
			res := tmpl.Execute(w, map[string]interface{} {
				"Title": title,
				"Subtitle": subtitle,
				"Msg": msg,
				"Menu": m,
			})

			if res != nil {
				log.Println("Error executing template: ", res)
				return
			}

		} else {
			len, err  := strconv.Atoi(r.PostFormValue("length"))
			utils.Ckeck(err)
			pwd := utils.GeneratePassword(len)
			lines, dataFile := utils.GetContentMail(msg.Greetings, msg.System, msg.Length, msg.Username, pwd)
			result := os.WriteFile("./files/"+msg.Filename+".txt", dataFile, 0644)
			utils.Ckeck(result)
			tmpl, _ := template.ParseFiles(gen, header, menu, footer, js)
			res := tmpl.Execute(w, map[string]interface{} {
				"Title": title,
				"Subtitle": subtitle,
				"Menu": m,
				"Msg": msg,
				"Pwd": pwd,
				"Lines": lines,
			})

			if res != nil {
				log.Println("Error executing template: ", res)
				return
			}
		}
	}
}