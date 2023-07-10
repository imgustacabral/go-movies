package repository

import (
	"database/sql"

	"github.com/imgustacabral/go-movies/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
