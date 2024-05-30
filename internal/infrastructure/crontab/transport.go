package crontab

import (
	"github.com/go-kratos/kratos/v2/transport"
)

const (
	KIND_CRON = "cron"
)

type header map[string]string

// Add implements transport.Header.
func (hc header) Add(key string, value string) {
	panic("unimplemented")
}

// Values implements transport.Header.
func (hc header) Values(key string) []string {
	panic("unimplemented")
}

// Get returns the value associated with the passed key.
func (hc header) Get(key string) string {
	return hc[key]
}

// Set stores the key-value pair.
func (hc header) Set(key string, value string) {
	hc[key] = value
}

// Keys lists the keys stored in this carrier.
func (hc header) Keys() []string {
	keys := make([]string, 0, len(hc))
	for k := range hc {
		keys = append(keys, k)
	}
	return keys
}

type Transport struct {
	reqHeader   header
	replyHeader header
}

// Kind returns the transport kind.
func (tr *Transport) Kind() transport.Kind {
	return KIND_CRON
}

// Endpoint returns the transport endpoint.
func (tr *Transport) Endpoint() string {
	return ""
}

// Operation returns the transport operation.
func (tr *Transport) Operation() string {
	return ""
}

// RequestHeader returns the request header.
func (tr *Transport) RequestHeader() transport.Header {
	return tr.reqHeader
}

// ReplyHeader returns the reply header.
func (tr *Transport) ReplyHeader() transport.Header {
	return tr.replyHeader
}
