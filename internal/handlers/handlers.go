package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlersRoot(w http.ResponseWriter, r *http.Request) {

	data, err := os.ReadFile("../index.html")
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(data))
}

func HandlersUpload(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ищем в форме имя файла
	out1 := strings.Split(string(body), "filename=")
	out2 := strings.Split(out1[1], "\"")
	filename := out2[1]

	if filename != "" { // поверяем наличие файла в форме

		data, err := os.ReadFile("C:\\Users\\User\\go23-basic-sprint-6-final\\" + filename)
		if err != nil {
			log.Fatal(err)
		}

		out := service.Service(string(data))

		ftext, err := os.OpenFile(string(time.Now().Format("02.01.06 15_04_05")+filepath.Ext(".txt")), os.O_CREATE|os.O_RDWR, 0755)
		if err != nil {
			log.Fatal(err)
		}
		_, err = ftext.WriteString(out)
		if err != nil {
			log.Fatal(err)
		}
		ftext.Close()

		w.Write([]byte(out))
	} else { // если файл не выбран повторяем вывод формы запроса
		data, err := os.ReadFile("../index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(data))
	}
}
