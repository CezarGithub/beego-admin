package conf

// Upload attachment configuration
type Attachment struct {
	ThumbPath    string `mapstructure:"thumb_path" json:"thumb_path" yaml:"thumb_path"`          //Thumbnail path
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                            //Upload directory configuration (relative to the root directory)
	Url          string `mapstructure:"url" json:"url" yaml:"url"`                               //url (relative to the web directory)
	ValidateSize string `mapstructure:"validate_size" json:"validate_size" yaml:"validate_size"` //The default is not more than 50mb
	ValidateExt  string `mapstructure:"validate_ext" json:"validate_ext" yaml:"validate_ext"`    //url (relative to the web directory)
}
