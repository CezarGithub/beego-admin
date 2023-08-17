package conf

// Login settings: background login related settings
type Login struct {
	Token      string `mapstructure:"token" json:"token" yaml:"token"`                //Login token verification
	Captcha    string `mapstructure:"captcha" json:"captcha" yaml:"captcha"`          //Verification code
	Background string `mapstructure:"background" json:"background" yaml:"background"` //Login background
}
