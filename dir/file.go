package dir

import "time"

type file struct {
	Name       string
	Path       string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func NewFile(name string, path string, createdAt time.Time, modifiedAt time.Time) file {
	return file{
		Name:       name,
		Path:       path,
		CreatedAt:  createdAt,
		ModifiedAt: modifiedAt,
	}
}
