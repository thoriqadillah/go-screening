package persistent

import (
	"io"
	"os"
	"sync"
)

const PATH = "./tmp/"

var lock sync.Mutex

func Save(name string, value interface{}) error {
	lock.Lock()
	defer lock.Unlock()

	f, err := os.Create(PATH + name)
	if err != nil {
		return err
	}
	defer f.Close()

	r, err := toJSON(value)
	if err != nil {
		return err
	}

	if _, err := io.Copy(f, r); err != nil {
		return err
	}

	return nil
}

func Load(name string, value interface{}) error {
	lock.Lock()
	defer lock.Unlock()

	f, err := os.Open(PATH + name)
	if err != nil {
		return err
	}
	defer f.Close()

	return toStruct(f, value)
}
