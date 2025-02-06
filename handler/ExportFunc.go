package ascii

import (
	"net/http"
)

func Exportfunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/export" {
		errore := ErrorPage{
			Code:         http.StatusNotFound,
			ErrorMessage: "The page you are looking for might have been removed, had its name changed, or is temporarily unavailable.",
		}

		w.WriteHeader(http.StatusNotFound)
		// execute the not found template
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}
	// check if the method is post
	if r.Method != http.MethodPost {
		errore := ErrorPage{
			Code:         http.StatusMethodNotAllowed,
			ErrorMessage: "The request method is not supported for the requested resource.",
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		// execute the not allowed method template
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}

	// here i should check if the export is empty
	fileContent := r.FormValue("export")

	if fileContent == "" {
		errore := ErrorPage{
			Code:         http.StatusBadRequest,
			ErrorMessage: " bad request   you can't export an empty file",
		}

		w.WriteHeader(http.StatusBadRequest)
		// execute the not allowed method template
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=ascii-art-web.txt")
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte(fileContent))
	if err != nil {
		errore := ErrorPage{
			Code:         http.StatusInternalServerError,
			ErrorMessage: " Something went wrong on our end. We are working to resolve the issue. Please try again later.  ",
		}
		w.WriteHeader(http.StatusInternalServerError)
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}
}
