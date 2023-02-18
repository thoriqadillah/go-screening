package dir

import "time"

type file struct {
	Name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
	DeletedAt  time.Time
}

func NewFile(name string, createdAt time.Time, modifiedAt time.Time) file {
	return file{
		Name:       name,
		CreatedAt:  createdAt,
		ModifiedAt: modifiedAt,
		DeletedAt:  time.Time{},
	}
}
