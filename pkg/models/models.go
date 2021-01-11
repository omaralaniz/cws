package models

import (
	"errors"
	"html/template"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Post struct {
	ID       int
	Title    string
	Author   string
	Content  string
	Modified time.Time
	Category string
	Summary  string
}

type PostEx struct {
	ID       int
	Title    string
	Author   string
	Content  template.HTML
	Modified time.Time
	Category string
	Summary  string
}

type Author struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}
