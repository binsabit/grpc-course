package serializer

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
)

func WriteProtobufToJSON(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("could not marshal to json: %v", err)
	}
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("could not write data to file: %v", err)
	}
	return nil
}

func WriteProtobufToFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("could not marshal protobuf to binary: %v", err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("could not write data to file: %v", err)
	}
	return nil
}

func ReadProtobufFromFile(filename string, message proto.Message) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read file: %v", err)
	}
	err = proto.Unmarshal(bytes, message)
	log.Println(message)
	if err != nil {
		return fmt.Errorf("could not unmarshal protobuf message: %v", err)
	}
	return nil
}
