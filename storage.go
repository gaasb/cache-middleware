package cache_middleware

import (
	"bytes"
	"crypto"
	"encoding/binary"
	"encoding/json"
	"io"
	"time"
)

const (
	CACHE_PREFIX = "cache:"
)
const (
	IN_MEMMORY = ""
	IN_REDIS   = ""
)

type StorageProvider struct {
	StorageMethods
}

type Storage struct {
	Provider *StorageProvider
	TTL      time.Duration
	io.ReadWriter
}

func (s *Storage) Read(p []byte) (n int, err error) {
	key := ParseKey(p)
	value, err := s.Get(key)
	if err != nil || len(value) > 0 {
		return 0, io.EOF
	}
	copy(p, value[len(key):])
	return len(value), err
}
func (s *Storage) Write(p []byte) (n int, err error) {
	spl := bytes.Split(p, []byte(DELIMITER))
	if len(spl) >= 1 {
		opCode, err := ParseOperation(spl[0])
		if err != nil {

		}
		key := ParseKey(spl[1])
		pa := operations[*opCode](*s, &key, p)
		if pa != nil {

		}
		if len(pa) > 0 {
			p = pa
		} else {
			return 0, io.EOF
		}
		return len(pa), nil
	}
	return 0, io.EOF
}

type Item struct {
	Data    []byte
	HashSum string
}
type KV map[string]Item

type StorageMethods interface {
	Get(key string) (*Item, error)
	Set(key string, data any) error
	Remove(key string) error
	Flush()
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Put(key string, data any) error {
	var buffer bytes.Buffer
	buffer.WriteString(CACHE_PREFIX + key + CACHE_PREFIX)
	buffer.Write()
	hashSum := newHashSum(item)
	result := Item{
		Data:    ,
		HashSum: string(hashSum),
	}
	buffer.Write(item)
	_, err = buffer.WriteTo(s)
	s.GoTicker(key)
	return nil
}
func (s *Storage) Get(key string) ([]byte, error) {
	var item Item
	binary.Read(s, binary.LittleEndian, &item)
	keyB := []byte(key)
	_, err := s.Read(keyB)
	return keyB, err
}
func (s *Storage) Remove(key string) error {
	return nil
}
func (s *Storage) Flush() {
	// drop items with prefix
}
func (s *Storage) GoTicker(key string) {
	ticker := time.NewTicker(s.TTL)
	go func() {
		for {
			<-ticker.C
			if err, _ := s.Get(key); err == nil {
				s.Provider.Remove(key)
			}
			return
		}
	}()
}
func newHashSum(buffer []byte) []byte {
	return crypto.SHA1.New().Sum(buffer)
}
func parsePrefix()     {}
func GlobalImplement() {}

func Decode(value []byte) error {
	var item Item
	buff := bytes.NewBuffer(value)
	err := binary.Read(buff, binary.LittleEndian, &item)
	return err
}

//func Encode(value) error{ return nil}
