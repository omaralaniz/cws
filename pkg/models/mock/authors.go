package mock

import (
	"time"

	"github.com/omaralaniz/computerandwebstuff/pkg/models"
)

var mockAuthor = &models.Author{
	ID:      1,
	Name:    "Mr. Mockaton",
	Email:   "mockaton@gmail.com",
	Created: time.Now(),
	Active:  true,
}

type AuthorModel struct{}

func (pM *AuthorModel) Insert(name, email, password string) error {
	switch email {
	case "duplicate@email.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (a *AuthorModel) Authenticate(email, password string) (int, error) {
	switch email {
	case "mockaton@gmail.com":
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

func (pM *AuthorModel) Get(id int) (*models.Author, error) {
	switch id {
	case 1:
		return mockAuthor, nil
	default:
		return nil, models.ErrNoRecord
	}
}
