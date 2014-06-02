package mumgo

import (
	"os"
	"testing"
)

// Should fix missing *host*
func TestFixNilHost(t *testing.T) {
	orig := Config{}
	fixed := orig.ToValid()

	expect(t, defaultCnf.host, fixed.host)
	refute(t, defaultCnf.host, orig.host)
}

// Shouldnt change existing *host*
func TestNoChangeValidHost(t *testing.T) {
	orig := Config{host: "192.168.0.254"}
	fixed := orig.ToValid()

	refute(t, defaultCnf.host, fixed.host)
	refute(t, defaultCnf.host, orig.host)
}

// Should fix missing *port*
func TestFixNilPort(t *testing.T) {
	orig := Config{}
	fixed := orig.ToValid()

	expect(t, defaultCnf.port, fixed.port)
	refute(t, defaultCnf.port, orig.port)
}

// Shouldnt change existing *port*
func TestNoChangeValidPort(t *testing.T) {
	orig := Config{port: 56789}
	fixed := orig.ToValid()

	refute(t, defaultCnf.port, fixed.port)
	refute(t, defaultCnf.port, orig.port)
}

// Should fix missing *username*
func TestFixNilUsername(t *testing.T) {
	orig := Config{}
	fixed := orig.ToValid()

	expect(t, defaultCnf.username, fixed.username)
	refute(t, defaultCnf.username, orig.username)
}

// Shouldnt change existing *username*
func TestNoChangeValidUsername(t *testing.T) {
	orig := Config{username: "alice"}
	fixed := orig.ToValid()

	refute(t, defaultCnf.username, fixed.username)
	refute(t, defaultCnf.username, orig.username)
}

// Should fix missing *certFile*
func TestFixNilCertFile(t *testing.T) {
	orig := Config{}
	fixed := orig.ToValid()

	expect(t, defaultCnf.certFile, fixed.certFile)
	refute(t, defaultCnf.certFile, orig.certFile)
}

// Shouldnt change existing *certFile*
func TestNoChangeValidCertFile(t *testing.T) {
	orig := Config{certFile: "~/.mumgo/mumgo.crt"}
	fixed := orig.ToValid()

	refute(t, defaultCnf.certFile, fixed.certFile)
	refute(t, defaultCnf.certFile, orig.certFile)
}

// Should fix missing *keyFile*
func TestFixNilKeyFile(t *testing.T) {
	orig := Config{}
	fixed := orig.ToValid()

	expect(t, defaultCnf.keyFile, fixed.keyFile)
	refute(t, defaultCnf.keyFile, orig.keyFile)
}

// Shouldnt change existing *keyFile*
func TestNoChangeValidKeyFile(t *testing.T) {
	orig := Config{keyFile: "~/.mumgo/mumgo.crt"}
	fixed := orig.ToValid()

	refute(t, defaultCnf.keyFile, fixed.keyFile)
	refute(t, defaultCnf.keyFile, orig.keyFile)
}

// Shouldnt fix missing *password*
func TestNoFixNilPassword(t *testing.T) {
	orig := Config{}
	fixed := orig.ToValid()

	expect(t, zeroCnf.password, fixed.password)
	expect(t, zeroCnf.password, orig.password)
}

// Should succeed in loading valid cert
func TestGetTLSConfigWithValidCertAndKeyFiles(t *testing.T) {
	certFile, keyFile := os.Getenv("TEST_CRT"), os.Getenv("TEST_KEY")

	config := Config{keyFile: keyFile, certFile: certFile}
	tlsConf, err := config.GetTLSConfig(true)

	expect(t, 1, len(tlsConf.Certificates))
	expect(t, nil, err)
}

// Should fail in loading missing cert
func TestGetTLSConfigWithMissingCertAndKeyFiles(t *testing.T) {
	certFile, keyFile := "/tmp/missing.crt", "/tmp/missing.key"

	config := Config{keyFile: keyFile, certFile: certFile}
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

	config := Config{keyFile: keyFile, certFile: certFile}
	tlsConf, err := config.GetTLSConfig(true)

	expect(t, 0, len(tlsConf.Certificates))
	refute(t, nil, err)
}
