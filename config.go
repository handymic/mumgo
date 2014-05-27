package mumgo

type Config struct {
	username string
	password string
	host     string
	port     int
}

// Zero config
var zeroCnf = Config{}

// Default config
var defaultCnf = Config{
	username: "mumgo", host: "localhost", port: 64738}

// Fix the zero values in the provided config
func (c Config) FixNils() Config {
	if c.username == zeroCnf.username {
		c.username = defaultCnf.username
	}

	if c.host == zeroCnf.host {
		c.host = defaultCnf.host
	}

	if c.port == zeroCnf.port {
		c.port = defaultCnf.port
	}

	return c
}
