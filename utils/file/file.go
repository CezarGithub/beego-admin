package file

import (
	"fmt"
	"io"
	"os"

	"github.com/beego/beego/v2/core/logs"
)

func AppendToFile(sourceFile string, destFile string) error {
	source, err := os.Open(sourceFile)
	if err != nil {
		logs.Error("failed to open file: %s ", err)
	}
	defer source.Close()
	dest, err := os.OpenFile(destFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logs.Error("failed to open file: %s ", err)
	}
	defer dest.Close()
	dest.WriteString(fmt.Sprintf("\n#%s  \n", "Generated file. DO NOT MODIFY !!!"))
	dest.WriteString(fmt.Sprintf("\n#%s  \n", sourceFile))
	n, err := io.Copy(dest, source)
	if err != nil {
		logs.Error("failed to append file : %s %d", err, n)
	}
	//logs.Info("wrote %d bytes of %s to %s\n", n, sourceFile, destFile)

	return err
}

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
			logs.Info(err.Error())
		}

	}

	return
}

func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}
