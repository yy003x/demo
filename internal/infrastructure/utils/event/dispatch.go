package event

import (
	"context"
	"fmt"
	"reflect"
	"sync"
)

type Event interface {
	Type() string
	Source() interface{}
	Value() interface{}
}

type Listener interface {
	Handle(context.Context, Event) error
}

type FuncListener func(ctx context.Context, e Event) error

func (fn FuncListener) Handle(ctx context.Context, e Event) error {
	return fn(ctx, e)
}

type Dispatcher interface {
	AddEventListener(string, Listener)
	RemoveEventListener(string, Listener)
	DispatchEvent(context.Context, Event)
	AsyncDispatchEvent(context.Context, Event)
	HasEventListener(string, Listener) bool
	GetEventListeners(typ string) []Listener
}

type event struct {
	typ    string
	source interface{}
	value  interface{}
}

func NewEvent(typ string, value interface{}) *event {
	return &event{
		typ:   typ,
		value: value,
	}
}

func (e *event) Type() string {
	return e.typ
}

func (e *event) Source() interface{} {
	return e.source
}

func (e *event) Value() interface{} {
	return e.value
}

type eventListeners []Listener

type eventDispatcher struct {
	sync.RWMutex
	source    interface{}
	listeners map[string]eventListeners
}

func NewEventDispatcher(source interface{}) Dispatcher {
	return &eventDispatcher{
		source:    source,
		listeners: make(map[string]eventListeners),
	}
}

func (d *eventDispatcher) AddEventListener(typ string, listener Listener) {
	d.Lock()
	defer d.Unlock()
	d.listeners[typ] = append(d.listeners[typ], listener)
}

func (d *eventDispatcher) RemoveEventListener(typ string, listener Listener) {
	d.Lock()
	defer d.Unlock()

	ptr := reflect.ValueOf(listener).Pointer()
	listeners := d.listeners[typ]
	for i, l := range listeners {
		if reflect.ValueOf(l).Pointer() == ptr {
			d.listeners[typ] = append(listeners[:i], listeners[i+1:]...)
		}
	}
}

func (d *eventDispatcher) DispatchEvent(ctx context.Context, e Event) {
	d.RLock()
	defer d.RUnlock()
	fmt.Println(e)
	if ev, ok := e.(*event); ok {
		ev.source = d.source
	}
	listeners, ok := d.listeners[e.Type()]
	if !ok {
		fmt.Println("none dispatch")
		return
	}

	for _, l := range listeners {
		fmt.Println(e.Type())
		l.Handle(ctx, e)
	}
}

func (d *eventDispatcher) AsyncDispatchEvent(ctx context.Context, e Event) {
	d.RLock()
	defer d.RUnlock()

	if ev, ok := e.(*event); ok {
		ev.source = d.source
	}
	for _, l := range d.listeners[e.Type()] {
		go l.Handle(ctx, e)
	}
}

func (d *eventDispatcher) HasEventListener(typ string, listener Listener) bool {
	d.Lock()
	defer d.Unlock()

	ptr := reflect.ValueOf(listener).Pointer()
	listeners := d.listeners[typ]
	for _, l := range listeners {
		if reflect.ValueOf(l).Pointer() == ptr {
			return true
		}
	}
	return false
}

func (d *eventDispatcher) GetEventListeners(typ string) []Listener {
	d.RLock()
	defer d.RUnlock()
	return d.listeners[typ]
}
