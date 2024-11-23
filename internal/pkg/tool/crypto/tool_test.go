package crypto

import (
	"fmt"
	"game-go/internal/model/res"
	"log"
	"testing"
)

func TestMarshal(t *testing.T) {
	pb := &res.UserData{}
	pb.UserId = 1
	pb.Score = 10000

	// 轉成 crypto
	msg, err := New().Marshal(1, 2, pb)
	if err != nil {
		log.Fatalf("Failed to Marshal: %v", err)
	}

	// 解出 mid
	mid, err := New().Mid(msg)
	if err != nil {
		log.Fatalf("Failed to Mid: %v", err)
	}
	fmt.Println(mid)

	// 解出 sid
	sid, err := New().Sid(msg)
	if err != nil {
		log.Fatalf("Failed to Mid: %v", err)
	}
	fmt.Println(sid)

	// 解出 payload
	var u = res.UserData{}
	if err := New().Payload(msg, &u); err != nil {
		log.Fatalf("Failed to Payload: %v", err)
	}
	fmt.Println(u.UserId)
	fmt.Println(u.Score)
}
