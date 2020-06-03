package handlers

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Handle handles
func Handle() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/download/cv", downloadCV).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(notFoundPage)

	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("images/"))))

	return r
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/home.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func notFoundPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("views/404.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func downloadCV(w http.ResponseWriter, r *http.Request) {
	fileName := "CV_HugoBlanc.pdf"

	openfile, err := os.Open(fileName)
	if err != nil {
		logrus.Errorf("error opening file %s: %v", fileName, err)
		http.Error(w, "File not found.", 404)
		return
	}
	defer openfile.Close()

	// Read the 512 first bytes of the file's headers
	FileHeader := make([]byte, 512)
	openfile.Read(FileHeader)

	FileContentType := http.DetectContentType(FileHeader)

	// Get informations about to file for the headers
	FileStat, _ := openfile.Stat()
	FileSize := strconv.FormatInt(FileStat.Size(), 10)

	parts := strings.Split(fileName, "/")
	// Send the headers
	w.Header().Set("Content-Disposition", "attachment; filename="+parts[len(parts)-1])
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	// Send the file
	openfile.Seek(0, 0)
	io.Copy(w, openfile)
}
