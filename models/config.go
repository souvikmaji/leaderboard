package models

// Configurations exported
type Configurations struct {
	Server   ServerConfigurations
	Database DatabaseConfigurations
}

// ServerConfigurations represents http server configurations
type ServerConfigurations struct {
	Host string
	Port int
}

// DatabaseConfigurations represents postgre db credentials
type DatabaseConfigurations struct {
	URL      string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
	Logmode  bool
}
