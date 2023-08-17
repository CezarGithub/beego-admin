package conf

// sqlite Related configuration
type Sqlite struct {
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	Version int `mapstructure:"versiom" json:"version" yaml:"version"`
}
