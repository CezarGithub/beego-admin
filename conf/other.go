package conf

// Other additional configuration
type Other struct {
	LogAesKey string `mapstructure:"log_aes_key" json:"log_aes_key" yaml:"log_aes_key"` //Log encryption key
}
