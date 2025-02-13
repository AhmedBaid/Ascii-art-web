package main

import (
	"fmt"
	"net/http"
	"text/template"

	tools "ascii/tools"
	handler "ascii/handler"
)

func main() {
	var err error
	// parse all the html file from the template folder to variable Tp
	tools.Tp, err = template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Println("Error parsing templates: ", err)
		return
	}

	// Register handlers
	http.HandleFunc("/static/", handler.StyleFunc)
	http.HandleFunc("/export", handler.Exportfunc)
	http.HandleFunc("/ascii-art", handler.ResultFunc)
	http.HandleFunc("/", handler.FormFunc)
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
