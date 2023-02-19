package dir

import "time"

type file struct {
	Name       string
	Path       string
	Size       int64
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func NewFile(name string, path string, size int64, createdAt time.Time, modifiedAt time.Time) file {
	return file{
		Name:       name,
		Path:       path,
		Size:       size,
		CreatedAt:  createdAt,
		ModifiedAt: modifiedAt,
	}
}
