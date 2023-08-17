package url

import (
	"net/url"
	"path"
)

// keep full path (host + path) )without query params
func FullPath(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	result := u.Host + u.Path
	return result, nil
}
func Host(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	result := u.Host
	return result, nil
}
func Path(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	result := u.Path
	return result, nil
}
func Query(urlString string) (url.Values, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	result := u.Query()
	return result, nil
}

func LastSegment(urlString string) (string, error) {
	_, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	result := path.Base(urlString)
	return result, nil
}
