package mumgo

import (
	"os"
	"testing"
)

// Should fix missing *host*
func TestFixNilHost(t *testing.T) {
	orig := Config{}
	fixed := orig.ToValid()

	expect(t, defaultCnf.Host, fixed.Host)
	refute(t, defaultCnf.Host, orig.Host)
}

// Shouldnt change existing *host*
func TestNoChangeValidHost(t *testing.T) {
	orig := Config{Host: "192.168.0.254"}
	fixed := orig.ToValid()

	refute(t, defaultCnf.Host, fixed.Host)
	refute(t, defaultCnf.Host, orig.Host)
}

// Should fix missing *port*
func TestFixNilPort(t *testing.T) {
	orig := Config{}
	fixed := orig.ToValid()

	expect(t, defaultCnf.Port, fixed.Port)
	refute(t, defaultCnf.Port, orig.Port)
}

// Shouldnt change existing *port*
func TestNoChangeValidPort(t *testing.T) {
	orig := Config{Port: 56789}
	fixed := orig.ToValid()

	refute(t, defaultCnf.Port, fixed.Port)
	refute(t, defaultCnf.Port, orig.Port)
}

// Should fix missing *Username*
func TestFixNilUsername(t *testing.T) {
	orig := Config{}
	fixed := orig.ToValid()

	expect(t, defaultCnf.Username, fixed.Username)
	refute(t, defaultCnf.Username, orig.Username)
}

// Shouldnt change existing *Username*
func TestNoChangeValidUsername(t *testing.T) {
	orig := Config{Username: "alice"}
	fixed := orig.ToValid()

	refute(t, defaultCnf.Username, fixed.Username)
	refute(t, defaultCnf.Username, orig.Username)
}

// Should fix missing *certFile*
func TestFixNilCertFile(t *testing.T) {
	orig := Config{}
	fixed := orig.ToValid()

	expect(t, defaultCnf.CertFile, fixed.CertFile)
	refute(t, defaultCnf.CertFile, orig.CertFile)
}

// Shouldnt change existing *certFile*
func TestNoChangeValidCertFile(t *testing.T) {
	orig := Config{CertFile: "~/.mumgo/mumgo.crt"}
	fixed := orig.ToValid()

	refute(t, defaultCnf.CertFile, fixed.CertFile)
	refute(t, defaultCnf.CertFile, orig.CertFile)
}

// Should fix missing *keyFile*
func TestFixNilKeyFile(t *testing.T) {
	orig := Config{}
	fixed := orig.ToValid()

	expect(t, defaultCnf.KeyFile, fixed.KeyFile)
	refute(t, defaultCnf.KeyFile, orig.KeyFile)
}

// Shouldnt change existing *keyFile*
func TestNoChangeValidKeyFile(t *testing.T) {
	orig := Config{KeyFile: "~/.mumgo/mumgo.crt"}
	fixed := orig.ToValid()

	refute(t, defaultCnf.KeyFile, fixed.KeyFile)
	refute(t, defaultCnf.KeyFile, orig.KeyFile)
}

// Shouldnt fix missing *password*
func TestNoFixNilPassword(t *testing.T) {
	orig := Config{}
	fixed := orig.ToValid()

	expect(t, zeroCnf.Password, fixed.Password)
	expect(t, zeroCnf.Password, orig.Password)
}

// Should succeed in loading valid cert
func TestGetTLSConfigWithValidCertAndKeyFiles(t *testing.T) {
	certFile, keyFile := os.Getenv("TEST_CRT"), os.Getenv("TEST_KEY")

	config := Config{KeyFile: keyFile, CertFile: certFile}
	tlsConf, err := config.GetTLSConfig(true)

	expect(t, 1, len(tlsConf.Certificates))
	expect(t, nil, err)
}

// Should fail in loading missing cert
func TestGetTLSConfigWithMissingCertAndKeyFiles(t *testing.T) {
	certFile, keyFile := "/tmp/missing.crt", "/tmp/missing.key"

	config := Config{KeyFile: keyFile, CertFile: certFile}
	tlsConf, err := config.GetTLSConfig(true)

	expect(t, 0, len(tlsConf.Certificates))
	refute(t, nil, err)
}

// Should fail in loading invalid cert
func TestGetTLSConfigWithInvalidCertAndKeyFiles(t *testing.T) {
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
