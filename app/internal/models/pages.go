package models

import (
	"errors"
)

// Define a Page type to hold the information for a single page.

type Page struct {
	ID      string
	Title   string
	Content string
	Created string
	Updated string
}

// Define a PageModel type which wraps a mocked up sql.DB connection pool.
type PageModel struct {
	pages map[string]*Page
}

func (m *PageModel) Get(id string) (string, error) {

	p, err := Page{ID: id}, errors.New("Not implemented")
	if err != nil {
		return p.ID, err
	}

	return p.ID, err
}
