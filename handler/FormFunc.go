package handler

import (
	"net/http"

	tools "ascii/tools"
	helpers "ascii/helpers"
)

func FormFunc(w http.ResponseWriter, r *http.Request) {
	// cheack if the path is correct
	if r.URL.Path != "/" {
		errore := tools.ErrorPage{
			Code:         http.StatusNotFound,
			ErrorMessage: "The page you are looking for might have been removed, had its name changed, or is temporarily unavailable.",
		}

		w.WriteHeader(http.StatusNotFound)
		// execute the not found  template
		tools.Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}
	// check if the method is get
	if r.Method != http.MethodGet {
		errore := tools.ErrorPage{
			Code:         http.StatusMethodNotAllowed,
			ErrorMessage: "The request method is not supported for the requested resource.",
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		// execute the not allowed method template
		tools.Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}
	// execute the index template if the is no error
	helpers.RenderTemplates(w, "index.html", "", 200)
}
