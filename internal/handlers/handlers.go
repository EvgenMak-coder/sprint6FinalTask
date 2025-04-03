package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlerRoot(w http.ResponseWriter, r *http.Request) {

	filePath := filepath.Join("..", "index.html")

	// Читаем содержимое HTML-файла
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Ошибка при чтении файла: %v", err)
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Устанавливаем заголовок Content-Type и отправляем содержимое
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err = w.Write(content)
	if err != nil {
		log.Printf("Ошибка при записи в файл: %v", err)
	}
}
func UploadHandler(w http.ResponseWriter, r *http.Request) {

	// Парсинг multipart-формы
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Получение файла из формы
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Failed to get file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Чтение содержимого файла
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Конвертация через сервис
	convertedData := service.Encode(string(fileBytes))

	// Генерация уникального имени файла
	timestamp := time.Now().UTC().Format("20060102-150405")
	cleanTimestamp := strings.ReplaceAll(timestamp, ":", "")
	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".txt"
	}
	newFilename := fmt.Sprintf("%s%s", cleanTimestamp, ext)

	// Сохранение файла
	if err := os.WriteFile(newFilename, []byte(convertedData), 0644); err != nil {
		http.Error(w, "Failed to save file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка результата клиенту
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err = w.Write([]byte(convertedData))
	if err != nil {
		log.Printf("Ошибка при записи в файл: %v", err)
	}
}
