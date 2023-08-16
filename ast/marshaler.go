package ast

import (
	"fmt"

	v3 "github.com/cncf/xds/go/xds/type/v3"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"
)

// NewTypedStruct creates a new v3.TypedStruct instance based on the provided protoreflect.ProtoMessage and protoType.
// It marshals the given ProtoMessage into JSON, then unmarshals it into a structpb.Struct, and constructs a TypedStruct
// with the appropriate type URL and structpb.Value.
// It returns the created TypedStruct instance or nil in case of errors during marshaling or unmarshaling.
func NewTypedStruct(m protoreflect.ProtoMessage, protoType string) *v3.TypedStruct {
	jsonBytes, err := protojson.Marshal(m)
	if err != nil {
		zap.L().Error("failed to marshal proto to json", zap.Error(err))
		return nil
	}

	s := &structpb.Struct{}
	if err := protojson.Unmarshal(jsonBytes, s); err != nil {
		zap.L().Error("failed to unmarshal json to structpb", zap.Error(err))
		return nil
	}

	return &v3.TypedStruct{
		TypeUrl: fmt.Sprintf(
			"github.com/txpull/protos/txpull.v1.ast.%s",
			protoType,
		),
		Value: s,
	}
}
