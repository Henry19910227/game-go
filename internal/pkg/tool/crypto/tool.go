package crypto

import (
	"encoding/base64"
	"encoding/json"
	"google.golang.org/protobuf/proto"
)

type Message struct {
	Mid     int    `json:"mid"`
	Sid     int    `json:"sid"`
	Payload string `json:"payload"` // base64字串
}

type tool struct {
}

func New() Tool {
	return &tool{}
}

func (t *tool) Marshal(mid int, sid int, payload proto.Message) ([]byte, error) {

	// 轉成 pb data
	pbData, err := proto.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// 設置 payload
	msg := Message{}
	msg.Mid = mid
	msg.Sid = sid
	msg.Payload = base64.StdEncoding.EncodeToString(pbData)

	// 轉成 json data
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t *tool) Mid(b []byte) (int, error) {
	var m Message
	if err := json.Unmarshal(b, &m); err != nil {
		return 0, err
	}
	return m.Mid, nil
}

func (t *tool) Sid(b []byte) (int, error) {
	var m Message
	if err := json.Unmarshal(b, &m); err != nil {
		return 0, err
	}
	return m.Sid, nil
}

func (t *tool) Payload(b []byte, out proto.Message) error {
	var m Message
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	data, err := base64.StdEncoding.DecodeString(m.Payload)
	if err != nil {
		return err
	}
	if err := proto.Unmarshal(data, out); err != nil {
		return err
	}
	return nil
}
