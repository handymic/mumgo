package mumgo

type Config struct {
	username string
	password string
	host     string
	port     int
	certFile string
	keyFile  string
}

var certDir = pwd() + "/certs"
var zeroCnf = Config{}
var defaultCnf = Config{
	username: "mumgo", host: "localhost", port: 64738,
	certFile: certDir + "/mumgo.crt",
	keyFile:  certDir + "/mumgo.key"}

// Fix any zero values & return a new Config.
func (c Config) ToValid() Config {
	if c.username == zeroCnf.username {
		c.username = defaultCnf.username
	}

	if c.host == zeroCnf.host {
		c.host = defaultCnf.host
	}

	if c.port == zeroCnf.port {
		c.port = defaultCnf.port
	}

	if c.certFile == zeroCnf.certFile {
		c.certFile = defaultCnf.certFile
	}

	if c.keyFile == zeroCnf.keyFile {
		c.keyFile = defaultCnf.keyFile
	}

	return c
}
