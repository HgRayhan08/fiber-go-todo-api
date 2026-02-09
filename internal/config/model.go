package config

type Config struct {
	Server   Server
	Database Database
	Jwt      Jwt
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Tz       string
}

type Jwt struct {
	Key    string
	Expire int
}
