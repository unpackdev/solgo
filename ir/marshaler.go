package ir

import (
	"fmt"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"
)

// NewTypedStruct creates a v3.TypedStruct from the given proto message and protoType.
func NewTypedStruct(m protoreflect.ProtoMessage, protoType string) *v3.TypedStruct {
	// Marshal the proto message to JSON.
	jsonBytes, err := protojson.Marshal(m)
	if err != nil {
		zap.L().Error("failed to marshal proto to json", zap.Error(err))
		return nil
	}

	// Unmarshal the JSON data into a structpb.Struct.
	s := &structpb.Struct{}
	if err := protojson.Unmarshal(jsonBytes, s); err != nil {
		zap.L().Error("failed to unmarshal json to structpb", zap.Error(err))
		return nil
	}

	// Create and return a v3.TypedStruct.
	return &v3.TypedStruct{
		TypeUrl: fmt.Sprintf("github.com/unpackdev/protos/unpack.v1.ir.%s", protoType),
		Value:   s,
	}
}
