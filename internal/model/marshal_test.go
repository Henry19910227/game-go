package model

import (
	"fmt"
	"game-go/internal/req"
	"google.golang.org/protobuf/proto"
	"log"
	"testing"
)

func TestMarshal(t *testing.T) {
	model := &req.TestAssign{}
	model.MiniGameId = 1
	model.ElementsArray = []int32{1, 2, 3}

	data, err := proto.Marshal(model)
	if err != nil {
		log.Fatalf("Failed to marshal: %v", err)
	}
	fmt.Printf("Serialized data: %v\n", data) // [8, 1, 18, 3, 1, 2, 3]
}

func TestUnMarshal(t *testing.T) {

	data := []byte{8, 1, 18, 3, 1, 2, 3}

	var model req.TestAssign
	err := proto.Unmarshal(data, &model)
	if err != nil {
		log.Fatalf("Failed to Unmarshal: %v", err)
	}
	fmt.Printf("Deserialized object: %+v\n", model.ElementsArray)
	fmt.Printf("Deserialized object: %+v\n", model.MiniGameId)
}
