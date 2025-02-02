package handler

import (
	"bytes"
	"html/template"
	"net/http"
	"path"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.GET("/", func(context *Context) {
		// Define template path
		tmplPath := path.Join("templates", "index.html")

		// Parse template file
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			context.String(500, "Error loading template")
			return
		}

		// Create a buffer to store rendered template output
		var renderedTemplate bytes.Buffer
		data := map[string]interface{}{
			"Title":   "Welcome",
			"Message": "Hello, Go from Vercel!",
		}

		// Execute template and write output to buffer
		if err := tmpl.Execute(&renderedTemplate, data); err != nil {
			context.String(500, "Error rendering template")
			return
		}

		// Send rendered HTML as a response
		context.HTML(200, renderedTemplate.String(), nil)
	})
	server.Handle(w, r)
}
