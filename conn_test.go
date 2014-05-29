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

func TestNewConnWithAllValid(t *testing.T) {
	_, err := NewConn(newConfig())

	expect(t, nil, err)
}

func TestNewConnWithInvalidAddr(t *testing.T) {
	config := newConfig()
	config.host = "missinghost"
	_, err := NewConn(config)

	refute(t, nil, err)
}

func TestNewConnWithInvalidCert(t *testing.T) {
	config := newConfig()
	config.certFile = "/tmp/missing.crt"
	_, err := NewConn(config)

	refute(t, nil, err)
}

func TestNewConnWithInvalidKey(t *testing.T) {
	config := newConfig()
	config.keyFile = "/tmp/missing.key"
	_, err := NewConn(config)

	refute(t, nil, err)
}
