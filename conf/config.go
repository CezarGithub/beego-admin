package conf

type Server struct {
	Base       Base       `mapstructure:"base" json:"base" yaml:"base"`
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Sqlite     Sqlite     `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	Redis      Mysql      `mapstructure:"redis" json:"redis" yaml:"redis"`
	Attachment Attachment `mapstructure:"attachment" json:"attachment" yaml:"attachment"`
	DataFile   DataFile   `mapstructure:"datafile" json:"datafile" yaml:"datafile"`
	Login      Login      `mapstructure:"login" json:"login" yaml:"login"`
	Other      Other      `mapstructure:"other" json:"other" yaml:"other"`
	Ueditor    Ueditor    `mapstructure:"ueditor" json:"ueditor" yaml:"ueditor"`
	Database   Database   `mapstructure:"database" json:"database" yaml:"database"`
}
