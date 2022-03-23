package messages

type KafkaMessage struct {
	SchemaPrimary *SchemaPrimary `json:"schemaPrimary"`
	Payload       *Payload       `json:"payload"`
}

type SchemaPrimary struct {
	Type string `json:"type"`

	Optional bool `json:"optional"`
	Name     bool `json:"name"`
}

type Payload struct {
	Before *CustomerMessage `json:"before"`
	After  *CustomerMessage `json:"after"`
}

type CustomerMessage struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	Active    bool   `json:"active"`
}
