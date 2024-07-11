package event

import (
	"sync"
	"time"
)

// Event 事件结构体
type Event struct {
	// 事件类型
	Type string
	// 事件数据
	Data interface{}
	// 事件发生的时间
	Time time.Time
}

// EventHandler 事件处理函数类型
type EventHandler struct {
	// 事件处理函数
	Handle func(event Event)
	// 函数元数据，用于唯一标识函数
	Metadata string
	// 是否为一次性事件处理函数
	once bool
}

// EventEmitter 事件发射器接口
type EventEmitter interface {
	// On 注册事件处理函数
	On(eventType string, handler EventHandler)
	// Once 注册一次性事件处理函数
	Once(eventType string, handler EventHandler)
	// Off 注销事件处理函数
	Off(eventType string, handler EventHandler)
	// Emit 触发事件
	Emit(eventType string, data interface{})
}

// SimpleEventEmitter 简单的基于map的事件发射器实现
type SimpleEventEmitter struct {
	events map[string]map[*EventHandler]bool
	mu     sync.RWMutex
}

// NewSimpleEventEmitter 创建一个简单的基于map的事件发射器
func NewSimpleEventEmitter() *SimpleEventEmitter {
	return &SimpleEventEmitter{
		events: make(map[string]map[*EventHandler]bool),
	}
}

// On 注册事件处理函数
func (e *SimpleEventEmitter) On(eventType string, handler EventHandler) {
	e.mu.Lock()
	if _, ok := e.events[eventType]; !ok {
		e.events[eventType] = make(map[*EventHandler]bool)
	}
	e.events[eventType][&handler] = true
	e.mu.Unlock()
}

// Once 注册一次性事件处理函数
func (e *SimpleEventEmitter) Once(eventType string, handler EventHandler) {
	handler.once = true
	e.mu.Lock()
	if _, ok := e.events[eventType]; !ok {
		e.events[eventType] = make(map[*EventHandler]bool)
	}
	e.events[eventType][&handler] = true
	e.mu.Unlock()
}

// Off 注销事件处理函数
func (e *SimpleEventEmitter) Off(eventType string, handler EventHandler) {
	e.mu.Lock()
	if _, ok := e.events[eventType]; !ok {
		e.mu.Unlock()
		return
	}
	delete(e.events[eventType], &handler)
	e.mu.Unlock()
}

// Emit 触发事件
func (e *SimpleEventEmitter) Emit(eventType string, data interface{}) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	if _, ok := e.events[eventType]; !ok {
		return
	}

	event := Event{
		Type: eventType,
		Data: data,
		Time: time.Now(),
	}

	for h, _ := range e.events[eventType] {
		h.Handle(event)
	}
}
