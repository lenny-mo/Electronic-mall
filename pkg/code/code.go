package code

import "sync"

// 业务错误码
const (
	Success       = 200
	InvalidParams = 400
	Error         = 500
	UserExist     = 1001
)

var msgMapping = NewMsgMapping() // 初始化MsgMapping

type MsgMapping struct {
	MsgFlags map[int]string
	sync.RWMutex
}

// NewMsgMapping 初始化MsgMapping
//
// 并发安全
func NewMsgMapping() *MsgMapping {
	return &MsgMapping{
		MsgFlags: map[int]string{
			Success:       "ok",
			InvalidParams: "请求参数错误",
			Error:         "服务器错误",
			UserExist:     "用户已存在",
		},
	}
}

// ReadMap 读取map
//
// 并发安全, 读取失败返回默认值Error
func (m *MsgMapping) ReadMap(key int) string {
	m.RLock()
	defer m.RUnlock()
	if v, ok := m.MsgFlags[key]; ok {
		return v
	}
	return m.MsgFlags[Error]
}

// SetMap 设置map
//
// 并发安全
func (m *MsgMapping) SetMap(key int, val string) {
	m.Lock()
	defer m.Unlock()
	m.MsgFlags[key] = val
}

func GetMsg(code int) string {
	return msgMapping.ReadMap(code) // 并发安全读取map
}
