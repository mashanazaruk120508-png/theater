package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"theater-backend/config"
	"theater-backend/internal/handler"
	"theater-backend/internal/repository"
	"theater-backend/internal/service"

	"github.com/jackc/pgx/v5"
)

func main() {
	cfg := config.LoadConfig()

	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Инициализация слоёв
	playRepo := repository.NewPlayRepository(conn)
	playService := service.NewPlayService(playRepo)
	playHandler := handler.NewPlayHandler(playService)

	// Маршруты (Endpoints)
	http.HandleFunc("/plays", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			playHandler.GetAllPlays(w, r)
		case http.MethodPost:
			playHandler.CreatePlay(w, r)
		default:
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}
	})

	fmt.Printf("Сервер запущен на порту %s\n", cfg.ServerPort)
	if err := http.ListenAndServe(cfg.ServerPort, nil); err != nil {
		log.Fatalf("Ошибка работы сервера: %v\n", err)
	}
}
