package mumgo

import (
	"fmt"
	mumble "github.com/handymic/MumbleProto-go"

	"os"
	"testing"
)

func newConfig() Config {
	return NewConfig(Config{
		Host:     "localhost",
		Port:     64738,
		CertFile: os.Getenv("TEST_CRT"),
		KeyFile:  os.Getenv("TEST_KEY")})
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
	config.Host = "missinghost"

	conn, err := NewConn(config)
	defer closeConn(&conn)

	refute(t, nil, err)
}

func TestNewConnWithInvalidCert(t *testing.T) {
	config := newConfig()
	config.CertFile = "/tmp/missing.crt"

	conn, err := NewConn(config)
	defer closeConn(&conn)

	refute(t, nil, err)
}

func TestNewConnWithInvalidKey(t *testing.T) {
	config := newConfig()
	config.KeyFile = "/tmp/missing.key"

	conn, err := NewConn(config)
	defer closeConn(&conn)

	refute(t, nil, err)
}

func TestWrite(t *testing.T) {
	config := newConfig()

	conn, err := NewConn(config)
	defer closeConn(&conn)

	message := "OMG"
	size, err := conn.Write(&mumble.TextMessage{
		Message: &message})

	fmt.Println("WRITE:", size, err)

}
