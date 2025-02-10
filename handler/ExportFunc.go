package handler

import (
	"net/http"
	"strconv"

	tools "ascii/tools"
)

func Exportfunc(w http.ResponseWriter, r *http.Request) {
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

	// here i should check if the export is empty
	fileContent := r.FormValue("export")

	if fileContent == "" {
		errore := tools.ErrorPage{
			Code:         http.StatusBadRequest,
			ErrorMessage: " bad request you can't export an empty file",
		}

		w.WriteHeader(http.StatusBadRequest)
		// execute the bad request template
		tools.Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}

	// Set the response header to force a file download with the specified filename.
	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art-web.txt")
	// Set the Content-Length header to specify the size of the response body in bytes.
	w.Header().Set("Content-Length", strconv.Itoa(len(fileContent)))
	// Specify that the file content is plain text.
	w.Header().Set("Content-Type", "text/plain")
	// Write the file content to the response body, sending it to the client.
	_, err := w.Write([]byte(fileContent))
	if err != nil {
		errore := tools.ErrorPage{
			Code:         http.StatusInternalServerError,
			ErrorMessage: " Something went wrong on our end. We are working to resolve the issue. Please try again later.  ",
		}
		w.WriteHeader(http.StatusInternalServerError)
		tools.Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}
}
