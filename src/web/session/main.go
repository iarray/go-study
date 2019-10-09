package session

import (
	"fmt"
	"sync"
)

var providers = make(map[string]SessionProvider)

type SessionManager struct {
	CookieName  string
	Provider    SessionProvider
	MaxLeftTime int64
	mu          sync.Mutex
}

type SessionProvider interface {
	Init(sid string) (Session, error)
	GetOrInit(sid string) (Session, error)
	Remove(sid string) error
	GC(maxLeftTime int64)
}

type Session interface {
	Get(key interface{}) interface{}
	Set(key, value interface{})
	Remove(key interface{})
	Id() string
}

func NewSessionManager(providerName, cookieName string, maxLeftTime int64) (*SessionManager, error) {
	if providerName == "" || cookieName == "" {
		return nil, fmt.Errorf("providerName or cookieName is Empty")
	}
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("providerName not Exists")
	}

	return &SessionManager{Provider: provider, CookieName: cookieName, MaxLeftTime: maxLeftTime}, nil
}

func RegisterProvider(providerName string, provider SessionProvider) {
	if provider == nil {
		panic("provider is nil")
	}
	providers[providerName] = provider
}
