/**
* @Author: scjtqs
* @Date: 2022/7/18 11:28
* @Email: scjtqs@qq.com
 */
package config

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func LoadConf(confFile string) ([]byte, error) {
	var data []byte
	if strings.IndexByte(confFile, filepath.Separator) != 0 {
		cwd, err := GetCurrentPath()
		if err != nil {
			return data, errors.New("svc/invalid_configfile: " + err.Error())
		}
		confFile = filepath.Join(cwd, confFile)
	}

	file, err := os.OpenFile(confFile, os.O_RDONLY, 0666)
	if err != nil {
		return data, err
	}

	defer file.Close()

	data, err = ioutil.ReadAll(file)
	if err != nil {
		return data, err
	}

	// err = yaml.Unmarshal(data, &conf)

	return data, nil
}

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	if runtime.GOOS == "windows" {
		path = strings.Replace(path, "\\", "/", -1)
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		return "", errors.New("system/path_error: " + `Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}
