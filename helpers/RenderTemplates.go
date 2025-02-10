package helpers

import (
	tools "ascii/tools"
	"bytes"
	"net/http"
)

func RenderTemplates(w http.ResponseWriter, temp string, LastResult string, status int) {
	var buf bytes.Buffer
	err := tools.Tp.ExecuteTemplate(&buf, temp, LastResult)
	if err != nil {
		errore := tools.ErrorPage{
			Code:         http.StatusInternalServerError,
			ErrorMessage: "Something went wrong on our end. Please try again later.",
		}

		w.WriteHeader(http.StatusInternalServerError)
		tools.Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}
