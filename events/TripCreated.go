package events

type TripCreated struct {
	ID        string          `json:"id" avro:"id"`
	Type      string          `json:"type" avro:"type"`
	Body      TripCreatedBody `json:"body" avro:"body"`
	CreatedAt int64           `json:"created_at" avro:"created_at"`
	Expiry    int64           `json:"expiry" avro:"expiry"`
	Version   int             `json:"version" avro:"version"`
	TraceInfo struct {
		TraceID struct {
			High int64 `json:"high" avro:"high"`
			Low  int64 `json:"low" avro:"low"`
		} `json:"trace_id" avro:"trace_id"`
		SpanID   int64 `json:"span_id" avro:"span_id"`
		ParentID int64 `json:"parent_id" avro:"parent_id"`
		Sampled  bool  `json:"sampled" avro:"sampled"`
	} `json:"trace_info" avro:"trace_info"`
}

type TripCreatedBody struct {
	Module           int    `json:"module" avro:"module"`
	BookedBy         int    `json:"booked_by" avro:"booked_by"`
	ServiceGroupCode string `json:"service_group_code" avro:"service_group_code"`
	TripID           int    `json:"trip_id" avro:"trip_id"`
	VehicleType      int    `json:"vehicle_type" avro:"vehicle_type"`
	PreBooking       bool   `json:"pre_booking" avro:"pre_booking"`
	SeatRequirement  int    `json:"seat_requirement" avro:"seat_requirement"`
	RouteID          int    `json:"route_id" avro:"route_id"`
	Passenger        struct {
		ID int `json:"id" avro:"id"`
	} `json:"passenger" avro:"passenger"`
	Driver struct {
		ID int `json:"id" avro:"id"`
	} `json:"driver" avro:"driver"`
	Corporate struct {
		ID    int `json:"id" avro:"id"`
		DepID int `json:"dep_id" avro:"dep_id"`
	} `json:"corporate" avro:"corporate"`
	Pickup struct {
		Time     int        `json:"time" avro:"time"`
		Location []Location `json:"location" avro:"location"`
	} `json:"pickup" avro:"pickup"`
	Drop struct {
		Location []Location `json:"location" avro:"location"`
	} `json:"drop" avro:"drop"`
	Promotion struct {
		Code string `json:"code" avro:"code"`
	} `json:"promotion" avro:"promotion"`
	Region struct {
		ID []int `json:"ids" avro:"ids"`
	} `json:"region" avro:"region"`
	Payment struct {
		PrimaryMethod   int `json:"primary_method" avro:"primary_method"`
		SecondaryMethod int `json:"secondary_method" avro:"secondary_method"`
	} `json:"payment" avro:"payment"`
	Comments struct {
		Remark      string `json:"remark" avro:"remark"`
		DriverNotes string `json:"driver_notes" avro:"driver_notes"`
	} `json:"comments" avro:"comments"`
	Filters struct {
		Driver struct {
			LanguageID int `json:"language_id" avro:"language_id"`
		} `json:"driver" avro:"driver"`
		Vehicle struct {
			CompanyID int `json:"company_id" avro:"company_id"`
			BrandID   int `json:"brand_id" avro:"brand_id"`
			ColorID   int `json:"color_id" avro:"color_id"`
		} `json:"vehicle" avro:"vehicle"`
	} `json:"filters" avro:"filters"`

	Boost struct {
		Name     string  `json:"name" avro:"name"`
		RegionID int     `json:"region_id" avro:"region_id"`
		Value    float64 `json:"value" avro:"value"`
		Type     string  `json:"type" avro:"type"`
	} `json:"boost" avro:"boost"`
	Surge struct {
		RegionID int     `json:"region_id" avro:"region_id"`
		Value    float32 `json:"value" avro:"value"`
	} `json:"surge" avro:"surge"`
	FareDetails struct {
		FareType         string  `json:"fare_type" avro:"fare_type"`
		MinKm            float32 `json:"min_km" avro:"min_km"`
		MinFare          float32 `json:"min_fare" avro:"min_fare"`
		AdditionalKmFare float32 `json:"additional_km_fare" avro:"additional_km_fare"`
		WaitingTimeFare  float32 `json:"waiting_time_fare" avro:"waiting_time_fare"`
		FreeWaitingTime  int     `json:"free_waiting_time" avro:"free_waiting_time"`
		NightFare        float32 `json:"night_fare" avro:"night_fare"`
		RideHours        float32 `json:"ride_hours" avro:"ride_hours"`
		ExtraRideFare    float32 `json:"extra_ride_fare" avro:"extra_ride_fare"`
		DriverBata       float32 `json:"driver_bata" avro:"driver_bata"`
		TripType         int     `json:"trip_type" avro:"trip_type"`
	} `json:"fare_details" avro:"fare_details"`
}

type Location struct {
	Address string  `json:"address" avro:"address"`
	Lat     float32 `json:"lat" avro:"lat"`
	Lng     float32 `json:"lng" avro:"lng"`
}
