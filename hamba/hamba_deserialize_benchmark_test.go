package hamba

import (
	"testing"

	"github.com/ShenalPelpola/go-avro-libraries-benchmark/events/templates"
	"github.com/ShenalPelpola/go-avro-libraries-benchmark/hamba/reg"
)

// go test -bench={methodname}  -benchtime 1m -benchmem -memprofile memprofile3.out -cpuprofile profile3.out
func BenchmarkDeserializeTripAcceptedEvent(b *testing.B) {
	schema, err := reg.RegisterSchema("trip_accepted.avsc")

	if err != nil {
		b.Error(err)
	}

	event := templates.GetTripAcceptedEventTemplate()

	data, err := SerializeEvent(schema, event)

	if err != nil {
		b.Fatal(err)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err = DeserializeTripAcceptedEvent(schema, data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkDeserializePassengerProfileEvent(b *testing.B) {
	schema, err := reg.RegisterSchema("passenger_profile_schema.avsc")

	if err != nil {
		b.Error(err)
	}

	event := templates.GetPassengerProfileEvent()

	data, err := SerializeEvent(schema, event)
	if err != nil {
		b.Fatal(err)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := DeserializePassengerProfileEvent(schema, data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkDeserializeTripCreatedEvent(b *testing.B) {
	schema, err := reg.RegisterSchema("trip_created.avsc")

	if err != nil {
		b.Error(err)
	}

	event := templates.GetTripCreatedEventTemplate()

	data, err := SerializeEvent(schema, event)

	if err != nil {
		b.Fatal(err)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err = DeserializeTripCreatedEvent(schema, data)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
