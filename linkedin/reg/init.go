package reg

import (
	"io/ioutil"

	"github.com/linkedin/goavro/v2"
)

// read the avro schema
func RegisterSchema(file string) (*goavro.Codec, error) {

	content, err := ioutil.ReadFile("../schemas/" + file)

	schema, err := goavro.NewCodec(string(content))

	if err != nil {
		return nil, err
	}
	return schema, nil
}
