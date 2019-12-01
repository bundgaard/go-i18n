package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

// Language ...
type Language map[string]map[string]string

func main() {
	languageContent, _ := ioutil.ReadFile("language.json")
	var language map[string]map[string]string
	fmt.Println(string(languageContent))

	if err := json.Unmarshal(languageContent, &language); err != nil {
		fmt.Println(err)
	}

	fmt.Println(language)

	root := http.NewServeMux()
	root.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		clientLanguage := r.Header.Get("Accept-Language")
		log.Println(clientLanguage)

		tpls, err := template.New("index.html").Funcs(template.FuncMap{"getField": func(name string) string {
			return language["sv_SE"][name]
		}}).ParseFiles(filepath.Join("public", "index.html"))
		if err != nil {
			log.Fatal(err)
		}
		if err := tpls.Execute(w, language); err != nil {
			log.Fatal(err)
		}

	})

	log.Fatal(http.ListenAndServe(":8000", root))
}
