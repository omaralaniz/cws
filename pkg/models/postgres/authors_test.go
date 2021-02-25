package postgres

import (
	"reflect"
	"testing"
	"time"

	"github.com/omaralaniz/computerandwebstuff/pkg/models"
)

func TestAuthorGet(t *testing.T) {
	if testing.Short() {
		t.Skip("postgres: we are skipping the integration test station")
	}

	// Set up a suite of table-driven tests and expected results.
	authors := []struct {
		name           string
		userID         int
		ExpectedAuthor *models.Author
		ExpectedError  error
	}{
		{
			name:   "Valid ID",
			userID: 1,
			ExpectedAuthor: &models.Author{
				ID:      1,
				Name:    "Mockaton Smith",
				Email:   "mockaton@gmail.com",
				Created: time.Date(2020, 12, 18, 0, 0, 0, 0, time.UTC),
				Active:  false,
			},
			ExpectedError: nil,
		},
		{
			name:           "NO ID",
			userID:         0,
			ExpectedAuthor: nil,
			ExpectedError:  models.ErrNoRecord,
		},
		{
			name:           "NON-EXIST ID",
			userID:         2,
			ExpectedAuthor: nil,
			ExpectedError:  models.ErrNoRecord,
		},
	}

	for _, aa := range authors {
		t.Run(aa.name, func(t *testing.T) {
			db, teardown := newTestDB(t)
			defer teardown()

			m := AuthorModel{db}

			user, err := m.Get(aa.userID)

			if err != aa.ExpectedError {
				t.Errorf("want %v; got %s", aa.ExpectedError, err)
			}

			if !reflect.DeepEqual(user, aa.ExpectedAuthor) {
				t.Errorf("want %v; got %v", aa.ExpectedAuthor, user)
			}
		})
	}
}
