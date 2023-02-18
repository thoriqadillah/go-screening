package dir

import "time"

type file struct {
	name       string
	CreatedAt  time.Time
	ModifiedAt time.Time
	DeletedAt  time.Time
}

func NewFile(name string, now time.Time) file {
	return file{
		name:       name,
		CreatedAt:  now,
		ModifiedAt: now,
		DeletedAt:  time.Time{},
	}
}
