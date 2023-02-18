package persistent

import (
	"bytes"
	"encoding/json"
	"io"
)

func toJSON(value interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(value, "", "\t")
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}

func toStruct(r io.Reader, value interface{}) error {
	return json.NewDecoder(r).Decode(value)
}
