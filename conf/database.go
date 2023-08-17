package conf

// sqlite Related configuration
type Database struct {
	Type string `mapstructure:"type" json:"type" yaml:"type"`
	JsonFolder string `mapstructure:"jsonfolder" json:"jsonfolder" yaml:"jsonfolder"`
}
