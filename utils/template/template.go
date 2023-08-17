//Package template
package template

import (
	"math"
	"strconv"
	"time"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.AddFuncMap("UnixTimeForFormat", UnixTimeForFormat)
	web.AddFuncMap("FormatSize", FormatSize)
	web.AddFuncMap("TimeFormat", TimeFormat)
}

// UnixTimeForFormat
func UnixTimeForFormat(timeUnix int) string {
	//
	timeLayout := "2006-01-02 15:04:05"
	return time.Unix(int64(timeUnix), 0).Format(timeLayout)
}
func TimeFormat(date time.Time) string {
	//
	timeLayout := "2006-01-02 15:04:05"
	return date.Format(timeLayout)
}

// FormatSize
func FormatSize(size, delimiter string) string {
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		return ""
	}
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	var i int
	for i = 0; sizeInt >= 1024 && i < 5; i++ {
		sizeInt /= 1024
	}
	return strconv.FormatFloat(math.Round(float64(sizeInt)), 'f', -1, 64) + delimiter + units[i]
}
