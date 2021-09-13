package events

type PassengerProfile struct {
	ID   string `json:"id" avro:"id"`
	Type string `json:"type" avro:"type"`
	Body struct {
		ReferralCode         string `json:"referral_code" avro:"referral_code"`
		FcmToken             string `json:"fcm_token" avro:"fcm_token"`
		UserStatus           string `json:"user_status" avro:"user_status"`
		CreatedBy            int    `json:"created_by" avro:"created_by"`
		DeviceID             string `json:"device_id" avro:"device_id"`
		Email                string `json:"email" avro:"email"`
		Dob                  string `json:"dob" avro:"dob"`
		Firstname            string `json:"firstname" avro:"firstname"`
		PushVersion          int    `json:"push_version" avro:"push_version"`
		LoginStatus          string `json:"login_status" avro:"login_status"`
		LastActiveDate       string `json:"last_active_date" avro:"last_active_date"`
		DeviceType           int    `json:"device_type" avro:"device_type"`
		Lastname             string `json:"lastname" avro:"lastname"`
		ID                   int    `json:"id" avro:"id"`
		DeviceToken          string `json:"device_token" avro:"device_token"`
		Salutation           string `json:"salutation" avro:"salutation"`
		Address              string `json:"address" avro:"address"`
		CreatedAt            int64  `json:"created_at" avro:"created_at"`
		LastLogin            int64  `json:"last_login" avro:"last_login"`
		DefaultPaymentMethod int    `json:"default_payment_method" avro:"default_payment_method"`
		UpdatedAt            int64  `json:"updated_at" avro:"updated_at"`
		CountryID            int    `json:"country_id" avro:"country_id"`
		Phone                string `json:"phone" avro:"phone"`
		Password             string `json:"password" avro:"password"`
	} `json:"body" avro:"body"`
	Expiry    int64 `json:"expiry" avro:"expiry"`
	CreatedAt int64 `json:"created_at" avro:"created_at"`
	Version   int   `json:"version" avro:"version"`
	TraceInfo struct {
		TraceID struct {
			High int64 `json:"high" avro:"high"`
			Low  int64 `json:"low" avro:"low"`
		} `json:"trace_id" avro:"trace_id"`
		SpanID   int64 `json:"span_id"   avro:"span_id"`
		ParentID int64 `json:"parent_id" avro:"parent_id"`
		Sampled  bool  `json:"sampled" avro:"sampled"`
	} `json:"trace_info" avro:"trace_info"`
}
