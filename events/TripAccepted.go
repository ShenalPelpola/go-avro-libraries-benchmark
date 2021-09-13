package events

type TripAccepted struct {
	ID   string `json:"id" avro:"id"`
	Type string `json:"type" avro:"type"`
	Body struct {
		TripID   int `json:"trip_id" avro:"trip_id"`
		DriverID int `json:"driver_id" avro:"driver_id"`
		Location struct {
			Address string  `json:"address" avro:"address"`
			Lat     float32 `json:"lat" avro:"lat"`
			Lng     float32 `json:"lng" avro:"lng"`
		} `json:"location" avro:"location"`
	} `json:"body" avro:"body"`
	CreatedAt int64 `json:"created_at" avro:"created_at"`
	Expiry    int64 `json:"expiry" avro:"expiry"`
	Version   int   `json:"version" avro:"version"`
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
