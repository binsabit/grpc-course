package serializer_test

import (
	"testing"

	"github.com/binsabit/grpc-course/protogen/pb"
	"github.com/binsabit/grpc-course/samples"
	"github.com/binsabit/grpc-course/serializer"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	binaryFile := "../temp/laptop.bin"
	jsonFile := "../temp/laptop.json"
	laptop := samples.NewLaptop()
	err := serializer.WriteProtobufToFile(laptop, binaryFile)
	assert.Nil(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromFile(binaryFile, laptop2)
	assert.Nil(t, err)
	assert.True(t, proto.Equal(laptop2, laptop))

	err = serializer.WriteProtobufToJSON(laptop, jsonFile)
	assert.Nil(t, err)
}
