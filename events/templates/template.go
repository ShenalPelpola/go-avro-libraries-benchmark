package templates

import (
	"time"

	"github.com/ShenalPelpola/go-avro-libraries-benchmark/events"
	"github.com/google/uuid"
)

func GetPassengerProfileEvent() events.PassengerProfile {
	event := events.PassengerProfile{}

	event.ID = uuid.New().String()
	event.Type = "com.pickme.events.passenger.passenger_profile"
	event.CreatedAt = int64(time.Now().UnixNano() / 1e6)
	event.Body.ReferralCode = ""
	event.Body.FcmToken = ""
	event.Body.UserStatus = "A"
	event.Body.CreatedAt = 0
	event.Body.DeviceID = ""
	event.Body.Email = "simulator_Test_user_123@gmail.com"
	event.Body.Dob = "1994/01/01"
	event.Body.Firstname = "simulator_Test"
	event.Body.PushVersion = 0
	event.Body.LoginStatus = "D"
	event.Body.LastActiveDate = "2021/06/24"
	event.Body.DeviceType = 4
	event.Body.Lastname = "user_123"
	event.Body.ID = 0
	event.Body.DeviceToken = ""
	event.Body.Salutation = ""
	event.Body.Address = "14 maga pickme"
	event.Body.CreatedAt = int64(time.Now().UnixNano() / 1e6)
	event.Body.LastLogin = 0
	event.Body.DefaultPaymentMethod = 1
	event.Body.UpdatedAt = 1
	event.Body.CountryID = 1
	event.Body.Phone = "233532345"
	event.Body.Password = ""

	return event
}

func GetTripCreatedEventTemplate() events.TripCreated {
	event := events.TripCreated{}

	event.ID = uuid.New().String()
	event.Type = "trip_created"
	event.Version = 6
	event.CreatedAt = time.Now().UnixNano() / 1e6
	event.Body.TripID = 10000
	event.Body.Module = 1
	event.Body.RouteID = 0
	event.Body.Passenger.ID = 4242446
	event.Body.SeatRequirement = 1
	event.Body.FareDetails.MinFare = 0
	event.Body.FareDetails.MinKm = 0
	event.Body.FareDetails.FareType = ""
	event.Body.FareDetails.DriverBata = 0
	event.Body.FareDetails.AdditionalKmFare = 0
	event.Body.FareDetails.WaitingTimeFare = 0
	event.Body.FareDetails.FreeWaitingTime = 0
	event.Body.FareDetails.NightFare = 0
	event.Body.FareDetails.RideHours = 0
	event.Body.FareDetails.ExtraRideFare = 0
	event.Body.FareDetails.TripType = 0
	event.Body.Surge.Value = 0
	event.Body.Surge.RegionID = 0
	event.Body.Boost.RegionID = 0
	event.Body.Boost.Value = 0
	event.Body.Boost.Type = ""
	event.Body.Boost.Name = ""
	event.Body.Corporate.ID = 0
	event.Body.Corporate.DepID = 0
	event.Body.Payment.PrimaryMethod = 1
	event.Body.Payment.SecondaryMethod = 0
	event.Body.Driver.ID = 0
	event.Body.Comments.Remark = ""
	event.Body.Comments.DriverNotes = ""
	event.Body.Promotion.Code = ""
	event.Body.Filters.Driver.LanguageID = 0
	event.Body.Filters.Vehicle.BrandID = 0
	event.Body.Filters.Vehicle.ColorID = 0
	event.Body.Filters.Vehicle.CompanyID = 0
	event.Body.Region.ID = []int{17319, 17300, 17000, 10000}
	event.Body.Promotion.Code = ""
	event.Body.PreBooking = true
	event.Body.VehicleType = 3
	event.Body.BookedBy = 1
	event.Body.Pickup.Time = 1624013637
	event.Body.Pickup.Location = []events.Location{}
	event.Body.Drop.Location = []events.Location{}

	return event
}

func GetTripAcceptedEventTemplate() events.TripAccepted {
	event := events.TripAccepted{}

	event.ID = uuid.New().String()
	event.Type = "trip_created"
	event.Version = 6
	event.CreatedAt = time.Now().UnixNano() / 1e6
	event.Body.TripID = 10000
	event.Body.DriverID = 92556
	event.Body.Location.Address = ""
	event.Body.Location.Lat = 62.232
	event.Body.Location.Lng = 5.3434
	return event
}
