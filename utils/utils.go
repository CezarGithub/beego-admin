package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/beego/beego/v2/core/logs"

	"golang.org/x/crypto/bcrypt"
	"quince/internal/captcha"
)

// CaptchaResponse struct
type CaptchaResponse struct {
	CaptchaId  string
	CaptchaUrl string
}

// GetCaptcha 获取验证码
func GetCaptcha() *CaptchaResponse {
	captchaID := captcha.NewLen(4)
	return &CaptchaResponse{
		CaptchaId:  captchaID,
		CaptchaUrl: fmt.Sprintf("/admin/auth/captcha/%s.png", captchaID),
	}
}

// KeyInMap Imitate PHP's array_key_exists to determine whether it exists in the map
func KeyInMap(key string, m map[string]interface{}) bool {
	_, ok := m[key]
	if ok {
		return true
	}
	return false
}

// InArrayForString Imitate php's in_array to determine whether it exists in the string array
func InArrayForString(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// InArrayForInt Imitate php's in_array to determine whether it exists in the int array
func InArrayForInt(items []int64, item int64) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// PasswordHash php function password_hash
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// PasswordVerify php function password_verify
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// IntArrToStringArr int array to string array
func IntArrToStringArr(arr []int) []string {
	var stringArr []string
	for _, v := range arr {
		stringArr = append(stringArr, strconv.Itoa(v))
	}
	return stringArr
}

// GetMd5String MD5 hash the string
func GetMd5String(str string) string {
	t := md5.New()
	io.WriteString(t, str)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// GetSha1String SHA1 hash a string
func GetSha1String(str string) string {
	t := sha1.New()
	io.WriteString(t, str)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// ParseName String naming style conversion
func ParseName(name string, ptype int, ucfirst bool) string {
	if ptype > 0 {
		//Explain regular expressions
		reg := regexp.MustCompile(`_([a-zA-Z])`)
		if reg == nil {
			logs.Error("MustCompile err")
			return ""
		}
		//Extract key information
		result := reg.FindAllStringSubmatch(name, -1)
		for _, v := range result {
			name = strings.ReplaceAll(name, v[0], strings.ToUpper(v[1]))
		}

		if ucfirst {
			return Ucfirst(name)
		}
		return Lcfirst(name)
	}
	//Explain regular expressions
	reg := regexp.MustCompile(`[A-Z]`)
	if reg == nil {
		logs.Error("MustCompile err")
		return ""
	}
	//Extract key information
	result := reg.FindAllStringSubmatch(name, -1)

	for _, v := range result {
		name = strings.ReplaceAll(name, v[0], "_"+v[0])
	}
	return strings.ToLower(name)
}

// Ucfirst Capitalize the first letter
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// Lcfirst Lowercase first letter
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// GetStringInBetweenTwoString - return string between two strings ex : {{bla.bla}}
// r,_:=GetStringInBetweenTwoString("{{bla.bla}}","{{","}}") will return r=bla.bla
func GetStringInBetweenTwoString(str string, startS string, endS string) (result string, found bool) {
	s := strings.Index(str, startS)
	if s == -1 {
		return result, false
	}
	newS := str[s+len(startS):]
	e := strings.Index(newS, endS)
	if e == -1 {
		return result, false
	}
	result = newS[:e]
	return result, true
}
