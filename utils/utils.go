package utils

import (
	"os"
	"regexp"
)

// check mobile phone number
func IsTel(tel string) (isMobile bool) {
	isMobile, _ = regexp.MatchString(`^(1[3|4|5|8|7][0-9]\d{4,8})$`, tel)
	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
	}
	return err
}
