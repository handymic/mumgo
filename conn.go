package mumgo

import (
	"crypto/tls"
	"fmt"
)

type Conn struct {
	conn *tls.Conn
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
		InsecureSkipVerify: true}

	conn, err := tls.Dial("tcp", addr, tlsConf)

	if err != nil {
		return Conn{}, err
	}

	return Conn{conn: conn}, nil
}

//