package dir

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

const PATH = "./tmp/"

func save(name string, value []file) error {
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

func load(name string, value *[]file) error {
	f, err := os.Open(PATH + name)
	if err != nil {
		return err
	}
	defer f.Close()

	return toStruct(f, value)
}

func toJSON(value []file) (io.Reader, error) {
	b, err := json.MarshalIndent(value, "", "\t")
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}

func toStruct(r io.Reader, value *[]file) error {
	return json.NewDecoder(r).Decode(value)
}
