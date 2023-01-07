package avros

type AccountCreateOrUpdateEvent struct {
	Email       string `json:"email" avro:"email"`
	FullNumber  string `json:"full_number" avro:"full_number"`
	Alias       string `json:"alias" avro:"alias"`
	City        string `json:"city" avro:"city"`
	District    string `json:"district" avro:"district"`
	Name        string `json:"name" avro:"name"`
	PublicPlace string `json:"public_place" avro:"public_place"`
	Status      string `json:"status" avro:"status"`
	ZipCode     string `json:"zip_code" avro:"zip_code"`
}

type AccountDeleteEvent struct {
	Email string `json:"email" avro:"email"`
}

type AccountGetEvent struct {
	Email string `json:"email" avro:"email"`
}

type AccountGetResponseEvent struct {
	Email       string `json:"email" avro:"email"`
	FullNumber  string `json:"full_number" avro:"full_number"`
	Alias       string `json:"alias" avro:"alias"`
	City        string `json:"city" avro:"city"`
	DateTime    string `json:"date_time" avro:"date_time"`
	District    string `json:"district" avro:"district"`
	Name        string `json:"name" avro:"name"`
	PublicPlace string `json:"public_place" avro:"public_place"`
	Status      string `json:"status" avro:"status"`
	ZipCode     string `json:"zip_code" avro:"zip_code"`
}
