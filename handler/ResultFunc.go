package handler

import (
	"net/http"
	"unicode"

	tools "ascii/tools"
	ascii "ascii/functions"
	helpers "ascii/helpers"
)

var LastResult string

func ResultFunc(w http.ResponseWriter, r *http.Request) {
	// check if the method is post
	if r.Method != http.MethodPost {
		errore := tools.ErrorPage{
			Code:         http.StatusMethodNotAllowed,
			ErrorMessage: "The request method is not supported for the requested resource.",
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		// execute the not allowed method template
		tools.Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}

	Word := r.FormValue("word")
	typee := r.FormValue("typee")
	// this is to check the length of the word is correct without the \r
	ctr := 0
	for _, char := range Word {
		if char != '\r' {
			ctr++
		}
	}
	var errorMessage string

	if Word == "" {
		errorMessage = "Please enter a word."
	} else if typee == "" {
		errorMessage = "Please select a type."
	} else if ctr > 1000 {
		errorMessage = "The word length should not exceed 1000 characters."
	}
	// check if the word has invalid characters
	for i := 0; i < len(Word); i++ {
		if unicode.IsLetter(rune(Word[i])) && (Word[i] < 32 || Word[i] > 126) {
			errorMessage = "Invalid characters."
			break
		}
	}
	// ascci function to get the result as a ascci
	LastResult = ascii.Ascii(Word, typee)
	//  chack if there is an error in opening the file
	if LastResult == "" && Word != "" && typee != "" {
		errorMessage = "invalid file name !!!!!"
	}
	if errorMessage != "" {
		helpers.RenderTemplates(w, "index.html", errorMessage, 400)
		return
	}
	helpers.RenderTemplates(w, "result.html", LastResult, 200)
}
