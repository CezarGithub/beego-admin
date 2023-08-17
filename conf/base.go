package conf

//Basic information settings in the background
type Base struct {
	Name            string `mapstructure:"name" json:"name" yaml:"name"`                                     //Backstage name
	ShortName       string `mapstructure:"short_name" json:"short_name" yaml:"short_name"`                   //Backstage abbreviation
	Author          string `mapstructure:"author" json:"author" yaml:"author"`                               //Backstage author
	Version         string `mapstructure:"version" json:"version" yaml:"version"`                            //Background version
	Link            string `mapstructure:"link" json:"link" yaml:"link"`                                     //footer link address
	PasswordWarning string `mapstructure:"password_warning" json:"password_warning" yaml:"password_warning"` //Default password warning
	ShowNotice      string `mapstructure:"show_notice" json:"show_notice" yaml:"show_notice"`                //Whether to display prompt information
	NoticeContent   string `mapstructure:"notice_content" json:"notice_content" yaml:"notice_content"`       //Prompt message content
}
