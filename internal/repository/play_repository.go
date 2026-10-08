package repository

import (
	"context"
	"theater-backend/internal/models"

	"github.com/jackc/pgx/v5"
)

type PlayRepository struct {
	db *pgx.Conn
}

func NewPlayRepository(db *pgx.Conn) *PlayRepository {
	return &PlayRepository{db: db}
}

// GetAll возвращает список всех спектаклей
func (r *PlayRepository) GetAll(ctx context.Context) ([]models.Play, error) {
	query := `SELECT id, title, genre, duration_minutes, description, age_restriction FROM plays`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plays []models.Play
	for rows.Next() {
		var p models.Play
		err := rows.Scan(&p.ID, &p.Title, &p.Genre, &p.DurationMinutes, &p.Description, &p.AgeRestriction)
		if err != nil {
			return nil, err
		}
		plays = append(plays, p)
	}

	return plays, nil
}

// Create добавляет новый спектакль в базу
func (r *PlayRepository) Create(ctx context.Context, p *models.Play) error {
	query := `
		INSERT INTO plays (title, genre, duration_minutes, description, age_restriction)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	return r.db.QueryRow(ctx, query, p.Title, p.Genre, p.DurationMinutes, p.Description, p.AgeRestriction).Scan(&p.ID)
}
