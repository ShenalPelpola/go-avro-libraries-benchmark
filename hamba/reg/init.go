package reg

import (
	"io/ioutil"

	"github.com/hamba/avro"
)

// read the avro schema
func RegisterSchema(file string) (avro.Schema, error) {
	content, err := ioutil.ReadFile("../schemas/" + file)

	schema, err := avro.Parse(string(content))

	if err != nil {
		return nil, err
	}

	return schema, nil
}
