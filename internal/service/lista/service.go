package lista

import (
	"errors"
	"fmt"

	"github.com/CristianMarsico/seminario_Go/internal/config"
	"github.com/jmoiron/sqlx"
)

// Service interface
type Service interface {
	AddList(*Lista) (*Lista, error)
	GetAll() ([]*Lista, error)
	GetByID(int64) (*Lista, error)
	Delete(int64) error
	Edit(string, string) string
}

// Lista ...
type Lista struct {
	ID   int64
	Name string
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// NewBand_Art ...
func NewBand_Art(s string) *Lista {
	return &Lista{
		0,
		s,
	}
}

// New ...
func New(db *sqlx.DB, c *config.Config) (Service, error) {
	return service{db, c}, nil
}

// GetAll ...
func (s service) GetAll() ([]*Lista, error) {
	var l []*Lista

	query := "SELECT * FROM lista"
	err := s.db.Select(&l, query)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// GetByID ...
func (s service) GetByID(i int64) (*Lista, error) {
	query := `SELECT * FROM lista WHERE id = (?)`

	var l Lista
	err := s.db.Get(&l, query, i)
	if err != nil {
		return nil, err
	}
	return &l, nil
}

// AddList...
func (s service) AddList(l *Lista) (*Lista, error) {
	query := `INSERT INTO lista (name) VALUES (?)`

	res, err := s.db.Exec(query, l.Name)
	if err != nil {
		return nil, errors.New("DATABASE ERROR - " + err.Error())
	}
	LastID, _ := res.LastInsertId()
	l.ID = LastID
	return l, nil
}

// Delete ...
func (s service) Delete(l int64) error {
	query := `DELETE FROM lista WHERE id = (?)`
	res, err := s.db.Exec(query, l)

	if err != nil {
		return errors.New("DATABASE ERROR - " + err.Error())
	}
	col, _ := res.RowsAffected()
	if col == 0 {
		return errors.New("no existe tal ID")
	}
	return nil
}

// Edit ...
func (s service) Edit(n string, l string) string {
	query := `UPDATE lista SET name = ? WHERE id = ?`
	res, err := s.db.Exec(query, n, l)

	if err != nil {
		return fmt.Sprintf("%v", errors.New("DATABASE ERROR - "+err.Error()))
	}
	RowsAffected, _ := res.RowsAffected()

	return fmt.Sprintf("modificado: %d", RowsAffected)
}
