package crypto

import "google.golang.org/protobuf/proto"

type Tool interface {
	Marshal(mid int, sid int, payload proto.Message) ([]byte, error)
	Mid(b []byte) (int, error)
	Sid(b []byte) (int, error)
	Payload(b []byte, out proto.Message) error
}
