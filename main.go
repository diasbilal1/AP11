package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Payload struct {
	Content string `json:"content"`
}

type Reply struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func serveEndpoint(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод разрешен только POST", http.StatusMethodNotAllowed)
		return
	}

	var data Payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Ошибка при обработке JSON данных", http.StatusBadRequest)
		return
	}

	if data.Content == "" {
		respondJSON(w, http.StatusBadRequest, Reply{
			StatusCode: 400,
			Message:    "Неверное содержимое JSON",
		})
		return
	}

	log.Printf("Получено сообщение: %s\n", data.Content)

	respondJSON(w, http.StatusOK, Reply{
		StatusCode: 200,
		Message:    "Данные успешно получены",
	})
}

func respondJSON(w http.ResponseWriter, statusCode int, data Reply) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", serveEndpoint)

	log.Print("Запуск сервера на порту :8080")

	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
