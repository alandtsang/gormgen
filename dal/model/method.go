package model

import "gorm.io/gen"

type Method interface {
	// SELECT * FROM @@table WHERE id=@id
	GetByID(id uint64) (*gen.T, error)
	// SELECT * FROM @@table WHERE name IN @names
	ListByNames(names ...string) ([]*gen.T, error)
	// DELETE FROM @@table WHERE id IN @ids
	DeleteByIDs(ids ...uint64) error
}
