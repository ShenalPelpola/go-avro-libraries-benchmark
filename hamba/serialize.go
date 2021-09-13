package hamba

import (
	"github.com/ShenalPelpola/go-avro-libraries-benchmark/events/templates"
	"github.com/hamba/avro"
)

func SerializeEvent(schema avro.Schema, event interface{}) ([]byte, error) {
	data, err := avro.Marshal(schema, event)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func SerializePassengerProfileEvent(schema avro.Schema) ([]byte, error) {
	record := templates.GetPassengerProfileEvent()

	data, err := avro.Marshal(schema, record)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func SerializeTripCreatedEvent(schema avro.Schema) ([]byte, error) {
	record := templates.GetTripCreatedEventTemplate()

	data, err := avro.Marshal(schema, record)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func SerializeTripAcceptedEvent(schema avro.Schema) ([]byte, error) {
	record := templates.GetTripAcceptedEventTemplate()

	data, err := avro.Marshal(schema, record)

	if err != nil {
		return nil, err
	}

	return data, nil
}
