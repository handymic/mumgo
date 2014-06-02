package mumgo

import (
	"fmt"
	"os"
	"testing"
)

// Should fix missing *host*
func TestNewConfig_FixNilHost(t *testing.T) {
	config := NewConfig()
	expect(t, defaultCnf.Host, config.Host)
}

// Shouldnt change existing *host*
func TestNewConfig_NoChangeValidHost(t *testing.T) {
	valid := "192.168.0.254"
	config := NewConfig(Config{Host: valid})
	expect(t, valid, config.Host)
}

// Should fix missing *port*
func TestNewConfig_FixNilPort(t *testing.T) {
	config := NewConfig()
	expect(t, defaultCnf.Port, config.Port)
}

// Shouldnt change existing *port*
func TestNewConfig_NoChangeValidPort(t *testing.T) {
	valid := 56789
	config := NewConfig(Config{Port: valid})
	expect(t, valid, config.Port)
}

// Should fix missing *Username*
func TestNewConfig_FixNilUsername(t *testing.T) {
	config := NewConfig()
	expect(t, defaultCnf.Username, config.Username)
}

// Shouldnt change existing *Username*
func TestNewConfig_NoChangeValidUsername(t *testing.T) {
	valid := "alice"
	config := NewConfig(Config{Username: valid})
	expect(t, valid, config.Username)
}

// Should fix missing *certFile*
func TestNewConfig_FixNilCertFile(t *testing.T) {
	config := NewConfig()
	expect(t, defaultCnf.CertFile, config.CertFile)
}

// Shouldnt change existing *certFile*
func TestNewConfig_NoChangeValidCertFile(t *testing.T) {
	valid := "~/.mumgo/mumbot.crt"
	config := NewConfig(Config{CertFile: valid})
	expect(t, valid, config.CertFile)
}

// Should fix missing *keyFile*
func TestNewConfig_FixNilKeyFile(t *testing.T) {
	config := NewConfig()
	expect(t, defaultCnf.KeyFile, config.KeyFile)
}

// Shouldnt change existing *keyFile*
func TestNewConfig_NoChangeValidKeyFile(t *testing.T) {
	valid := "~/.mumgo/mumbot.crt"
	config := NewConfig(Config{KeyFile: valid})
	expect(t, valid, config.KeyFile)
}

// Shouldnt fix missing *password*
func TestNewConfig_NoFixNilPassword(t *testing.T) {
	config := NewConfig()
	expect(t, zeroCnf.Password, config.Password)
}

// Should succeed in loading valid cert
func TestGetTLSConfig_WithValidCertAndKeyFile(t *testing.T) {
	certFile, keyFile := os.Getenv("TEST_CRT"), os.Getenv("TEST_KEY")

	config := Config{KeyFile: keyFile, CertFile: certFile}
	tlsConf, err := config.GetTLSConfig(true)

	expect(t, 1, len(tlsConf.Certificates))
	expect(t, nil, err)
}

// Should fail in loading missing cert
func TestGetTLSConfig_WithMissingCertAndKeyFile(t *testing.T) {
	certFile, keyFile := "/tmp/missing.crt", "/tmp/missing.key"

	config := Config{KeyFile: keyFile, CertFile: certFile}
	tlsConf, err := config.GetTLSConfig(true)

	expect(t, 0, len(tlsConf.Certificates))
	refute(t, nil, err)
}

// Should fail in loading invalid cert
func TestGetTLSConfig_WithInvalidCertAndKeyFile(t *testing.T) {
	certFile := os.TempDir() + "/invalid.crt"
	keyFile := os.TempDir() + "/invalid.key"

	os.Create(certFile)
	os.Create(keyFile)

	defer func() { // trash created tmp files prior eventual return
		os.Remove(certFile)
		os.Remove(keyFile)
	}()

	config := Config{KeyFile: keyFile, CertFile: certFile}
	tlsConf, err := config.GetTLSConfig(true)

	expect(t, 0, len(tlsConf.Certificates))
	refute(t, nil, err)
}

func TestGetAddr(t *testing.T) {
	host, port := "aweso.me", 99999
	config := NewConfig(Config{Host: host, Port: port})
	expect(t, fmt.Sprint(host, ":", port), config.GetAddr())
}
