package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	s := server.Server(logger)

	if err := s.ListenAndServe(); err != nil {

		logger.Printf("ошибка запуска сервера: %s\n", err.Error())
		fmt.Print(&buf)
		return
	}
}
