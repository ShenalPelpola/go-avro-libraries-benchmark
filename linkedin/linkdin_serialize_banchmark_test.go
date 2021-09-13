package linkedin

import (
	"testing"

	"github.com/ShenalPelpola/go-avro-libraries-benchmark/events/templates"
	"github.com/ShenalPelpola/go-avro-libraries-benchmark/linkdin/reg"
)

// go test -bench={methodname}  -benchtime 1m -benchmem -memprofile memprofile3.out -cpuprofile profile3.out

func BenchmarkSerializeTripAcceptedEvent(b *testing.B) {
	schema, err := reg.RegisterSchema("trip_accepted.avsc")

	if err != nil {
		b.Error(err)
	}

	event := templates.GetTripAcceptedEventTemplate()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := SerializeTripAcceptedEvent(schema, event)

			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkSerializePassengerProfileEvent(b *testing.B) {
	schema, err := reg.RegisterSchema("passenger_profile_schema.avsc")

	if err != nil {
		b.Error(err)
	}

	event := templates.GetPassengerProfileEvent()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := SerializePassengerProfileEvent(schema, event)

			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkSerializeTripCreatedEvent(b *testing.B) {
	schema, err := reg.RegisterSchema("trip_created.avsc")

	if err != nil {
		b.Error(err)
	}

	event := templates.GetTripCreatedEventTemplate()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := SerializeTripCreatedEvent(schema, event)

			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
