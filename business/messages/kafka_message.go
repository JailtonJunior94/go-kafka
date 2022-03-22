package messages

type KafkaMessage struct {
	Before   *CustomerMessage `json:"before"`
	After    *CustomerMessage `json:"after"`
	Op       string           `json:"op"`
	TsMs     int64            `json:"ts_ms"`
	Customer interface{}      `json:"customer"`
}

type CustomerMessage struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	Active    bool   `json:"active"`
}
