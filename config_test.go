package mumgo

import "testing"

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

// Shouldnt fix missing *password*
func TestNoFixNilPassword(t *testing.T) {
	orig := Config{}
	fixed := orig.ToValid()

	expect(t, zeroCnf.password, fixed.password)
	expect(t, zeroCnf.password, orig.password)
}
