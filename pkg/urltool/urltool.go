package urltool

import (
	"errors"
	"net/url"
	"path"
)

// GetBasePath 获取url路径的最后一节
func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(targetUrl) // 基本上都能解析通过
	if err != nil {
		return "", err
	}
	if len(myUrl.Host) == 0 {
		return "", errors.New("no host on targetUrl")
	}
	return path.Base(myUrl.Path), nil
}
