package memory

import(
	"../../../session"
	"fmt"
	"time"
	"sync"
)

func init(){
	provider := MemoryProvider{}
	session.RegisterProvider("memory", &provider)
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
		sessions[sid] = Session{_id: sid, _createTime:time.Now().Unix(), data: make(map[interface{}]interface{})}
		ses,_ = sessions[sid]
	}
	return &ses, nil
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
  
	val, ok := s.data[key]
	fmt.Println(val, ok)
	if !ok{
		return nil
	}

	return val
}

func (session *Session)Set(key, value interface{}){
	if key == nil || value == nil{
		panic("key or value is nil")
	}
	session._lock.Lock()
	defer session._lock.Unlock()
	 
	session.data[key] = value
	fmt.Printf("Set key=%s, value=%s\n", key, value)
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
