package internal

import (
	"context"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/MrGru/GruGo/cookbook/microservices/grpcjson/keyvalue"
)

type KeyValue struct {
	mutex sync.RWMutex
	m     map[string]string
}

func NewKeyValue() *KeyValue {
	return &KeyValue{
		m: make(map[string]string),
	}
}

func (k *KeyValue) Set(ctx context.Context, r *keyvalue.SetKeyValueRequest) (*keyvalue.KeyValueResponse, error) {
	k.mutex.Lock()
	k.m[r.GetKey()] = r.GetValue()
	k.mutex.Unlock()
	return &keyvalue.KeyValueResponse{Value: r.GetValue()}, nil
}

func (k *KeyValue) Get(ctx context.Context, r *keyvalue.GetKeyValueRequest) (*keyvalue.KeyValueResponse, error) {
	k.mutex.RLock()
	defer k.mutex.RUnlock()
	val, ok := k.m[r.GetKey()]
	if !ok {
		return nil, grpc.Errorf(codes.NotFound, "key not set")
	}
	return &keyvalue.KeyValueResponse{Value: val}, nil
}
