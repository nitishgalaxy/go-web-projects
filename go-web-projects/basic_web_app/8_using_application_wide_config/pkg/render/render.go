package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/nitishgalaxy/go-course/pkg/config"
	"github.com/nitishgalaxy/go-course/pkg/models"
)

// Variable functions is of type template.FuncMap
// A funcMap is a map of functions I can use in a template
var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the templates package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// Data that I want available on every page
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	/*
		// Get the template cache
		tc, err := CreateTemplateCache()
		if err != nil {
			fmt.Println("Error getting template cache", err)
			// Stop the application
			log.Fatal(err)
		}
	*/

	var tc map[string]*template.Template

	if app.UseCache {
		// Production mode
		// Use caching
		// Get the template cache from App Config
		tc = app.TemplateCache
	} else {
		// Development mode
		// Get the template from disk
		// Reuild the cache
		tc, _ = CreateTemplateCache()

	}

	// Get the template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// Create bytes buffer
	buf := new(bytes.Buffer)

	// Add data common for every page
	td = AddDefaultData(td)

	// Take our template, execute it, and store in buf vaariable
	// _ = t.Execute(buf, nil)   // Without template data
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)

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
