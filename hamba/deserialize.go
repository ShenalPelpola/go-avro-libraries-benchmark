package hamba

import (
	"github.com/ShenalPelpola/go-avro-libraries-benchmark/events"
	"github.com/hamba/avro"
)

func DeserializePassengerProfileEvent(schema avro.Schema, data []byte) (events.PassengerProfile, error) {
	deserialized := events.PassengerProfile{}

	err := avro.Unmarshal(schema, data, &deserialized)

	if err != nil {
		return deserialized, err
	}

	return deserialized, nil
}

func DeserializeTripCreatedEvent(schema avro.Schema, data []byte) (events.TripCreated, error) {
	deserialized := events.TripCreated{}

	err := avro.Unmarshal(schema, data, &deserialized)

	if err != nil {
		return deserialized, err
	}

	return deserialized, nil
}

func DeserializeTripAcceptedEvent(schema avro.Schema, data []byte) (events.TripAccepted, error) {
	deserialized := events.TripAccepted{}

	err := avro.Unmarshal(schema, data, &deserialized)

	if err != nil {
		return deserialized, err
	}

	return deserialized, nil
}
