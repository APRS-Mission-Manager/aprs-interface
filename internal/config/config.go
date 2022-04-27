package config

type Config struct {
	APRS   APRSConfig
	Amazon AmazonConfig
}

type APRSConfig struct {
	Server           string
	Port             int
	Username         string
	Password         string
	CallsignPatterns []string
}

type AmazonConfig struct {
	DBAPRSLog DatabaseConfig
}

type DatabaseConfig struct {
	Name string
}
