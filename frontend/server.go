package main

import (
	"log"
	"net/http"
)

func main() {
	// Создаём файловый сервер, раздающий файлы из папки frontend
	fs := http.FileServer(http.Dir(".")) // "." означает текущую папку (frontend)
	http.Handle("/", fs)

	// Запускаем сервер на порту 8081
	log.Println("🚀 Frontend запущен на http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
