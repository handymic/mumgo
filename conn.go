package mumgo

import (
	proto "code.google.com/p/goprotobuf/proto"
	mumble "github.com/handymic/MumbleProto-go"

	"bytes"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"reflect"
)

// Message to type mappings to ease building the appropriate
// protobuf message on the fly
var messageTypes = map[reflect.Type]int{
	reflect.TypeOf(&mumble.Version{}):             0,
	reflect.TypeOf(&mumble.UDPTunnel{}):           1,
	reflect.TypeOf(&mumble.Authenticate{}):        2,
	reflect.TypeOf(&mumble.Ping{}):                3,
	reflect.TypeOf(&mumble.Reject{}):              4,
	reflect.TypeOf(&mumble.ServerSync{}):          5,
	reflect.TypeOf(&mumble.ChannelRemove{}):       6,
	reflect.TypeOf(&mumble.ChannelState{}):        7,
	reflect.TypeOf(&mumble.UserRemove{}):          8,
	reflect.TypeOf(&mumble.UserState{}):           9,
	reflect.TypeOf(&mumble.BanList{}):             10,
	reflect.TypeOf(&mumble.TextMessage{}):         11,
	reflect.TypeOf(&mumble.PermissionDenied{}):    12,
	reflect.TypeOf(&mumble.ACL{}):                 13,
	reflect.TypeOf(&mumble.QueryUsers{}):          14,
	reflect.TypeOf(&mumble.CryptSetup{}):          15,
	reflect.TypeOf(&mumble.ContextActionModify{}): 16,
	reflect.TypeOf(&mumble.ContextAction{}):       17,
	reflect.TypeOf(&mumble.UserList{}):            18,
	reflect.TypeOf(&mumble.VoiceTarget{}):         19,
	reflect.TypeOf(&mumble.PermissionQuery{}):     20,
	reflect.TypeOf(&mumble.CodecVersion{}):        21,
	reflect.TypeOf(&mumble.UserStats{}):           22,
	reflect.TypeOf(&mumble.RequestBlob{}):         23,
	reflect.TypeOf(&mumble.ServerConfig{}):        24,
	reflect.TypeOf(&mumble.SuggestConfig{}):       25}

//
type Conn struct {
	conn        *tls.Conn
	initialized bool
}

// Creates a new initialized connection
func NewConn(config Config) (Conn, error) {
	cert, err := config.LoadCert()

	if err != nil {
		return Conn{}, err
	}

	addr := fmt.Sprint(config.host, ":", config.port)
	tlsConf := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true} // TODO: fix hardcoding

	conn, err := tls.Dial("tcp", addr, tlsConf)

	if err != nil {
		return Conn{}, err
	}

	return Conn{conn: conn, initialized: true}, nil
}

// Close the connection
func (c *Conn) Close() {
	if c.initialized {
		c.conn.Close()
	}
}

// Write protobuf message to connection, taking care of header construction
func (c *Conn) Write(message proto.Message) (int, error) {
	payload, err := proto.Marshal(message)

	if err != nil {
		return -1, err
	}

	var buffer bytes.Buffer
	var chunk []byte

	// Prepare *type* prefix
	mtype := messageTypes[reflect.TypeOf(message)]
	chunk = make([]byte, 2)
	binary.PutVarint(chunk, int64(mtype))
	_, err = buffer.Write(chunk)

	if err != nil {
		return -1, err
	}

	// Prepare *size* prefix
	chunk = make([]byte, 4)
	binary.PutVarint(chunk, int64(len(payload)))
	_, err = buffer.Write(chunk)

	if err != nil {
		return -1, err
	}

	// Prepare *payload* body
	_, err = buffer.Write(payload)

	if err != nil {
		return -1, err
	}

	return c.conn.Write(buffer.Bytes())
}
