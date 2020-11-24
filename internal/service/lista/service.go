package lista

import (
	"github.com/CristianMarsico/seminario_Go/internal/config"
	"github.com/jmoiron/sqlx"
)

// Lista ...
type Lista struct {
	ID   int64
	Text string
}

// ChatService ...
type ChatService interface {
	AddMessage(Lista) error
	FindByID(int) *Lista
	FindAll() []*Lista
}
type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (ChatService, error) {
	return service{db, c}, nil //instancia de la estructura que respeta la interfaz
}

// AddMessage ...
func (s service) AddMessage(m Lista) error {
	return nil
}

// FindByID ...
func (s service) FindByID(ID int) *Lista {
	return nil

}

// FindAll ...
func (s service) FindAll() []*Lista {
	var list []*Lista
	//list = append(list, &Message{0, "Hello World"})
	s.db.Select(&list, "SELECT * FROM lista") //lo aloja en lugar de memoria de list
	return list

}
