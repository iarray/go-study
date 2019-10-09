package session

import (
	"fmt"
	"sync"
	"math/rand"
	"encoding/base64"
	"net/http"
	"net/url"
)

var providers = make(map[string]SessionProvider)

type SessionManager struct {
	cookieName  string
	provider    SessionProvider
	maxLeftTime int64
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

	return &SessionManager{provider: provider, cookieName: cookieName, maxLeftTime: maxLeftTime}, nil
}

/* 注册session存储提供者 */
func RegisterProvider(providerName string, provider SessionProvider) {
	if provider == nil {
		panic("provider is nil")
	}
	providers[providerName] = provider
	fmt.Printf("RegisterProvider name=%s, type=%T\n", providerName, provider)
}

/* 生成唯一的sessionid */
func (manager *SessionManager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

/* 创建session */
func(manager *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request)(session Session){
	manager.mu.Lock()
	defer manager.mu.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == ""{
		//cookie不存在则创建session
		sid := manager.sessionId()
		fmt.Printf("New SessionId=%s\n", sid)
		session, _ = manager.provider.Init(sid)
		//设置cookie值为sessionId
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxLeftTime)}
		http.SetCookie(w, &cookie)
	}else{
		//cookie值转换为sessionId
		sid, _ := url.QueryUnescape(cookie.Value)
		fmt.Printf("Exists SessionId=%s\n", sid)
		session, _ = manager.provider.GetOrInit(sid)
	}

	return
}