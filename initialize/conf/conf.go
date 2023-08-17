package conf

import (
	"fmt"
	"os"
	"quince/global"
	"quince/utils"
	"gopkg.in/yaml.v3"
)

//Load configuration file
func init() {
	var config string

	if configEnv := os.Getenv(utils.ConfigEnv); configEnv == "" {
		config = utils.ConfigFile
		fmt.Printf("You are using the default value of config, the path of config is %v\n", utils.ConfigFile)
	} else {
		config = configEnv
		fmt.Printf("You are using the default value of config, the path of config is %v\n", config)
	}
	if yamlFile, err := os.ReadFile(config); err != nil {
		panic(err)
	} else {
		if err := yaml.Unmarshal(yamlFile, &global.BA_CONFIG); err != nil {
			panic(err)
		}
	}
	fmt.Println("Configuration file loaded ! ")

}
