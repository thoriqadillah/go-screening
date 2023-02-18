package dir

import "time"

type file struct {
	Name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
	DeletedAt  time.Time
}

func NewFile(name string, now time.Time) file {
	return file{
		Name:       name,
		CreatedAt:  now,
		ModifiedAt: now,
		DeletedAt:  time.Time{},
	}
}
