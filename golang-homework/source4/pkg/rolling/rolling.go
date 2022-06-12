package rolling

import (
	"sync"
	"time"
)


type StatusMapAction interface {
	ModifyStatus(data *StatusNode)
	GetStatus() *StatusNode
	deleteStatus()
	GetAllStatus() *map[int64]*StatusNode
}

func (s *StatusMap) GetAllStatus() map[int64]*StatusNode {
	return s.nodeMap
}

type StatusNodeAction interface {
	Add(success, failure, timeout, reject int)
}

func (s *StatusNode) Add(success, failure, timeout, reject int32) {
	s.success += success
	s.failure += failure
	s.timeout += timeout
	s.reject += reject
	//if success != 0 {
	//	atomic.AddInt32(&s.success, success)
	//}
	//if failure != 0 {
	//	atomic.AddInt32(&s.failure, failure)
	//}
	//if timeout != 0 {
	//	atomic.AddInt32(&s.timeout, timeout)
	//}
	//if reject != 0 {
	//	atomic.AddInt32(&s.reject, reject)
	//}
}

func (s *StatusMap) GetStatus() (int64, *StatusNode) {
	t := time.Now().Unix()
	var node *StatusNode
	var ok bool
	if node, ok = s.nodeMap[t]; !ok {
		node = NewStatusNode()
		s.nodeMap[t] = node
	}
	return t, node
}

func (s *StatusMap) deleteStatus() {
	// 10s以下的状态清除
	t := time.Now().Unix() - 10
	for timestamp := range s.nodeMap {
		if timestamp < t {
			//s.mutex.Lock()
			//defer s.mutex.Unlock()
			delete(s.nodeMap, timestamp)
		}
	}
}

func (s *StatusMap) ModifyStatus(data *StatusNode){
	_, node := s.GetStatus()
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if data.success != 0 {
		node.success += data.success
	}
	if data.failure != 0 {
		node.failure += data.failure
	}
	if data.reject != 0 {
		node.reject += data.reject
	}
	if data.timeout != 0 {
		node.timeout += data.timeout
	}
	//node = &res
	//s.nodeMap[timestamp] = node
}

type StatusMap struct {
	nodeMap map[int64]*StatusNode
	mutex *sync.RWMutex
}

type StatusNode struct {
	// 成功量、失败量、超时量、拒绝量
	success, failure, timeout, reject int32
}

func NewStatusNode() *StatusNode {
	return &StatusNode{
		0,0,0,0,
	}
}

func NewStatusMap() *StatusMap {
	statusMap := make(map[int64]*StatusNode, 10)
	return &StatusMap{
		statusMap,
		&sync.RWMutex{},
	}
}