package ascii

import (
	"net/http"
	"os"
	"strings"
)

func StyleFunc(w http.ResponseWriter, r *http.Request) {
	filePath := strings.TrimPrefix(r.URL.Path, "/")
	File, err := os.Stat(filePath)
	if err != nil || File.IsDir() {

		errore := ErrorPage{
			Code:         http.StatusNotFound,
			ErrorMessage: "The page you are looking for might have been removed, had its name changed, or is temporarily unavailable.",
		}

		w.WriteHeader(http.StatusNotFound)
		Tp.ExecuteTemplate(w, "statusPage.html", errore)
		return
	}
	http.StripPrefix("/styles", http.FileServer(http.Dir("styles"))).ServeHTTP(w, r)
}
