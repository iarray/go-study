package provider

import(
	"../session"
	"fmt"
	"time"
	"sync"
)

func init(){
	provider := MemoryProvider{}
	session.RegisterProvider("memory", session.SessionProvider(&provider))
}

var sessions = make(map[string]Session)

type MemoryProvider struct{

}

type Session struct {
	_id string
	_createTime int64
	data map[interface{}]interface{}
	_lock sync.RWMutex
}

func (m * MemoryProvider)Init(sid string) (session.Session, error) {
	return m.GetOrInit(sid)
}

func  (* MemoryProvider)GetOrInit(sid string) (session.Session, error){
	ses, ok := sessions[sid]
	if !ok {
		sessions[sid] = Session{_id: sid, _createTime:time.Now().Unix()}
	}
	return session.Session(&ses), nil
}

func  (* MemoryProvider)Remove(sid string) error{
	_, ok := sessions[sid]
	if !ok {
		fmt.Errorf("session id not exists")
	}
	delete(sessions, sid)
	return nil
}
func (* MemoryProvider)GC(maxLeftTime int64){
	now := time.Now().Unix()
	for k,v := range sessions {
		if now - v._createTime > maxLeftTime{
			defer delete(sessions, k)
		}
	}
}

func (s *Session)Get(key interface{}) interface{}{
	if key == nil {
		return nil
	}
	
	s._lock.RLock()
	defer s._lock.RUnlock()

	if s.data == nil{
		return nil
	}
	
	val, ok := s.data[key]
	if !ok{
		return nil
	}

	return val
}

func (s *Session)Set(key, value interface{}){
	if key == nil || value == nil{
		panic("key or value is nil")
	}
	s._lock.Lock()
	defer s._lock.Unlock()
	if s.data == nil{
		s.data = make(map[interface{}]interface{})
	}
	s.data[key] = value
}

func (s *Session)Remove(key interface{}){
	if key == nil {
		panic("key is nil")
	}
	if s.data == nil{
		return
	}
	s._lock.Lock()
	defer s._lock.Unlock()
	delete(s.data, key)
}

func (m *Session)Id()string{
	return m._id
}
