package config

// Config holds all Config Models
type Config struct {
	App   *AppConfig   `json:"app" yaml:"app"`
	DB    *DBConfig    `json:"db" yaml:"db"`
	Cache *CacheConfig `json:"cache" yaml:"cache"`
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
	Enabled bool   `json:"enabled" yaml:"enabled"`
}

// CacheConfig holds Cache relevant Configuration
type CacheConfig struct {
	Host    string `json:"host" yaml:"host"`
	Port    int    `json:"port" yaml:"port"`
	Enabled bool   `json:"enabled" yaml:"enabled"`
}
