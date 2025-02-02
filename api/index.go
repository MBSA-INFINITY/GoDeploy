package handler

import (
	"net/http"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.GET("/", func(context *Context) {
		context.JSON(200, H{
			"message": "hello go from vercel !!!!",
		})
		// fp := path.Join("templates", "index.html")

		// templ, err := template.ParseFiles(fp)

		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }

		// if err := templ.Execute(w, nil); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }
	})
	server.Handle(w, r)
}
