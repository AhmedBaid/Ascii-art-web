package ascii

import "net/http"

func FormFunc(w http.ResponseWriter, r *http.Request) {
	//cheack if the path is correct 
	if r.URL.Path != "/" {
		errore := ErrorPage{
			Code:         http.StatusNotFound,
			ErrorMessage: "The page you are looking for might have been removed, had its name changed, or is temporarily unavailable.",
		}

		w.WriteHeader(http.StatusNotFound)
		// execute the not found  template
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}	
	//check if the method is get
	if r.Method != http.MethodGet {
		errore := ErrorPage{
			Code:         http.StatusMethodNotAllowed,
			ErrorMessage: "The request method is not supported for the requested resource.",
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		// execute the not allowed method template
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}
	// execute the index template if the is no error
	err := Tp.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		errore := ErrorPage{
			Code:         http.StatusInternalServerError,
			ErrorMessage: "Something went wrong on our end. Please try again later.",
		}

		w.WriteHeader(http.StatusInternalServerError)
		// execute the internal server error template if there is an error
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
	}
}
