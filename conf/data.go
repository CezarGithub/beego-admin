package conf

// Data files folder - for export files only !!!
type DataFile struct {
	Path   string `mapstructure:"path" json:"path" yaml:"path"`       //data files  directory configuration (relative to the root directory)
	Prefix string `mapstructure:"prefix" json:"prefix" yaml:"prefix"` //prefix - for data files and name of the arhive
	Ext    string `mapstructure:"ext" json:"ext" yaml:"ext"`          //ext - data files extension
}
