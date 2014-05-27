package mumgo

type Config struct {
	username string
	password string
	host     string
	port     int
}

var zeroCnf = Config{}
var defaultCnf = Config{
	username: "mumgo", host: "localhost", port: 64738}

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

	return c
}
