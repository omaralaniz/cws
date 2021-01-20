package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/omaralaniz/computerandwebstuff/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthorModel struct {
	DB *pgxpool.Pool
}

var pgErr *pgconn.PgError

func (a *AuthorModel) Insert(name, email, password string) error {
	hasedpassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	query := `INSERT INTO authors (name, email, hashed_password, created)
		VALUES($1, $2, $3, NOW())`

	_, err = a.DB.Exec(context.Background(), query, name, email, string(hasedpassword))
	if err != nil {

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}
	return nil
}

func (a *AuthorModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte
	query := "SELECT id, hashed_password FROM authors WHERE email = $1 AND active = TRUE"
	err := a.DB.QueryRow(context.Background(), query, email).Scan(&id, &hashedPassword)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, models.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (a *AuthorModel) Get(id int) (*models.Author, error) {
	author := &models.Author{}

	query := `SELECT id, name, email, created, active FROM authors WHERE id = $1`
	err := a.DB.QueryRow(context.Background(), query, id).Scan(&author.ID, &author.Name, &author.Email, &author.Created, &author.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return author, nil
}
