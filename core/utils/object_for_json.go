package utils

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
)

func JsonForObject[T any](s []byte) (error, T) {
	var x T

	err := json.Unmarshal(s, &x)
	if err != nil {
		return err, x
	}
	return nil, x
}

func ReaderForByte(o io.Reader) []byte {
	body, _ := ioutil.ReadAll(o)
	return body
}

func ObjectForJsonString(o any) string {
	byteJson, err := json.Marshal(o)
	if err != nil {
		return ""
	}
	return string(byteJson)
}

func ObjectForByte(o any) []byte {
	byteJson, err := json.Marshal(o)
	if err != nil {
		return nil
	}
	return byteJson
}

func StringForGuid(guid string) uuid.UUID {
	var id, _ = uuid.Parse(guid)
	return id
}
