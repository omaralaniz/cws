package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/omaralaniz/computerandwebstuff/pkg/models"
)

type PostModel struct {
	DB *pgxpool.Pool
}

func (pM *PostModel) Insert(title, author, content string) (int, error) {
	var id int
	query := `INSERT INTO POSTS(title, author, content, modified) 
		VALUES($1,$2,$3, CURRENT_DATE) RETURNING id`

	err := pM.DB.QueryRow(context.Background(), query, title, author, content).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (pM *PostModel) Get(id int) (*models.Post, error) {
	post := &models.Post{}
	query := `SELECT id, title, author, content, modified FROM posts WHERE id = $1`
	err := pM.DB.QueryRow(context.Background(), query, id).Scan(&post.ID, &post.Title, &post.Author, &post.Content, &post.Modified)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}

	return post, nil
}

func (pM *PostModel) Latest() ([]*models.Post, error) {
	query := `SELECT id, title, author, content, modified FROM posts ORDER by modified DESC LIMIT 10`

	rows, err := pM.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []*models.Post{}

	for rows.Next() {
		post := &models.Post{}

		err = rows.Scan(&post.ID, &post.Title, &post.Author, &post.Content, &post.Modified)

		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}