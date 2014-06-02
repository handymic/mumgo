package mumgo

import (
	"crypto/tls"
	"fmt"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	CertFile string
	KeyFile  string
}

var certDir = pwd() + "/certs"
var zeroCnf = Config{}
var defaultCnf = Config{
	Username: "mumgo", Host: "localhost", Port: 64738,
	CertFile: certDir + "/mumgo.crt",
	KeyFile:  certDir + "/mumgo.key"}

// Creates a new Config with unassigned fields initialized to defaults
func NewConfig(base ...Config) Config {
	var c Config

	if len(base) == 0 {
		c = Config{}
	} else {
		c = base[0]
	}

	if c.Username == zeroCnf.Username {
		c.Username = defaultCnf.Username
	}

	if c.Host == zeroCnf.Host {
		c.Host = defaultCnf.Host
	}

	if c.Port == zeroCnf.Port {
		c.Port = defaultCnf.Port
	}

	if c.CertFile == zeroCnf.CertFile {
		c.CertFile = defaultCnf.CertFile
	}

	if c.KeyFile == zeroCnf.KeyFile {
		c.KeyFile = defaultCnf.KeyFile
	}

	return c
}

// Get the tls.Config describing *CertFile* & *KeyFile*
func (c *Config) GetTLSConfig(verify bool) (tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(c.CertFile, c.KeyFile)

	if err != nil {
		return tls.Config{}, err
	}

	config := tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: !verify}

	return config, nil
}

// Get the address string described by *Host* & *Port*
func (c *Config) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
