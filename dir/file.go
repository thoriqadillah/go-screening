package dir

import "time"

type file struct {
	name       string
	path       string
	createdAt  time.Time
	modifiedAt time.Time
}

func NewFile(name string, path string, createdAt time.Time, modifiedAt time.Time) file {
	return file{
		name:       name,
		path:       path,
		createdAt:  createdAt,
		modifiedAt: modifiedAt,
	}
}
