package gateway

import (
	"sync"

	"github.com/xh3b4sd/anna/gateway/spec"
)

func NewGateway() spec.Gateway {
	g := &gateway{
		Link:   make(chan spec.Signal, 1000),
		Closed: false,
		Mutex:  sync.Mutex{},
	}

	return g
}

type gateway struct {
	Link   chan spec.Signal `json:"link"`
	Closed bool             `json:"closed"`
	Mutex  sync.Mutex       `json:"mutex"`
}

func (g *gateway) Close() {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	g.Closed = true
}

func (g *gateway) Open() {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	g.Closed = false
}

func (g *gateway) ReceiveSignal() (spec.Signal, error) {
	// Note that we can NOT simply defer the call to Mutex.Unlock, because of the
	// Link channel at the end of this function. Link might block until a signal
	// can be read again. In this case Mutex.Unlock is never called and the mutex
	// causes blocking of Gateway.Close, Gateway.Open and Gateway.SendSignal. So
	// we need to explicitly unlock here.
	g.Mutex.Lock()
	if g.Closed {
		g.Mutex.Unlock()
		return nil, maskAny(gatewayClosedError)
	}
	g.Mutex.Unlock()

	return <-g.Link, nil
}

func (g *gateway) SendSignal(signal spec.Signal) error {
	// Note that we can NOT simply defer the call to Mutex.Unlock, because of the
	// Link channel at the end of this function. Link might block until a signal
	// can be sent again. In this case Mutex.Unlock is never called and the mutex
	// causes blocking of Gateway.Close, Gateway.Open and Gateway.ReceiveSignal.
	// So we need to explicitly unlock here.
	g.Mutex.Lock()
	if g.Closed {
		g.Mutex.Unlock()
		return maskAny(gatewayClosedError)
	}
	g.Mutex.Unlock()

	g.Link <- signal
	return nil
}
