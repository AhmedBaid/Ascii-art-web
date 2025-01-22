package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"
	"unicode"

	ascii "ascii/functions"
)

var tp *template.Template

type ErrorPage struct {
	Code         int
	ErrorMessage string
}

func main() {
	var err error

	tp, err = template.ParseGlob("template/*.html")
	if err != nil {
		fmt.Println("errror bro ", err)
	}

	http.HandleFunc("/styles/", styleFunc)
	http.HandleFunc("/ascii-art", ResultFunc)
	http.HandleFunc("/", formFunc)
	fmt.Println("Server running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func styleFunc(w http.ResponseWriter, r *http.Request) {
	filePath := "styles" + strings.TrimPrefix(r.URL.Path, "/styles")
	file, err := os.Stat(filePath)
	if err != nil || file.IsDir() {
		errore := ErrorPage{
			Code:         http.StatusNotFound,
			ErrorMessage: http.StatusText(404),
		}
		w.WriteHeader(http.StatusNotFound)
		tp.ExecuteTemplate(w, "notfound.html", errore)
		return
	}
	http.ServeFile(w, r, filePath)
}

func formFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errore := ErrorPage{
			Code:         http.StatusNotFound,
			ErrorMessage: http.StatusText(404),
		}
		w.WriteHeader(http.StatusNotFound)
		tp.ExecuteTemplate(w, "notfound.html", errore)
		return
	}

	if r.Method != http.MethodGet {
		errore := ErrorPage{
			Code:         http.StatusMethodNotAllowed,
			ErrorMessage: http.StatusText(400),
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		tp.ExecuteTemplate(w, "notfound.html", errore)
		return
	}
	tp.ExecuteTemplate(w, "index.html", nil)
}

func ResultFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		errore := ErrorPage{
			Code:         http.StatusNotFound,
			ErrorMessage: http.StatusText(400),
		}
		w.WriteHeader(http.StatusNotFound)
		tp.ExecuteTemplate(w, "notfound.html", errore)
		return
	}

	if r.Method != http.MethodPost {
		errore := ErrorPage{
			Code:         http.StatusMethodNotAllowed,
			ErrorMessage: http.StatusText(400),
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		tp.ExecuteTemplate(w, "notfound.html", errore)
		return
	}

	word := r.FormValue("word")
	typee := r.FormValue("typee")

	var errorMessage string
	if word == "" {
		errorMessage = "Please enter a word."
	} else if typee == "" {
		errorMessage = "Please select a type."
	} else if len(word) > 1000 {
		errorMessage = "The word length should not exceed 1000 characters."
	} else {
		for i := 0; i < len(word); i++ {
			if unicode.IsLetter(rune(word[i])) && (word[i] < 32 || word[i] > 126) {
				errorMessage = "invalid charts"
				break
			}
		}
	}

	if errorMessage != "" {

		w.WriteHeader(http.StatusBadRequest)
		tp.ExecuteTemplate(w, "index.html", errorMessage)

		return
	}
	LastResult := ascii.Ascii(word, typee)

	if LastResult == "" {

		errore := ErrorPage{
			Code:         http.StatusInternalServerError,
			ErrorMessage: http.StatusText(500),
		}

		w.WriteHeader(http.StatusInternalServerError)

		tp.ExecuteTemplate(w, "notfound.html", errore)
		return
	}

	tp.ExecuteTemplate(w, "result.html", LastResult)
}
