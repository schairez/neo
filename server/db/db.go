package db

import (
	"github.com/schairez/neo/server/model"
)

//DB interface
type DB interface {
	GetNotes() ([]*model.Note, error)
}
