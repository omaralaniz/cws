package mock

import (
	"time"

	"github.com/omaralaniz/computerandwebstuff/pkg/models"
)

var mockPost = &models.Post{
	ID:       1,
	Title:    "Mock Title",
	Author:   "Mr. Mockaton",
	Content:  "This mock info so pretend this is informative content",
	Modified: time.Now(),
	Category: "Codding Challenges",
	Summary:  "This mock info so pretend this is informative content",
}

type PostModel struct{}

func (pM *PostModel) Insert(title, author, category, content, summary string) (int, error) {
	return 2, nil
}

func (pM *PostModel) Update(id int, title, author, category, content, summary string) (int, error) {
	switch id {
	case 1:
		mockPost.Content = "Changed"
		return mockPost.ID, nil
	default:
		return 0, models.ErrNoRecord

	}
}

func (pM *PostModel) Get(id int) (*models.Post, error) {
	switch id {
	case 1:
		return mockPost, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (pM *PostModel) Latest() ([]*models.Post, error) {
	return []*models.Post{mockPost}, nil
}

func (pM *PostModel) Category(category string) ([]*models.Post, error) {
	switch category {
	case "Coding Challenges":
		return []*models.Post{mockPost}, nil
	default:
		return nil, models.ErrNoRecord
	}
}
