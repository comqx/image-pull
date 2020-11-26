package tcpServer

import (
	"net"
	"sync"
)

type SecurityMap struct {
	registeredMap map[string]map[string]net.Conn
	sync.Mutex    //相对效率更高
}

// 加锁写入数据
func (s *SecurityMap) add(s1, s2 string, n net.Conn) {
	s.Lock()
	defer s.Unlock()
	if len(s.registeredMap) == 0 {
		s.registeredMap[s1] = make(map[string]net.Conn, 5)
	}
	s.registeredMap[s1][s2] = n
}

// 加锁读数据
func (s *SecurityMap) get(s1 string) map[string]net.Conn {
	//u.RLock()
	//defer u.RUnlock()
	s.Lock()
	defer s.Unlock()
	if valMap, ok := s.registeredMap[s1]; ok {
		return valMap
	}
	return nil
}
