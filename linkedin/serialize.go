package linkdin

import (
	"encoding/json"

	"github.com/ShenalPelpola/go-avro-libraries-benchmark/events"
	"github.com/linkedin/goavro/v2"
)

/*
 BinaryFromNative method requires the data to be passed in a map.
 The struct has to be converted to a map.
*/

func SerializeTripAcceptedEvent(schema *goavro.Codec, event events.TripAccepted) ([]byte, error) {
	var mp map[string]interface{}
	d, err := json.Marshal(event)
	json.Unmarshal(d, &mp)

	binary, err := schema.BinaryFromNative(nil, mp)

	if err != nil {
		return nil, err
	}

	return binary, nil
}

func SerializePassengerProfileEvent(schema *goavro.Codec, event events.PassengerProfile) ([]byte, error) {
	var mp map[string]interface{}
	d, err := json.Marshal(event)
	json.Unmarshal(d, &mp)

	binary, err := schema.BinaryFromNative(nil, mp)

	if err != nil {
		return nil, err
	}

	return binary, nil
}

func SerializeTripCreatedEvent(schema *goavro.Codec, event events.TripCreated) ([]byte, error) {
	var mp map[string]interface{}
	d, err := json.Marshal(event)
	json.Unmarshal(d, &mp)

	binary, err := schema.BinaryFromNative(nil, mp)

	if err != nil {
		return nil, err
	}

	return binary, nil
}
