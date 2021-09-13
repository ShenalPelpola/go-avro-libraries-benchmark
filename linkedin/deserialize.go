package linkedin

import (
	"github.com/linkedin/goavro/v2"
)

func DeserializeEvent(schema *goavro.Codec, data []byte) (interface{}, error) {
	deserialized, _, err := schema.NativeFromBinary(data)

	if err != nil {
		return deserialized, err
	}

	return deserialized, nil
}
