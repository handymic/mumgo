package mumgo

import (
	"os"
	"testing"
)

func newConfig() Config {
	config := Config{
		host:     "localhost",
		port:     64738,
		certFile: os.Getenv("TEST_CRT"),
		keyFile:  os.Getenv("TEST_KEY")}

	return config.ToValid()
}

func closeConn(conn *Conn) {
	conn.Close()
}

func TestNewConnWithAllValid(t *testing.T) {
	conn, err := NewConn(newConfig())
	defer closeConn(&conn)

	expect(t, nil, err)
}

func TestNewConnWithInvalidAddr(t *testing.T) {
	config := newConfig()
	config.host = "missinghost"

	conn, err := NewConn(config)
	defer closeConn(&conn)

	refute(t, nil, err)
}

func TestNewConnWithInvalidCert(t *testing.T) {
	config := newConfig()
	config.certFile = "/tmp/missing.crt"

	conn, err := NewConn(config)
	defer closeConn(&conn)

	refute(t, nil, err)
}

func TestNewConnWithInvalidKey(t *testing.T) {
	config := newConfig()
	config.keyFile = "/tmp/missing.key"

	conn, err := NewConn(config)
	defer closeConn(&conn)

	refute(t, nil, err)
}
