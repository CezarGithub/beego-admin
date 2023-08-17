package i18n

import (
	"os"
	"path"

	"github.com/beego/beego/v2/server/web"

	"github.com/beego/beego/v2/core/logs"
)

//delete old language files
func init() {

	directory, _ := web.AppConfig.String("lang::folder")
	// Open the directory and read all its files.
	dirRead, _ := os.Open(directory)
	dirFiles, _ := dirRead.Readdir(0)
	// Loop over the directory's files.
	for index := range dirFiles {
		fileHere := dirFiles[index]

		// Get name of file and its full path.
		nameHere := fileHere.Name()
		fullPath := path.Join(directory, nameHere)

		// Remove the file.
		os.Remove(fullPath)
		logs.Info("Language file removed : %s", fullPath)
	}
}
