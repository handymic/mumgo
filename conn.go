package mumgo

import (
	proto "code.google.com/p/goprotobuf/proto"
	mumble "github.com/handymic/MumbleProto-go"

	"bytes"
	"crypto/tls"
	"reflect"
)

// Message to type mappings to ease building the appropriate
// protobuf message on the fly
var messageTypes = map[reflect.Type]uint16{
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
	conn      *tls.Conn
	connected bool
}

// Creates a new connected connection
func NewConn(config Config) (Conn, error) {

	// Gets tls.Config described by config.CertFile & config.KeyFile
	tlsConf, err := config.GetTLSConfig(false)
	if err != nil {
		return Conn{}, err
	}

	// Inits tls connection
	tlsConn, err := tls.Dial("tcp", config.GetAddr(), &tlsConf)
	if err != nil {
		return Conn{}, err
	}

	// Inits client connection
	conn := Conn{conn: tlsConn}
	if err := conn.init(config); err != nil {
		return Conn{}, err
	}

	return conn, nil
}

// Close the connection
func (c *Conn) Close() error {
	if c.connected {
		if err := c.conn.Close(); err != nil {
			return err
		}

		c.connected = false
	}

	return nil
}

// Write protobuf message to connection, taking care of header construction
func (c *Conn) Write(message proto.Message) (int, error) {
	payload, err := proto.Marshal(message)

	if err != nil {
		return -1, err
	}

	var buf bytes.Buffer

	// Prepare *type* prefix
	mtype := messageTypes[reflect.TypeOf(message)]
	if _, err := buf.Write(uint16tbs(mtype)); err != nil {
		return -1, err
	}

	// Prepare *size* prefix
	size := uint32(len(payload))
	if _, err := buf.Write(uint32tbs(size)); err != nil {
		return -1, err
	}

	// Prepare *payload* body
	if _, err := buf.Write(payload); err != nil {
		return -1, err
	}

	return c.conn.Write(buf.Bytes())
}

// Initializes the connection
func (c *Conn) init(config Config) error {
	if err := c.exchangeVersion(); err != nil {
		return err
	}
	if err := c.authenticate(config); err != nil {
		return err
	}

	c.connected = true
	return nil
}

// Exchanges version info
func (c *Conn) exchangeVersion() error {
	major, minor, patch := 1, 2, 5

	version := uint32((major << 16) | (minor << 8) | (patch & 0xFF))

	// TODO: fix hardcoding !!!
	release := "mumgo 0.0.1"
	os, osVersion := "linux", "#1 SMP PREEMPT Tue May 13 16:41:39 CEST 2014"

	_, err := c.Write(&mumble.Version{
		Version:   &version,
		Release:   &release,
		Os:        &os,
		OsVersion: &osVersion})

	return err
}

// Performs authentication
func (c *Conn) authenticate(config Config) error {
	opus := true

	_, err := c.Write(&mumble.Authenticate{
		Username: &config.Username,
		Password: &config.Password,
		Opus:     &opus})

	return err
}
