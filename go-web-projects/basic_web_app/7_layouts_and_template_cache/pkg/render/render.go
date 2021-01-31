package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// Variable functions is of type template.FuncMap
// A funcMap is a map of functions I can use in a template
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Get the template cache
	tc, err := CreateTemplateCache()
	if err != nil {
		fmt.Println("Error getting template cache", err)
		// Stop the application
		log.Fatal(err)
	}

	// Get the template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	// Create bytes buffer
	buf := new(bytes.Buffer)

	// Take our template, execute it, and store in buf vaariable
	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writing template to browser ", err)
	}

	/*
		// Old Logic: without template cache

		parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
		err = parsedTemplate.Execute(w, nil)

		if err != nil {
			fmt.Println("Error parsing template ", err)
		}

	*/
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	// Create a map of templates
	myCache := map[string]*template.Template{}

	// Get all templates which has word page in it
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// Loop through the pages
	for _, page := range pages {
		fmt.Println("Page is currently ", page)

		// Get actual name of the page from fullpath
		name := filepath.Base(page)

		// Create TemplateSet
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Find layouts
		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, err
}
