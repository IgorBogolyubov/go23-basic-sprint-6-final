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
	curDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Dir(curDir)
	absPath, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}

	indexFile := filepath.Join(absPath, "index.html")

	data, err := os.ReadFile(indexFile)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(data))
}

func HandlersUpload(w http.ResponseWriter, r *http.Request) {

	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Dir(curDir)
	absPath, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}

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

		textFile := filepath.Join(absPath, filename)

		data, err := os.ReadFile(textFile)
		if err != nil {
			log.Fatal(err)
		}

		out := service.Service(string(data))

		outFile := filepath.Join(dir, string(time.Now().Format("02.01.06 15_04_05")+filepath.Ext(".txt")))

		ftext, err := os.OpenFile(outFile, os.O_CREATE|os.O_RDWR, 0755)
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
		curDir, err := os.Getwd()

		if err != nil {
			log.Fatal(err)
		}

		dir := filepath.Dir(curDir)

		absPath, err := filepath.Abs(dir)
		if err != nil {
			log.Fatal(err)
		}

		indexFile := filepath.Join(absPath, "index.html")

		data, err := os.ReadFile(indexFile)
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(data))
	}
}
