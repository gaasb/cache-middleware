package cache_middleware

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	REMOVE CacheOperation = iota
	PUT
	GET
	FLUSH_ALL
)
const DELIMITER = ":\t"

type CacheOperation int

var operations = map[CacheOperation]func(storage Storage, key *string, data []byte) []byte{
	REMOVE:    RemoveOp,
	PUT:       PutOp,
	GET:       GetOp,
	FLUSH_ALL: FlushOp,
}

func RemoveOp(s Storage, key *string, data []byte) []byte {
	s.Provider.Remove(*key)
	return nil
}
func PutOp(s Storage, key *string, data []byte) []byte {
	err := s.Provider.Set(*key, data)
	if err == nil {
		s.GoTicker(*key)
	}
	fmt.Println(err)
	return nil
}
func GetOp(s Storage, key *string, data []byte) []byte {
	value, err := s.Provider.Get(*key)
	if err != nil {
		return nil
	}
	return value.Data
}
func FlushOp(s Storage, key *string, data []byte) []byte {
	s.Provider.Flush()
	return nil
}

func ParseOperation(raw []byte) (*CacheOperation, error) {
	var opCode int
	buffer := bytes.NewBuffer(raw)
	err := binary.Read(buffer, binary.BigEndian, opCode)
	if err != nil {
		return nil, err
	}
	result := CacheOperation(opCode)
	return &result, nil
}
func ParseKey(raw []byte) string {
	return string(raw)
}
