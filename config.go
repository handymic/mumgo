package mumgo

// Zero values for comparison
var stringZero string
var intZero int

type Config struct {
	username string
	password string
	host     string
	port     int
}

// Fix the zero values in the provided config
func (c *Config) FixNil(config Config) Config {
	if config.username == stringZero {
		config.username = "mumgo"
	}

	if config.host == stringZero {
		config.host = "localhost"
	}

	if config.port == intZero {
		config.port = 64738
	}

	return config
}
