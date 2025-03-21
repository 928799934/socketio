package socketio

import (
	"net/http"

	nmem "github.com/928799934/socketio/adaptor/transport/memory"
	eio "github.com/928799934/socketio/engineio"
	siop "github.com/928799934/socketio/protocol"
	siot "github.com/928799934/socketio/transport"
)

// https://socket.io/docs/v4/migrating-from-2-x-to-3-0/
// This is the revision 5 of the Socket.IO protocol, included in socket.io@3.0.0...latest.

type ServerV3 struct {
	inSocketV3

	doBinaryAckPacket func(SocketID, siot.Socket) error

	prev *ServerV2
}

func NewServerV3(opts ...Option) *ServerV3 {
	v3 := &ServerV3{}
	v3.new(opts...)

	v2 := v3.prev
	v1 := v2.prev
	v1.eio = eio.NewServerV4(eio.WithPath(*v1.path)).(eio.EIOServer)
	v1.eio.With(opts...)

	v3.With(opts...)
	return v3
}

func (v3 *ServerV3) new(opts ...Option) Server {
	v3.prev = (&ServerV2{}).new(opts...).(*ServerV2)
	v3.onConnect = make(map[Namespace]onConnectCallbackVersion3)

	v2 := v3.prev
	v1 := v2.prev

	v1.run = runV3(v3)

	v1.transport = nmem.NewInMemoryTransport(siop.NewPacketV5)
	v1.setTransporter(v1.transport)

	v1.protectedEventName = v3ProtectedEventName
	v1.doConnectPacket = doConnectPacketV3(v3)

	v3.doBinaryAckPacket = doBinaryAckPacket(v1)
	v3.inSocketV3.prev = v2.inSocketV2.clone()

	return v3
}

func (v3 *ServerV3) With(opts ...Option) {
	v3.prev.With(opts...)
	for _, opt := range opts {
		opt(v3)
	}
}

func (v3 *ServerV3) In(room Room) inToEmit {
	rtn := v3.clone()
	rtn.setIsServer(true)
	return rtn.In(room)
}

func (v3 *ServerV3) Of(ns Namespace) inSocketV3 {
	rtn := v3.clone()
	rtn.setIsServer(true)
	return rtn.Of(ns)
}

func (v3 *ServerV3) To(room Room) inToEmit {
	rtn := v3.clone()
	rtn.setIsServer(true)
	return rtn.To(room)
}

func (v3 *ServerV3) ServeHTTP(w http.ResponseWriter, r *http.Request) { v3.prev.ServeHTTP(w, r) }
