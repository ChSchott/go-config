package config

// Config holds all Config Models
type Config struct {
	App       *AppConfig       `json:"app" yaml:"app"`
	DB        *DBConfig        `json:"db" yaml:"db"`
	Cache     *CacheConfig     `json:"cache" yaml:"cache"`
	Messaging *MessagingConfig `json:"messaging" yaml:"messaging"`
	Mail      *MailConfig      `json:"mail" yaml:"mail"`
}

// AppConfig holds App relevant Configuration
type AppConfig struct {
	Host    string `json:"host" yaml:"host"`
	Port    int    `json:"port" yaml:"port"`
	Name    string `json:"name" yaml:"name"`
	Version string `json:"version" yaml:"version"`
}

// DBConfig holds Database relevant Configuration
type DBConfig struct {
	Host    string `json:"host" yaml:"host"`
	Port    int    `json:"port" yaml:"port"`
	User    string `json:"user" yaml:"user"`
	Pass    string `json:"pass" yaml:"pass"`
	Name    string `json:"name" yaml:"name"`
	SSL     bool   `json:"ssl" yaml:"ssl"`
	Enabled bool   `json:"enabled" yaml:"enabled"`
}

// CacheConfig holds Cache relevant Configuration
type CacheConfig struct {
	Host    string `json:"host" yaml:"host"`
	Port    int    `json:"port" yaml:"port"`
	Enabled bool   `json:"enabled" yaml:"enabled"`
}

// MessagingConfig holds Message Broker Configuration
type MessagingConfig struct {
	Host    string `json:"host" yaml:"host"`
	Port    int    `json:"port" yaml:"port"`
	Enabled bool   `json:"enabled" yaml:"enabled"`
}

// MailConfig holds Mail Server Configuration
type MailConfig struct {
	Host    string `json:"host" yaml:"host"`
	Port    int    `json:"port" yaml:"port"`
	User    int    `json:"user" yaml:"user"`
	Pass    int    `json:"pass" yaml:"pass"`
	Enabled bool   `json:"enabled" yaml:"enabled"`
}
