package message

type Model struct {
	Mid     int    `json:"mid"`
	Sid     int    `json:"sid"`
	Payload string `json:"payload"` // base64字串
}
